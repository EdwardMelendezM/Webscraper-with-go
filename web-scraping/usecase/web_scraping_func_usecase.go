package usecase

import (
	"golang.org/x/text/encoding/charmap"
	"strings"

	"github.com/PuerkitoBio/goquery"
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
		}
	}

	// 2. Verify repeated results
	for _, result := range results {
		if _, ok := seenMap[result.Url]; ok {
			continue
		}

		if strings.Contains(result.Url, "pdf") {
			continue
		}

		seenMap[result.Url] = domain.NewRecordWebScraping{
			Title:   result.Title,
			Url:     result.Url,
			Content: result.Content,
		}
	}

	// 3. Verify if existing results are in the new results
	for _, result := range seenMap {
		if _, ok := existingMap[result.Url]; !ok {
			notExistingResults[result.Url] = domain.NewRecordWebScraping{
				Title:   result.Title,
				Url:     result.Url,
				Content: result.Content,
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
	for _, notExistingResult := range notExistingResults {
		id := uuid.New().String()
		*lastNumber = *lastNumber + 1
		contentCleaned := extractText(notExistingResult.Content)
		contentUtf8, errUt8 := convertToUTF8(contentCleaned)
		if errUt8 != nil {
			break
		}
		body := domain.CreateRecordWebScraping{
			Title:   notExistingResult.Title,
			Url:     notExistingResult.Url,
			Content: contentUtf8,
			Number:  *lastNumber,
		}
		_, errCreateNewRecord := u.WebScrapingRepository.CreateRecord(id, projectId, body)
		if errCreateNewRecord != nil {
			break
		}
	}
	return true, nil
}

func extractText(htmlContent string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		panic(err)
	}
	return doc.Text()
}

func convertToUTF8(input string) (string, error) {
	decoder := charmap.ISO8859_1.NewDecoder()
	output, err := decoder.String(input)
	if err != nil {
		return "", err
	}
	return output, nil
}
