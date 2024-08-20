package usecase

import (
	"github.com/google/uuid"
	"webscraper-go/web-scraping/domain"
)

func (u *WebScrapingFuncUseCase) ExtractSearchResults() (bool bool, err error) {

	topics, err := u.TopicsRepository.GetTopics()
	if err != nil {
		return false, err
	}
	resultsChan := make(chan domain.SearchResult)
	for _, topic := range topics {
		go u.WebScrapingCollectRepository.CollectSearchResults(topic.Title, resultsChan)
	}
	var results []domain.SearchResult
	for result := range resultsChan {
		results = append(results, domain.SearchResult{
			Title:   result.Title,
			Url:     result.Url,
			Content: result.Content,
		})
	}
	close(resultsChan)
	if len(results) == 0 {
		return false, nil
	}
	for _, result := range results {
		exists, errVerify := u.WebScrapingRepository.VerifyExistsUrl(result.Url)
		if errVerify != nil {
			return false, errVerify
		}
		if exists == false {
			lastNumber, errLastNumber := u.WebScrapingRepository.GetLastNumber()
			if errLastNumber != nil {
				return false, errLastNumber
			}
			if lastNumber == nil {
				lastNumber = new(int)
				*lastNumber = 0
			}
			id := uuid.New().String()
			_, errCreateNewRecord := u.WebScrapingRepository.CreateRecord(id, domain.CreateRecordWebScraping{
				Title:  result.Title,
				Url:    result.Url,
				Number: *lastNumber + 1,
			})
			if errCreateNewRecord != nil {
				return false, errLastNumber
			}
		}
	}

	return true, nil
}
