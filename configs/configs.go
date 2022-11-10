package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	ElasticSearchConfigs ElasticSearchConfigs
}

type ElasticSearchConfigs struct {
	BaseURL string
}

func ReadConfigs() *Configs {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to read .env: %v", err)
	}

	elasticSearchConfig := ElasticSearchConfigs{
		BaseURL: os.Getenv("ELASTIC_SEARCH.BASE_URL"),
	}
	return &Configs{
		ElasticSearchConfigs: elasticSearchConfig,
	}
}
