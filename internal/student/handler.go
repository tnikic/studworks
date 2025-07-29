package student

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"hcw.ac.at/studworks/internal/domain"
	"hcw.ac.at/studworks/internal/errs"
)

type Handler struct{}

func (h *Handler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	student := &domain.Student{}
	err := json.NewDecoder(r.Body).Decode(student)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	service := &Service{}
	err = service.CreateStudent(student)
	if err != nil {
		var httpErr *errs.HttpError
		if errors.As(err, &httpErr) {
			httpError := err.(*errs.HttpError)
			http.Error(w, httpError.Message, httpError.Code)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}

func (h *Handler) SearchStudents(w http.ResponseWriter, r *http.Request) {
	className := r.PathValue("className")
	if className == "" {
		http.Error(w, "Class name is required", http.StatusBadRequest)
		return
	}

	service := &Service{}
	students, err := service.SearchStudents(className)
	if err != nil {
		var httpErr *errs.HttpError
		if errors.As(err, &httpErr) {
			http.Error(w, httpErr.Message, httpErr.Code)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	slog.Info("Found students", "count", len(students))

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(students)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
