package domain

import "github.com/google/uuid"

type Group struct {
	UUID uuid.UUID
	ID   string

	Course   *Course
	Students []*Student
}
