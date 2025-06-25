package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"strings"

	"webscraper-go/web-scraping/domain"
)

func (u *WebScrapingFuncUseCase) ExtractSearchResults() (bool, error) {
	projectId := "91da2ca7-6244-11ef-9d2f-0242ac110002"
	topics, err := u.TopicsRepository.GetTopics(projectId)
	if err != nil {
		return false, err
	}

	//results := make([]domain.SearchResult, 0)

	for index, topic := range topics {
		fmt.Printf(":=> Number: %d\n", index+1)
		//u.WebScrapingCollectRepository.CollectSearchResults(topic.Title, &results)
		results, errReturn := u.WebScrapingCollectRepository.CollectSearchResultsAndReturn(topic.Title)
		if errReturn != nil {
			return false, errReturn
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

			if strings.Contains(result.Content, "403") {
				continue
			}

			if strings.Contains(result.Content, "Just a moment...") {
				continue
			}

			if strings.Contains(result.Content, "Oops, something went wrong") {
				continue
			}

			if strings.Contains(result.Content, "Sorry, you have been blocked") {
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

		fmt.Printf("=== Start inserting new records ===\n")
		// 5: Add new record notExistingResults
		for _, notExistingResult := range notExistingResults {
			if strings.TrimSpace(notExistingResult.Content) == "" || notExistingResult.Title == "" {
				continue
			}
			id := uuid.New().String()
			*lastNumber = *lastNumber + 1

			// Limpiar y procesar contenido
			contentCleaned, errClean := extractText(notExistingResult.Content)
			if errClean != nil {
				break
			}

			// Procesar título y contenido con PLN
			titleTokens := tokenize(notExistingResult.Title)
			contentTokens := tokenize(contentCleaned)

			// Obtener la palabra clave más relevante
			documents := [][]string{titleTokens, contentTokens}
			wordKey, errGetKeyWord := getKeyword(contentCleaned, documents)
			if errGetKeyWord != nil {
				break
			}

			body := domain.CreateRecordWebScraping{
				Title:         notExistingResult.Title,
				Url:           notExistingResult.Url,
				Content:       contentCleaned,
				Number:        *lastNumber,
				TitleCorpus:   strings.Join(titleTokens, ","),
				ContentCorpus: strings.Join(contentTokens, ","),
				WordKey:       wordKey,
			}
			_, errCreateNewRecord := u.WebScrapingRepository.CreateRecord(id, projectId, body)
			if errCreateNewRecord != nil {
				break
			}
			fmt.Printf("Number inserted: %d\n", *lastNumber)
		}

	}

	//if len(results) == 0 {
	//	return false, nil
	//}
	//
	//existingResults, errResult := u.WebScrapingRepository.GetRecordResult(projectId, 1000)
	//if errResult != nil {
	//	return false, errResult
	//}
	//
	//existingMap := make(map[string]domain.NewRecordWebScraping)
	//notExistingResults := make(map[string]domain.NewRecordWebScraping)
	//seenMap := make(map[string]domain.NewRecordWebScraping)
	//
	//// 1. Map existing results
	//for _, existingResult := range existingResults {
	//	existingMap[existingResult.Url] = domain.NewRecordWebScraping{
	//		Title: existingResult.Title,
	//		Url:   existingResult.Url,
	//	}
	//}
	//
	//// 2. Verify repeated results
	//for _, result := range results {
	//	if _, ok := seenMap[result.Url]; ok {
	//		continue
	//	}
	//
	//	if strings.Contains(result.Url, "pdf") {
	//		continue
	//	}
	//
	//	if strings.Contains(result.Content, "PDF") {
	//		continue
	//	}
	//
	//	if strings.Contains(result.Url, "youtube") {
	//		continue
	//	}
	//
	//	if strings.Contains(result.Content, "404") {
	//		continue
	//	}
	//
	//	if strings.Contains(result.Content, "403") {
	//		continue
	//	}
	//
	//	if strings.Contains(result.Content, "Just a moment...") {
	//		continue
	//	}
	//
	//	if strings.Contains(result.Content, "Oops, something went wrong") {
	//		continue
	//	}
	//
	//	if strings.Contains(result.Content, "Sorry, you have been blocked") {
	//		continue
	//	}
	//
	//	seenMap[result.Url] = domain.NewRecordWebScraping{
	//		Title:   result.Title,
	//		Url:     result.Url,
	//		Content: result.Content,
	//	}
	//}
	//
	//// 3. Verify if existing results are in the new results
	//for _, result := range seenMap {
	//	if _, ok := existingMap[result.Url]; !ok {
	//		notExistingResults[result.Url] = domain.NewRecordWebScraping{
	//			Title:   result.Title,
	//			Url:     result.Url,
	//			Content: result.Content,
	//		}
	//		delete(existingMap, result.Url)
	//	}
	//}
	//
	//// 4: Get last number
	//lastNumber, errLastNumber := u.WebScrapingRepository.GetLastNumber(projectId)
	//if errLastNumber != nil {
	//	return false, errLastNumber
	//}
	//
	//if lastNumber == nil {
	//	lastNumber = new(int)
	//	*lastNumber = 0
	//}
	//
	//fmt.Printf("=== Start inserting new records ===\n")
	//// 5: Add new record notExistingResults
	//for _, notExistingResult := range notExistingResults {
	//	if notExistingResult.Content == "" || notExistingResult.Title == "" {
	//		continue
	//	}
	//	id := uuid.New().String()
	//	*lastNumber = *lastNumber + 1
	//
	//	// Limpiar y procesar contenido
	//	contentCleaned, errClean := extractText(notExistingResult.Content)
	//	if errClean != nil {
	//		break
	//	}
	//
	//	// Procesar título y contenido con PLN
	//	titleTokens := tokenize(notExistingResult.Title)
	//	contentTokens := tokenize(contentCleaned)
	//
	//	// Obtener la palabra clave más relevante
	//	documents := [][]string{titleTokens, contentTokens}
	//	wordKey, errGetKeyWord := getKeyword(contentCleaned, documents)
	//	if errGetKeyWord != nil {
	//		break
	//	}
	//
	//	body := domain.CreateRecordWebScraping{
	//		Title:         notExistingResult.Title,
	//		Url:           notExistingResult.Url,
	//		Content:       contentCleaned,
	//		Number:        *lastNumber,
	//		TitleCorpus:   strings.Join(titleTokens, ","),
	//		ContentCorpus: strings.Join(contentTokens, ","),
	//		WordKey:       wordKey,
	//	}
	//	_, errCreateNewRecord := u.WebScrapingRepository.CreateRecord(id, projectId, body)
	//	if errCreateNewRecord != nil {
	//		break
	//	}
	//	fmt.Printf("Number inserted: %d\n", *lastNumber)
	//}
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
	cleanText2 := cleanText(textWithoutTags)
	reSpaces := regexp.MustCompile(`\s+`)
	textWithSingleSpaces := reSpaces.ReplaceAllString(cleanText2, " ")

	// 4. Remove leading and trailing whitespace
	cleanedText := strings.TrimSpace(textWithSingleSpaces)

	// 5. Final cleanup for extra spaces after removing words
	cleanedText = reSpaces.ReplaceAllString(cleanedText, " ")
	cleanedText = strings.TrimSpace(cleanedText)

	return cleanedText, nil
}

// cleanText elimina secciones del texto si alguna línea en la sección tiene menos de 4 palabras.
func cleanText(text string) string {
	// Expresión regular para dividir el texto en secciones basadas en múltiples saltos de línea
	re := regexp.MustCompile(`(?m)(?:\n\s*){2,}`) // Coincide con dos o más saltos de línea

	// Dividir el texto en secciones usando la expresión regular
	sections := re.Split(text, -1)

	var cleanedSections []string

	for _, section := range sections {
		section = strings.TrimSpace(section) // Eliminar espacios en blanco
		if isValidSection(section) {
			cleanedSections = append(cleanedSections, section)
		}
	}

	// Unir todas las secciones válidas en el texto final
	return strings.Join(cleanedSections, "\n\n")
}

// isValidSection verifica si una sección es válida según la longitud de sus palabras en cada línea
func isValidSection(section string) bool {
	lines := strings.Split(section, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line) // Eliminar espacios en blanco
		if len(strings.Fields(line)) < 4 {
			return false
		}
	}
	return true
}
