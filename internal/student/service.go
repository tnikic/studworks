package student

import (
	"log/slog"

	"github.com/jackc/pgx/v5"
	"hcw.ac.at/studworks/internal/domain"
	"hcw.ac.at/studworks/internal/errs"
	"hcw.ac.at/studworks/internal/repository/db"
	"hcw.ac.at/studworks/internal/repository/directory"
)

type Service struct{}

func (s *Service) CreateStudent(student *domain.Student) error {
	existingStudent, err := s.GetStudent(student.UID)
	if !(err != nil && existingStudent != nil) {
		if existingStudent != nil {
			// Student already exists, return an error
			httpError := errs.NewHttpError(409, "Student already exists", nil)
			return httpError
		}
		return err
	}

	var pg db.Postgres
	err = pg.Connect()
	if err != nil {
		httpError := errs.NewHttpError(500, "Postgres connection failed", err)
		return httpError
	}
	defer pg.Close()

	params := db.CreateStudentParams{
		Uid:       student.UID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Email:     student.Email,
		Active:    student.Active,
	}

	_, err = pg.Queries.CreateStudent(pg.Ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetStudent(uid string) (*domain.Student, error) {
	var pg db.Postgres
	err := pg.Connect()
	if err != nil {
		return nil, err
	}
	defer pg.Close()

	student, err := pg.Queries.GetStudentByUID(pg.Ctx, uid)
	if err == pgx.ErrNoRows {
		// Student not found
		noStudent := &domain.Student{}
		httpError := errs.NewHttpError(404, "Student not found", nil)
		return noStudent, httpError
	} else if err != nil {
		// Other DB error
		slog.Error("Failed to get student from Postgres.", "error", err)
		httpError := errs.NewHttpError(500, "Database error", err)
		return nil, httpError
	}

	domainStudent := &domain.Student{}
	domainStudent.ConvertFromDB(&student)

	return domainStudent, nil
}

func (s *Service) SearchStudents(className string) ([]*domain.Student, error) {
	students := []*domain.Student{}

	l := directory.LDAP{}
	entries, err := l.Queries.ListStudents(className)
	if err != nil {
		return nil, err
	}

	class := &domain.Class{}
	err = class.ExpandClass(className)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		student := &domain.Student{}
		student.ConvertFromRegistry(entry)

		students = append(students, student)
	}

	return students, nil
}
