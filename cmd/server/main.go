package main

import (
	"fmt"
	"log"
	"my-hh-adapter/internal/api"
	"my-hh-adapter/internal/config"
	"my-hh-adapter/internal/hh"
	"my-hh-adapter/internal/repository"
	"my-hh-adapter/internal/service"
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
