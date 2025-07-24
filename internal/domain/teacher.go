package domain

type Teacher struct {
	UID       string
	FirstName string
	LastName  string
	Email     string
	Active    bool

	Courses []*Course
}
