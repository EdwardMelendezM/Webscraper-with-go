package mysql

import "webscraper-go/web-scraping/domain"

type WebScrapingMysqlRepo struct{}

func NewWebScrapingRepository() domain.WebScrapingRepository {
	return &WebScrapingMysqlRepo{}
}
