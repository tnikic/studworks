package db

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"hcw.ac.at/studworks/internal/errs"
)

type Postgres struct {
	connection *pgx.Conn
	Ctx        context.Context
	Queries    *Queries
}

func (p *Postgres) Connect() error {
	url := os.Getenv("DB_Url")
	user := os.Getenv("DB_User")
	password := os.Getenv("DB_Password")
	dbname := os.Getenv("DB_DBName")

	ctx := context.Background()
	p.Ctx = ctx

	connection, err := pgx.Connect(
		ctx,
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s",
			user,
			password,
			url,
			dbname,
		),
	)
	if err != nil {
		slog.Error("Postgres connection failed.", slog.Any("err", err))
		httpError := errs.NewHttpError(500, "Postgres connection failed", err)
		return httpError
	}
	p.connection = connection
	p.Queries = New(p.connection)

	slog.Debug("Postgres connection established successfully.")
	return nil
}

func (p *Postgres) Close() {
	if p.connection != nil {
		err := p.connection.Close(p.Ctx)
		if err != nil {
			slog.Error("Failed to close Postgres connection.", slog.Any("err", err))
		} else {
			slog.Debug("Postgres connection closed successfully.")
		}
	}

	p.connection = nil
	p.Ctx = nil
	p.Queries = nil
}
