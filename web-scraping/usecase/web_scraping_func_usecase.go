package usecase

import (
	"sync"

	"webscraper-go/web-scraping/domain"
)

func (u *WebScrapingFuncUseCase) ExtractSearchResults() (bool, error) {
	projectId := "91da2ca7-6244-11ef-9d2f-0242ac110002"
	topics, err := u.TopicsRepository.GetTopics(projectId)
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
			Path:    result.Path,
		})
	}

	if len(results) == 0 {
		return false, nil
	}

	existingResults, err := u.WebScrapingRepository.GetRecordResult(projectId, 1000)

	existingMap := make(map[string]domain.NewRecordWebScraping)
	notExistingResults := make(map[string]domain.NewRecordWebScraping)
	for _, existingResult := range existingResults {
		existingMap[existingResult.Url] = domain.NewRecordWebScraping{
			Title: existingResult.Title,
			Url:   existingResult.Url,
			Path:  existingResult.Path,
		}
	}
	for _, result := range results {
		if _, ok := existingMap[result.Url]; !ok {
			notExistingResults[result.Url] = domain.NewRecordWebScraping{
				Title: result.Title,
				Url:   result.Url,
				Path:  result.Path,
			}
			delete(existingMap, result.Url)
		}
	}

	//Step 1: Delete file in the path existingMap
	//Step 2: Get content of files in the path notExistingResults
	//Step 3: Add new record notExistingResults

	//for _, existingResult := range results {
	//	fmt.Println("Title: ", existingResult.Title)
	//
	//	exists, errVerify := u.WebScrapingRepository.VerifyExistsUrl(existingResult.Url)
	//	if errVerify != nil {
	//		return false, errVerify
	//	}
	//	if !exists {
	//		lastNumber, errLastNumber := u.WebScrapingRepository.GetLastNumber()
	//		if errLastNumber != nil {
	//			return false, errLastNumber
	//		}
	//		if lastNumber == nil {
	//			lastNumber = new(int)
	//			*lastNumber = 0
	//		}
	//		id := uuid.New().String()
	//		_, errCreateNewRecord := u.WebScrapingRepository.CreateRecord(id, domain.CreateRecordWebScraping{
	//			Title:  existingResult.Title,
	//			Url:    existingResult.Url,
	//			Number: *lastNumber + 1,
	//		})
	//		if errCreateNewRecord != nil {
	//			return false, errLastNumber
	//		}
	//	}
	//}

	return true, nil
}
