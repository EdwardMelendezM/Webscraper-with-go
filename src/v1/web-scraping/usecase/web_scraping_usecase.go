package usecase

import (
	domainTopics "webscraper-go/topics/domain"

	"webscraper-go/web-scraping/domain"
)

type WebScrapingFuncUseCase struct {
	TopicsRepository             domainTopics.TopicsRepository
	WebScrapingRepository        domain.WebScrapingRepository
	WebScrapingCollectRepository domain.WebScrapingCollectRepository
}

func NewWebScrapingFuncUseCase(
	webScrapingRepository domain.WebScrapingRepository,
	webScrapingCollectRepository domain.WebScrapingCollectRepository,
	topicsRepository domainTopics.TopicsRepository,
) domain.WebScrapingUseCase {
	return &WebScrapingFuncUseCase{
		TopicsRepository:             topicsRepository,
		WebScrapingRepository:        webScrapingRepository,
		WebScrapingCollectRepository: webScrapingCollectRepository,
	}
}
