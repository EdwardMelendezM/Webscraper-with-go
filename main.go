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

	webScraperUseCase.NewWebScrapingFuncUseCase(
		webScrapingRepository,
		webScrapingCollectRepository,
		topicsRepository)
	fmt.Printf("Finished scraping")
}
