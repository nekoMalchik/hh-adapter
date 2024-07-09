package api

import (
	"github.com/gorilla/mux"
	"hh-adapter/internal/service"
	"net/http"
)

func NewRouter(vacancyService service.VacancyService) *mux.Router {
	r := mux.NewRouter()
	h := Handler{vacancyService}
	r.HandleFunc("/jobs", h.CreateVacancyHandler).Methods(http.MethodGet)
	return r
}
