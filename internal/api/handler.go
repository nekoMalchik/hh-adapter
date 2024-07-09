package api

import (
	"encoding/json"
	"hh-adapter/internal/entity"
	"hh-adapter/internal/service"
	"net/http"
)

type Handler struct {
	VacancyService service.VacancyService
}

func (h *Handler) CreateVacancyHandler(w http.ResponseWriter, r *http.Request) {
	var jobs []entity.Vacancy

	if err := json.NewDecoder(r.Body).Decode(&jobs); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.VacancyService.CreateJobs(jobs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
