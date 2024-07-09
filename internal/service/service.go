package service

import (
	"hh-adapter/internal/entity"
	"hh-adapter/internal/hh"
	"hh-adapter/internal/repository"
)

type VacancyService interface {
	CreateJobs(vacancies []entity.Vacancy) error
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

func (s *vacancyService) CreateJobs(vacancies []entity.Vacancy) error {
	// Ваш код для создания вакансий в БД
	return nil
}
