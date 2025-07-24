package utils

import (
	"os"

	"gitlab.com/gitlab-org/api/client-go"
	"hcw.ac.at/studworks/internal/errs"
)

type Gitlab struct{}

func (g *Gitlab) Connect() (*gitlab.Client, error) {
	token := os.Getenv("GITLAB_Token")

	client, err := gitlab.NewClient(token)
	if err != nil {
		httpError := errs.NewHttpError(500, "Failed to create GitLab client", err)
		return nil, httpError
	}

	return client, nil
}
