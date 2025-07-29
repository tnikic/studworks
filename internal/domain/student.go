package domain

import (
	"github.com/go-ldap/ldap/v3"
	"hcw.ac.at/studworks/internal/repository/db"
)

type Student struct {
	UID       string `json:"uid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Active    bool   `json:"active"`
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
	s.Email = ldapStudent.GetAttributeValue("mail")
	s.Active = true
}
