// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package db

import (
	"github.com/google/uuid"
)

type Assignment struct {
	Uuid                   uuid.UUID
	AssignmentTemplateUuid uuid.UUID
	StudentUid             string
	CourseUuid             uuid.UUID
	PossiblePoints         int32
	PointsEarned           int32
}

type AssignmentTemplate struct {
	Uuid  uuid.UUID
	Title string
}

type Class struct {
	Name        string
	ProgramCode string
	Year        int32
	StudyType   string
	Active      bool
}

type Course struct {
	Uuid        uuid.UUID
	SubjectUuid uuid.UUID
	TeacherUid  string
	ClassName   string
}

type Group struct {
	Uuid       uuid.UUID
	ID         string
	CourseUuid uuid.UUID
}

type Project struct {
	Uuid           uuid.UUID
	Title          string
	Status         string
	PossiblePoints int32
	CourseUuid     uuid.UUID
}

type ProjectParticipant struct {
	Uuid         uuid.UUID
	PointsEarned int32
	StudentUid   string
	ProjectUuid  uuid.UUID
}

type Student struct {
	Uid       string
	FirstName string
	LastName  string
	Email     string
	Active    bool
	ClassName string
}

type Subject struct {
	Uuid        uuid.UUID
	Name        string
	ProgramCode string
	Semester    int32
}

type Teacher struct {
	Uid       string
	FirstName string
	LastName  string
	Email     string
	Active    bool
}
