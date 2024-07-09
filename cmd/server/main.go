package main

import (
	"fmt"
	"hh-adapter/internal/api"
	"hh-adapter/internal/config"
	"hh-adapter/internal/hh"
	"hh-adapter/internal/repository"
	"hh-adapter/internal/service"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	fmt.Println(cfg)

	hhClient := hh.NewHHClient(cfg.HH.APIKey)
	repo, err := repository.NewVacancyRepository(cfg.DB.Uri, cfg.DB.DbName, "vacancy")
	if err != nil {
		log.Fatalf("Error loading repository: %v", err)
	}

	jobService := service.NewVacancyService(repo, hhClient)

	router := api.NewRouter(jobService)

	log.Println("Starting server on port", cfg.Server.Port)
	if err := http.ListenAndServe(":"+cfg.Server.Port, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
