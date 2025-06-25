package domain

type WebScrapingUseCase interface {
	ExtractSearchResults() (bool bool, err error)
}
