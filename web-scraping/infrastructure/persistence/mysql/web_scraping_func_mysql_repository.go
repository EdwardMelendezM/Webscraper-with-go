package mysql

import (
	_ "database/sql"
	_ "embed"
	"time"

	"github.com/EdwardMelendezM/api-info-shared/db"
	"webscraper-go/web-scraping/domain"
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
	url string,
) (exists bool, err error) {
	err = db.Client.QueryRow(
		QueryVerifyExistsUrl,
		url,
	).Scan(&exists)

	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r WebScrapingMysqlRepo) GetLastNumber() (lastNumber *int, err error) {
	err = db.Client.QueryRow(
		QueryGetLastNumber,
	).Scan(&lastNumber)

	if lastNumber == nil {
		lastNumber = new(int)
		*lastNumber = 0
	}
	return lastNumber, nil
}

func (r WebScrapingMysqlRepo) CreateRecord(
	id string,
	body domain.CreateRecordWebScraping,
) (lastId *string, err error) {
	now := time.Now()
	_, err = db.Client.Exec(
		QueryCreateNewRecord,
		id,
		body.Title,
		body.Url,
		body.Number,
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
	body domain.UpdateRecordWebScraping,
) (err error) {
	_, err = db.Client.Exec(
		QueryUpdateRecordResult,
		body.Content,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r WebScrapingMysqlRepo) GetRecordResult(
	sizeRecord int,
) (webScrapingResults domain.WebScrapingResult) {

}
