package directory

import (
	"crypto/tls"
	"log/slog"
	"os"

	"github.com/go-ldap/ldap/v3"
	"hcw.ac.at/studworks/internal/errs"
)

type LDAP struct {
	conn    *ldap.Conn
	Queries *Queries
}

type Queries struct{}

func (l *LDAP) Connect() error {
	server := os.Getenv("LDAP_Server")
	bind := os.Getenv("LDAP_Bind")
	password := os.Getenv("LDAP_Password")

	tlsConfig := &tls.Config{
		ServerName:         "portal.fh-campuswien.ac.at",
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS10,
		MaxVersion:         tls.VersionTLS12,
	}

	conn, err := ldap.DialURL(server, ldap.DialWithTLSConfig(tlsConfig))
	if err != nil {
		slog.Error("Failed to connect to LDAP server.", slog.Any("err", err))
		httpError := errs.NewHttpError(500, "LDAP connection failed", err)
		return httpError
	}

	err = conn.Bind(bind, password)
	if err != nil {
		slog.Error("LDAP bind failed.", slog.Any("err", err))
		httpError := errs.NewHttpError(500, "LDAP bind failed", err)
		return httpError
	}

	l.conn = conn
	l.Queries = &Queries{}

	return nil
}

func (l *LDAP) Close() {
	if l.conn != nil {
		err := l.conn.Close()
		if err != nil {
			slog.Error("Failed to close LDAP connection.", slog.Any("err", err))
		}

		l.conn = nil
	}
}

func (l *LDAP) search(filter string, args []string) *ldap.SearchRequest {
	base := os.Getenv("LDAP_Base")

	return ldap.NewSearchRequest(
		base,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0, 0,
		false,
		filter,
		args,
		nil,
	)
}
