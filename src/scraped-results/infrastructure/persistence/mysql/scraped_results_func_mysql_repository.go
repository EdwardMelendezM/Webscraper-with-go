package mysql

import (
	"database/sql"
	_ "database/sql"
	_ "embed"

	"github.com/jackskj/carta"
	"github.com/stroiman/go-automapper"

	"github.com/EdwardMelendezM/api-info-shared/db"

	"webscraper-go/scraped-results/domain"
)

//go:embed sql/get_scraped_results.sql
var QueryGetScrapedResults string

func (r ScrapedResultsMysqlRepo) GetScrapedResults(projectId string) (
	topics []domain.ScrapedResult,
	err error,
) {
	results, err := db.Client.Query(
		QueryGetScrapedResults,
		projectId,
	)
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			return
		}
	}(results)

	topicsTmp := make([]ScrapedResult, 0)
	err = carta.Map(results, &topicsTmp)
	if err != nil {
		return nil, err
	}
	automapper.Map(topicsTmp, &topics)
	return topics, nil

}
