package domain

import "github.com/google/uuid"

type AssignmentTemplate struct {
	UUID  uuid.UUID
	Title string
}

type Assignment struct {
	UUID     uuid.UUID
	Template *AssignmentTemplate
	Student  *Student
	Course   *Course

	PossiblePoints int
	PointsEarned   int
}
