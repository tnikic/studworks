package directory

import (
	"log/slog"

	"hcw.ac.at/studworks/internal/domain"
	"hcw.ac.at/studworks/internal/errs"
)

func (q *Queries) ListClasses() ([]*domain.Class, error) {
	l := LDAP{}

	filter := "(|(departmentName=CSDC*)(departmentName=SDE*))"
	request := l.search(filter, []string{"departmentName"})

	err := l.Connect()
	if err != nil {
		return nil, err
	}

	response, err := l.conn.Search(request)
	if err != nil {
		slog.Error("LDAP search failed", "error", err)
		httpError := errs.NewHttpError(500, "LDAP search failed", err)
		return []*domain.Class{}, httpError
	}

	classes := []*domain.Class{}
	for _, entry := range response.Entries {
		name := entry.GetAttributeValue("departmentName")

		var class *domain.Class
		err := class.ExpandClass(name)
		if err != nil {
			return nil, err
		}

		classes = append(classes, class)
	}

	return classes, nil
}
