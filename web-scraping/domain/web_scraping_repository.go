package domain

type WebScrapingRepository interface {
	VerifyExistsUrl(projectId string, url string) (bool, error)
	GetLastNumber(projectId string) (lastNumber *int, err error)
	CreateRecord(id string, projectId string, body CreateRecordWebScraping) (lastId *string, err error)
	UpdateRecordResult(id string, projectId string, body UpdateRecordWebScraping) (err error)
	GetRecordResult(projectId string, sizeRecord int) (WebScrapingResults []WebScrapingResult, err error)
}

type WebScrapingCollectRepository interface {
	CollectSearchResults(topic string, resultsChan *[]SearchResult)
	CollectSearchResultsAndReturn(topic string) (resultsChan []SearchResult, err error)
}
