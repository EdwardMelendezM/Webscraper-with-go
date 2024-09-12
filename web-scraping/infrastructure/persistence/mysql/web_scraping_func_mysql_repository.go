package mysql

import (
	_ "embed"
	"time"

	"database/sql"
	_ "database/sql"
	"github.com/jackskj/carta"
	"github.com/stroiman/go-automapper"
	"webscraper-go/web-scraping/domain"

	"github.com/EdwardMelendezM/api-info-shared/db"
)

//go:embed sql/verify_exists_url.sql
var QueryVerifyExistsUrl string

//go:embed sql/get_last_number.sql
var QueryGetLastNumber string

//go:embed sql/create_new_record.sql
var QueryCreateNewRecord string

//go:embed sql/update_record_result.sql
var QueryUpdateRecordResult string

//go:embed sql/get_record_results.sql
var QueryGetRecordResult string

func (r WebScrapingMysqlRepo) VerifyExistsUrl(
	projectId string,
	url string,
) (exists bool, err error) {
	err = db.Client.QueryRow(
		QueryVerifyExistsUrl,
		projectId,
		url,
	).Scan(&exists)

	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r WebScrapingMysqlRepo) GetLastNumber(projectId string) (lastNumber *int, err error) {
	err = db.Client.QueryRow(
		QueryGetLastNumber,
		projectId,
	).Scan(&lastNumber)

	if lastNumber == nil {
		lastNumber = new(int)
		*lastNumber = 0
	}
	return lastNumber, nil
}

func (r WebScrapingMysqlRepo) CreateRecord(
	id string,
	projectId string,
	body domain.CreateRecordWebScraping,
) (lastId *string, err error) {
	now := time.Now()
	_, err = db.Client.Exec(
		QueryCreateNewRecord,
		id,
		projectId,
		body.Title,
		body.Url,
		body.Content,
		body.Number,
		body.TitleCorpus,
		body.ContentCorpus,
		body.WordKey,
		now,
	)
	if err != nil {
		return nil, err
	}
	lastId = &id
	return lastId, nil
}

func (r WebScrapingMysqlRepo) UpdateRecordResult(
	id string,
	projectId string,
	body domain.UpdateRecordWebScraping,
) (err error) {
	_, err = db.Client.Exec(
		QueryUpdateRecordResult,
		body.Content,
		id,
		projectId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r WebScrapingMysqlRepo) GetRecordResult(
	projectId string,
	sizeRecord int,
) (webScrapingResults []domain.WebScrapingResult, err error) {
	results, err := db.Client.Query(
		QueryGetRecordResult,
		projectId,
		sizeRecord,
	)
	if err != nil {
		return webScrapingResults, err
	}

	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			panic(errClose)
		}
	}(results)

	webScrapingResultsTmp := make([]WebScrapingResult, 0)
	err = carta.Map(results, &webScrapingResultsTmp)
	if err != nil {
		return webScrapingResults, err
	}
	automapper.Map(webScrapingResultsTmp, &webScrapingResults)
	return webScrapingResults, nil
}
