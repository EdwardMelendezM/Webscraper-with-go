package domain

type WebScrapingRepository interface {
	VerifyExistsUrl(url string) (bool, error)
	GetLastNumber() (lastNumber *int, err error)
	CreateRecord(id string, body CreateRecordWebScraping) (lastId *string, err error)
	UpdateRecordResult(id string, body UpdateRecordWebScraping) (err error)
	GetRecordResult() []WebScrapingResult
}

type WebScrapingCollectRepository interface {
	CollectSearchResults(topic string, resultsChan chan<- SearchResult)
}
