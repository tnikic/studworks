package class

import (
	"hcw.ac.at/studworks/internal/domain"
	"hcw.ac.at/studworks/internal/errs"
	"hcw.ac.at/studworks/internal/repository/db"
)

type Service struct{}

func (s *Service) CreateClass(name string) error {
	class := &domain.Class{}
	err := class.ExpandClass(name)
	if err != nil {
		return err
	}

	class.Active = true

	var pg db.Postgres
	err = pg.Connect()
	if err != nil {
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
		return err
	}

	return nil
}
