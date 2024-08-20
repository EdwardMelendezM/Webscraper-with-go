package domain

type WebScrapingRepository interface {
	VerifyExistsUrl(url string) (bool, error)
	GetLastNumber() (lastNumber *int, err error)
	CreateRecord(id string, body CreateRecordWebScraping) (lastId *string, err error)
}
