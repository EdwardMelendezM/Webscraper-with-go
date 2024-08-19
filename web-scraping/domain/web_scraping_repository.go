package domain

type WebScrapingRepository interface {
	VerifyExistsUrl(url string) (bool, error)
	GetLastNumber() (*int, error)
	CreateRecord()
}
