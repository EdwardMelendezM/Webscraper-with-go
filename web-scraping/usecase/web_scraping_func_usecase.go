package usecase

import (
	"github.com/google/uuid"
	"webscraper-go/web-scraping/domain"
)

func (u *WebScrapingFuncUseCase) ExtractSearchResults() (bool, error) {
	projectId := "91da2ca7-6244-11ef-9d2f-0242ac110002"
	topics, err := u.TopicsRepository.GetTopics(projectId)
	if err != nil {
		return false, err
	}

	results := make([]domain.SearchResult, 0)

	for _, topic := range topics {
		u.WebScrapingCollectRepository.CollectSearchResults(topic.Title, &results)
	}

	if len(results) == 0 {
		return false, nil
	}

	existingResults, errResult := u.WebScrapingRepository.GetRecordResult(projectId, 1000)

	if errResult != nil {
		return false, errResult
	}

	existingMap := make(map[string]domain.NewRecordWebScraping)
	notExistingResults := make(map[string]domain.NewRecordWebScraping)
	seenMap := make(map[string]domain.NewRecordWebScraping)

	// 1. Map existing results
	for _, existingResult := range existingResults {
		existingMap[existingResult.Url] = domain.NewRecordWebScraping{
			Title: existingResult.Title,
			Url:   existingResult.Url,
			Path:  existingResult.Path,
		}
	}

	// 2. Verify repeated results
	for _, result := range results {
		if _, ok := seenMap[result.Url]; ok {
			continue
		}
		seenMap[result.Url] = domain.NewRecordWebScraping{
			Title: result.Title,
			Url:   result.Url,
			Path:  result.Path,
		}
	}

	// 3. Verify if existing results are in the new results
	for _, result := range seenMap {
		if _, ok := existingMap[result.Url]; !ok {
			notExistingResults[result.Url] = domain.NewRecordWebScraping{
				Title: result.Title,
				Url:   result.Url,
				Path:  result.Path,
			}
			delete(existingMap, result.Url)
		}
	}

	// 4: Get last number
	lastNumber, errLastNumber := u.WebScrapingRepository.GetLastNumber(projectId)
	if errLastNumber != nil {
		return false, errLastNumber
	}

	if lastNumber == nil {
		lastNumber = new(int)
		*lastNumber = 0
	}

	// 5: Add new record notExistingResults
	for _, existingResult := range results {
		id := uuid.New().String()
		body := domain.CreateRecordWebScraping{
			Title:  existingResult.Title,
			Url:    existingResult.Url,
			Number: *lastNumber,
		}
		_, errCreateNewRecord := u.WebScrapingRepository.CreateRecord(id, projectId, body)
		if errCreateNewRecord != nil {
			return false, errCreateNewRecord
		}
		*lastNumber = *lastNumber + 1
	}
	return true, nil
}
