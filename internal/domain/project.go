package domain

import "github.com/google/uuid"

type ProjectParticipant struct {
	UUID    uuid.UUID
	Student *Student
	Project *Project

	PossiblePoints int
	PointsEarned   int
}

type Project struct {
	UUID  uuid.UUID
	Title string

	Course       *Course
	Participants []*ProjectParticipant
}
