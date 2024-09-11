package usecase

import (
	"regexp"
	"strings"
	"unicode/utf8"

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

		if strings.Contains(result.Content, "PDF") {
			continue
		}

		if strings.Contains(result.Url, "youtube") {
			continue
		}

		if strings.Contains(result.Content, "404") {
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

		// Limpiar y procesar contenido
		contentCleaned, errClean := extractText(notExistingResult.Content)
		if errClean != nil {
			break
		}
		contentUtf8, errUt8 := convertToUTF8(contentCleaned)
		if errUt8 != nil {
			break
		}

		// Eliminar espacios en blanco al principio y al final
		cleanStopWords := strings.TrimSpace(stopWords)
		// Dividir la constante en una slice de strings usando salto de línea como delimitador
		stopWordsList := strings.Split(cleanStopWords, "\n")

		// Procesar título y contenido con PLN
		titleTokens := tokenize(notExistingResult.Title, stopWordsList)
		contentTokens := tokenize(contentUtf8, stopWordsList)

		// Obtener la palabra clave más relevante
		documents := [][]string{titleTokens, contentTokens}
		wordKey, err := getKeyword(contentUtf8, documents, stopWordsList)
		if err != nil {
			break
		}

		body := domain.CreateRecordWebScraping{
			Title:         notExistingResult.Title,
			Url:           notExistingResult.Url,
			Content:       contentUtf8,
			Number:        *lastNumber,
			TitleCorpus:   strings.Join(titleTokens, " "),
			ContentCorpus: strings.Join(contentTokens, " "),
			WordKey:       wordKey,
		}
		_, errCreateNewRecord := u.WebScrapingRepository.CreateRecord(id, projectId, body)
		if errCreateNewRecord != nil {
			break
		}
	}
	return true, nil
}

// extractText removes HTML, CSS, JavaScript, and unwanted words from the content.
func extractText(htmlContent string) (string, error) {
	// 1. Remove content inside <script> and <style> tags
	reScript := regexp.MustCompile(`(?s)<script.*?>.*?</script>`)
	htmlWithoutScript := reScript.ReplaceAllString(htmlContent, "")

	reStyle := regexp.MustCompile(`(?s)<style.*?>.*?</style>`)
	htmlWithoutCSS := reStyle.ReplaceAllString(htmlWithoutScript, "")

	// 2. Remove all remaining HTML tags
	reTags := regexp.MustCompile(`<[^>]*>`)
	textWithoutTags := reTags.ReplaceAllString(htmlWithoutCSS, "")

	// 3. Replace multiple whitespace and newline characters with a single space
	reSpaces := regexp.MustCompile(`\s+`)
	textWithSingleSpaces := reSpaces.ReplaceAllString(textWithoutTags, " ")

	// 4. Remove leading and trailing whitespace
	cleanedText := strings.TrimSpace(textWithSingleSpaces)

	// 5. Remove unwanted words
	reRemoveWords := regexp.MustCompile(removeWords)
	cleanedText = reRemoveWords.ReplaceAllString(cleanedText, "")

	// 6. Final cleanup for extra spaces after removing words
	cleanedText = reSpaces.ReplaceAllString(cleanedText, " ")
	cleanedText = strings.TrimSpace(cleanedText)

	return cleanedText, nil
}

func convertToUTF8(input string) (string, error) {
	var output strings.Builder
	for _, r := range input {
		// Verifica si el carácter es válido en UTF-8
		if utf8.ValidRune(r) {
			output.WriteRune(r)
		}
	}
	return output.String(), nil

}
