package class

import (
	"errors"
	"net/http"

	"hcw.ac.at/studworks/internal/errs"
)

type Handler struct{}

func (h *Handler) CreateClass(w http.ResponseWriter, r *http.Request) {
	className := r.PathValue("className")
	if className == "" {
		http.Error(w, "Class name is required", http.StatusBadRequest)
		return
	}

	service := &Service{}
	err := service.CreateClass(className)
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
