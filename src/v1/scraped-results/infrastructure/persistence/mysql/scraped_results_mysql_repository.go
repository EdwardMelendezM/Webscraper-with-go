package mysql

import (
	"webscraper-go/v1/scraped-results/domain"
)

type ScrapedResultsMysqlRepo struct {
}

func NewScrapedResultRepository() domain.ScrapedResultRepository {
	return &ScrapedResultsMysqlRepo{}
}
