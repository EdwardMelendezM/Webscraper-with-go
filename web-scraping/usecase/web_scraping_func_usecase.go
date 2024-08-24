package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"sync"

	"webscraper-go/web-scraping/domain"
)

func (u *WebScrapingFuncUseCase) ExtractSearchResults() (bool, error) {

	topics, err := u.TopicsRepository.GetTopics()
	if err != nil {
		return false, err
	}

	resultsChan := make(chan domain.SearchResult)
	var wg sync.WaitGroup

	for _, topic := range topics {
		wg.Add(1)
		go func(topicTitle string) {
			defer wg.Done()
			u.WebScrapingCollectRepository.CollectSearchResults(topicTitle, resultsChan)
		}(topic.Title)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var results []domain.SearchResult
	for result := range resultsChan {
		results = append(results, domain.SearchResult{
			Title:   result.Title,
			Url:     result.Url,
			Content: result.Content,
		})
	}

	if len(results) == 0 {
		return false, nil
	}

	for _, result := range results {
		fmt.Println("Title: ", result.Title)

		exists, errVerify := u.WebScrapingRepository.VerifyExistsUrl(result.Url)
		if errVerify != nil {
			return false, errVerify
		}
		if !exists {
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
