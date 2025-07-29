package directory

import (
	"fmt"
	"log/slog"

	"github.com/go-ldap/ldap/v3"
	"hcw.ac.at/studworks/internal/errs"
)

func (q *Queries) ListStudents(className string) ([]*ldap.Entry, error) {
	l := LDAP{}

	filter := fmt.Sprintf("(departmentNumber=%s)", className)
	request := l.search(filter, nil)

	err := l.Connect()
	if err != nil {
		return nil, err
	}
	defer l.Close()

	response, err := l.conn.Search(request)
	if err != nil {
		slog.Error("LDAP search failed.", "error", err)
		httpError := errs.NewHttpError(500, "LDAP search failed", err)
		return nil, httpError
	}

	return response.Entries, nil
}
