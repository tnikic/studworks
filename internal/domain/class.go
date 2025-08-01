package domain

import (
	"fmt"
	"log/slog"
	"regexp"
	"strconv"
	"strings"

	"hcw.ac.at/studworks/internal/errs"
	"hcw.ac.at/studworks/internal/repository/db"
)

type Class struct {
	Name        string `json:"name"`
	ProgramCode string `json:"program_code"`
	Year        int    `json:"year"`
	StudyType   string `json:"study_type"`
	Active      bool   `json:"active"`

	Students []*Student `json:"students"`
}

func (c *Class) ExpandClass(name string) error {
	pattern := `^([A-Z]{3,4})(\d{2})([A-Z]{2})$`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(strings.ToUpper(name))
	if matches == nil {
		slog.Error("Invalid class name format", slog.String("name", name))
		httpError := errs.NewHttpError(400, "Invalid class name format", nil)
		return httpError
	}

	program := matches[1]
	yearString := matches[2]
	studyType := matches[3]

	year, err := strconv.Atoi(fmt.Sprintf("20%s", yearString))
	if err != nil {
		slog.Error("Invalid year in class name", slog.String("name", name), slog.Any("err", err))
		httpError := errs.NewHttpError(400, "Invalid year in class name", err)
		return httpError
	}

	c.Name = name
	c.ProgramCode = program
	c.Year = year
	c.StudyType = studyType

	return nil
}

func (c *Class) ConvertFromDB(dbClass *db.Class) {
	c.Name = dbClass.Name
	c.ProgramCode = dbClass.ProgramCode
	c.Year = int(dbClass.Year)
	c.StudyType = dbClass.StudyType
	c.Active = dbClass.Active
}
