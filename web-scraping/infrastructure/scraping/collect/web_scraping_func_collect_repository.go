package collect

import (
	"fmt"
	"github.com/gocolly/colly"
	"math/rand"
	"strings"
	"time"

	"webscraper-go/web-scraping/domain"
)

func (r WebScrapingCollectRepository) CollectSearchResults(
	topic string,
	resultsChan chan<- domain.SearchResult,
) {
	time.Sleep(time.Duration(2+rand.Intn(10)) * time.Second)
	c := colly.NewCollector(
		colly.AllowedDomains("www.google.com", "google.com"),
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 11.0; Win64; x64) AppleWebKit/537.36 (KHTML, como Gecko) Chrome/116.0.0.0 Safari/537.36"),
	)

	// Set a timeout to avoid infinite waiting
	c.SetRequestTimeout(5 * time.Second)

	// Set custom headers to simulate a real browser
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "es-ES,es;q=0.9,en;q=0.8")
		fmt.Printf("Visiting %s\n", r.URL)
	})

	// Handle errors during scraping
	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error: %v\n", e)
	})

	// Extract search results
	c.OnHTML("div.g", func(e *colly.HTMLElement) {
		title := e.ChildText("h3")
		url := e.ChildAttr("a", "href")
		content := e.ChildText("span.aCOpRe")

		if title != "" && url != "" {
			resultsChan <- domain.SearchResult{
				Title:   cleanText(title),
				Url:     cleanURL(url),
				Content: cleanText(content),
				Path:    "",
			}
			fmt.Println("New result: ", title)
		}
	})

	joinedTopic := strings.Join(strings.Fields(topic), "+")
	searchURL := "https://www.google.com/search?q=" + joinedTopic
	err := c.Visit(searchURL)
	if err != nil {
		fmt.Println("Error visiting URL: ", searchURL)
	}

}

// Helper function to clean up text
func cleanText(s string) string {
	return strings.TrimSpace(s)
}

// Helper function to clean up URLs (remove unnecessary Google redirect parts)
func cleanURL(s string) string {
	return strings.TrimPrefix(s, "/url?q=")
}
