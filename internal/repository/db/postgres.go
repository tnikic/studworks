package db

import (
	"context"
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

	ctx := context.Background()
	p.Ctx = ctx

	connection, err := pgx.Connect(
		ctx,
		url,
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

func TestPostgres() error {
	p := &Postgres{}
	err := p.Connect()
	if err != nil {
		return err
	}
	defer p.Close()
	return nil
}
