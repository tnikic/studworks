package domain

import (
	"github.com/go-ldap/ldap/v3"
	"hcw.ac.at/studworks/internal/repository/db"
)

type Student struct {
	UID       string
	FirstName string
	LastName  string
	Email     string
	Active    bool

	Class       *Class
	Courses     []*Course
	Assignments []*Assignment
	Projects    []*Project
}

func (s *Student) ConvertFromDB(dbStudent *db.Student) {
	s.UID = dbStudent.Uid
	s.FirstName = dbStudent.FirstName
	s.LastName = dbStudent.LastName
	s.Email = dbStudent.Email
	s.Active = dbStudent.Active
}

func (s *Student) ConvertFromRegistry(ldapStudent *ldap.Entry) {
	s.UID = ldapStudent.GetAttributeValue("uid")
	s.FirstName = ldapStudent.GetAttributeValue("givenName")
	s.LastName = ldapStudent.GetAttributeValue("sn")
	s.Email = ldapStudent.GetAttributeValue("email")
	s.Active = true
}
