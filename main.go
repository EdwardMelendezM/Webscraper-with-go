package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/google/uuid"

	TopicsRepository "webscraper-go/topics/infrastructure/persistence/mysql"

	WebScrapingRepository "webscraper-go/web-scraping/infrastructure/persistence/mysql"
	WebScrapingCollectRepository "webscraper-go/web-scraping/infrastructure/scraping/collect"
	webScraperUseCase "webscraper-go/web-scraping/usecase"
)

func main() {
	topicsRepository := TopicsRepository.NewTopicsRepository()
	webScrapingRepository := WebScrapingRepository.NewWebScrapingRepository()
	webScrapingCollectRepository := WebScrapingCollectRepository.NewWebScrapingCollectRepository()

	instance := webScraperUseCase.NewWebScrapingFuncUseCase(
		webScrapingRepository,
		webScrapingCollectRepository,
		topicsRepository)
	value, err := instance.ExtractSearchResults()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	if value {
		fmt.Printf("Scraping was successful")
	} else {
		fmt.Printf("Scraping was not successful")
	}
	fmt.Printf("Finished scraping")
}
