package service

import (
	"my-hh-adapter/internal/hh"
	"my-hh-adapter/internal/repository"
)

type VacancyService interface {
	CreateJobs(vacancies []Vacancy) error
}

type vacancyService struct {
	repo     *repository.VacancyRepository
	hhClient *hh.HHClient
}

func NewVacancyService(repo *repository.VacancyRepository, hhClient *hh.HHClient) VacancyService {
	return &vacancyService{
		repo:     repo,
		hhClient: hhClient,
	}
}

func (s *vacancyService) CreateJobs(vacancies []Vacancy) error {
	// Ваш код для создания вакансий в БД
	return nil
}
