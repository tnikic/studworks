package class

import (
	"log/slog"

	"github.com/jackc/pgx/v5"
	"hcw.ac.at/studworks/internal/domain"
	"hcw.ac.at/studworks/internal/errs"
	"hcw.ac.at/studworks/internal/repository/db"
)

type Service struct{}

func (s *Service) CreateClass(name string) error {
	existingClass, err := s.GetClass(name)
	if !(err != nil && existingClass != nil) {
		if existingClass != nil {
			// Class already exists, return an error
			httpError := errs.NewHttpError(409, "Class already exists", nil)
			return httpError
		}
		return err
	}

	class := &domain.Class{}
	err = class.ExpandClass(name)
	if err != nil {
		return err
	}

	class.Active = true

	var pg db.Postgres
	err = pg.Connect()
	if err != nil {
		slog.Error("Postgres connection failed.", slog.Any("err", err))
		httpError := errs.NewHttpError(500, "Postgres connection failed", err)
		return httpError
	}
	defer pg.Close()

	params := db.CreateClassParams{
		Name:        class.Name,
		Active:      class.Active,
		Year:        int32(class.Year),
		ProgramCode: class.ProgramCode,
		StudyType:   class.StudyType,
	}

	_, err = pg.Queries.CreateClass(pg.Ctx, params)
	if err != nil {
		slog.Error("Failed to create class in Postgres.", slog.Any("err", err))
		httpError := errs.NewHttpError(500, "Failed to create class in Postgres", err)
		return httpError
	}

	return nil
}

func (s *Service) GetClass(name string) (*domain.Class, error) {
	var pg db.Postgres
	err := pg.Connect()
	if err != nil {
		return nil, err
	}
	defer pg.Close()

	class, err := pg.Queries.GetClassByName(pg.Ctx, name)
	if err == pgx.ErrNoRows {
		noClass := &domain.Class{}
		httpError := errs.NewHttpError(404, "Class not found", nil)
		return noClass, httpError
	} else if err != nil {
		slog.Error("Failed to get class from Postgres.", slog.Any("err", err))
		httpError := errs.NewHttpError(500, "Database error", err)
		return nil, httpError
	}

	domainClass := &domain.Class{}
	domainClass.ConvertFromDB(&class)

	return domainClass, nil
}
