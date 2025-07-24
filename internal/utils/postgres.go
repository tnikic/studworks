package utils

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"hcw.ac.at/studworks/internal/errs"
)

type Postgres struct{}

func (p *Postgres) Connect() (*pgx.Conn, context.Context, error) {
	url := os.Getenv("DB_Url")
	user := os.Getenv("DB_User")
	password := os.Getenv("DB_Password")
	dbname := os.Getenv("DB_DBName")

	ctx := context.Background()

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
		return nil, nil, httpError
	}

	return connection, ctx, nil
}
