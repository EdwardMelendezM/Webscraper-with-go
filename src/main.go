package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/EdwardMelendezM/api-info-shared/config"
	"github.com/EdwardMelendezM/api-info-shared/db"

	TopicsRepository "webscraper-go/topics/infrastructure/persistence/mysql"
	WebScrapingRepository "webscraper-go/web-scraping/infrastructure/persistence/mysql"
	WebScrapingCollectRepository "webscraper-go/web-scraping/infrastructure/scraping/collect"
	webScraperUseCase "webscraper-go/web-scraping/usecase"
)

func main() {
	cfg := config.Configuration{
		ServerPort:  os.Getenv("SERVER_PORT"),
		StoragePath: os.Getenv("STORAGE_PATH"),
		DB: config.DB{
			DbDatabase: os.Getenv("DB_DATABASE"),
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbUsername: os.Getenv("DB_USERNAME"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
	}

	err := db.InitClients(cfg)
	if err != nil {
		return
	}
	defer func(Client *sql.DB) {
		errClient := Client.Close()
		if errClient != nil {
			fmt.Printf("Error db: %v", errClient)
		}
	}(db.Client)

	topicsRepository := TopicsRepository.NewTopicsRepository()
	webScrapingRepository := WebScrapingRepository.NewWebScrapingRepository()
	webScrapingCollectRepository := WebScrapingCollectRepository.NewWebScrapingCollectRepository()

	instance := webScraperUseCase.NewWebScrapingFuncUseCase(
		webScrapingRepository,
		webScrapingCollectRepository,
		topicsRepository)

	value, errExtract := instance.ExtractSearchResults()
	if errExtract != nil {
		fmt.Printf("Error: %v", errExtract)
	}
	if value {
		fmt.Printf("Scraping was successful")
	} else {
		fmt.Printf("Scraping was not successful")
	}
	fmt.Printf("Finished scraping")

}
