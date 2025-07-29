package student

import (
	"hcw.ac.at/studworks/internal/domain"
	"hcw.ac.at/studworks/internal/errs"
	"hcw.ac.at/studworks/internal/repository/db"
	"hcw.ac.at/studworks/internal/repository/directory"
)

type Service struct{}

func (s *Service) CreateStudent(student *domain.Student) error {
	var pg db.Postgres
	err := pg.Connect()
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

	dbStudent, err := pg.Queries.CreateStudent(pg.Ctx, params)
	if err != nil {
		return err
	}

	student.UID = dbStudent.Uid
	return nil
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
