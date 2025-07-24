package domain

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"hcw.ac.at/studworks/internal/errs"
)

type Class struct {
	Name        string
	ProgramCode string
	Year        int
	StudyType   string
	Active      bool

	Students []*Student
}

func (c *Class) ExpandClass(name string) error {
	pattern := `^([A-Z]{3,4})(\d{2})([A-Z]{2})$`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(strings.ToUpper(name))
	if matches == nil {
		httpError := errs.NewHttpError(400, "Invalid class name format", nil)
		return httpError
	}

	program := matches[1]
	yearString := matches[2]
	studyType := matches[3]

	year, err := strconv.Atoi(fmt.Sprintf("20%s", yearString))
	if err != nil {
		httpError := errs.NewHttpError(400, "Invalid year in class name", err)
		return httpError
	}

	c.Name = name
	c.ProgramCode = program
	c.Year = year
	c.StudyType = studyType

	return nil
}
