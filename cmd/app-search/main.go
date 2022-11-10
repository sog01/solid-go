package main

import (
	"log"

	"github.com/sog01/solid-go/configs"
	"github.com/sog01/solid-go/internal/search/handler"
	"github.com/sog01/solid-go/internal/search/repository"
	"github.com/sog01/solid-go/internal/search/service"
)

func main() {
	// read a configs instance
	configs := configs.ReadConfigs()

	// construct a repository
	repository := repository.NewElasticSearch(configs.ElasticSearchConfigs.BaseURL)

	// do a health check
	// if error abort operation immediately
	checkHealth(repository)

	// construct services
	searchService := service.NewSearchService(repository)
	syncService := service.NewSyncService(repository)
	constructService := service.NewConstructService(repository)

	// initiate an index
	constructService.CreateIndex()

	// construct a rest handler
	rest := handler.NewRest(searchService, syncService)

	// listen and serve server
	rest.ListenAndServe()
}

func checkHealth(repository repository.Repository) {
	if err := repository.CheckHealth(); err != nil {
		log.Fatalf("missing elasticsearch connection: %v", err)
	}
}
