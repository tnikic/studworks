package domain

import "github.com/google/uuid"

type Subject struct {
	UUID        uuid.UUID
	Name        string
	ProgramCode string
	Semester    int
}

type Course struct {
	UUID    uuid.UUID
	Subject *Subject
	Teacher *Teacher
	Class   *Class
	Year    int
}
