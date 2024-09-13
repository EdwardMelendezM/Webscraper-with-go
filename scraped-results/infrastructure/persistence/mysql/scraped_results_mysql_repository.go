package mysql

import "webscraper-go/scraped-results/domain"

type ScrapedResultsMysqlRepo struct {
}

func NewScrapedResultRepository() domain.ScrapedResultRepository {
	return &ScrapedResultsMysqlRepo{}
}
