package vcs

import (
	"log/slog"
	"os"

	"gitlab.com/gitlab-org/api/client-go"
	"hcw.ac.at/studworks/internal/errs"
)

type Gitlab struct {
	client  *gitlab.Client
	Queries *Queries
}

type Queries struct{}

func (g *Gitlab) Connect() error {
	url := os.Getenv("GL_URL")
	token := os.Getenv("GL_Token")

	client, err := gitlab.NewClient(
		token,
		gitlab.WithBaseURL(url),
	)
	if err != nil {
		slog.Error("Gitlab connection failed.", slog.Any("err", err))
		httpError := errs.NewHttpError(500, "Gitlab connection failed", err)
		return httpError
	}

	g.client = client
	g.Queries = &Queries{}

	return nil
}
