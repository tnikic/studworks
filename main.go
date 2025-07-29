package main

import (
	"log"
	"log/slog"
	"net/http"

	"hcw.ac.at/studworks/internal/class"
	"hcw.ac.at/studworks/internal/repository/db"
	"hcw.ac.at/studworks/internal/repository/directory"
	"hcw.ac.at/studworks/internal/student"
)

func main() {
	requirementsCheck()

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./web"))

	classHandler := class.Handler{}
	mux.HandleFunc("POST /api/classes", classHandler.CreateClass)

	studentHandler := student.Handler{}
	mux.HandleFunc("POST /api/students", studentHandler.CreateStudent)
	mux.HandleFunc("GET /api/students/{className}", studentHandler.SearchStudents)

	// Handle static files
	mux.Handle("/", files)

	slog.Info("Starting Server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func requirementsCheck() {
	errDB := db.TestPostgres()
	errLDAP := directory.TestLDAP()

	// Use unicode checkmark and cross for status
	const (
		check = "\u2713"
		cross = "\u2717"
	)

	slog.Info("Requirements Check:")
	slog.Info(
		"Postgres",
		slog.String("status", map[bool]string{true: check, false: cross}[errDB == nil]),
	)
	slog.Info(
		"LDAP",
		slog.String("status", map[bool]string{true: check, false: cross}[errLDAP == nil]),
	)
}
