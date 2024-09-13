package domain

type ScrapedResultRepository interface {
	GetScrapedResults(projectId string) ([]ScrapedResult, error)
}
