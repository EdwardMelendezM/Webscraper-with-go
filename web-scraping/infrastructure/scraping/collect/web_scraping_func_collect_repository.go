package collect

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocolly/colly"
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

		if title != "" && url != "" {
			pageContent, path := downloadPage(c, url, title)
			resultsChan <- domain.SearchResult{
				Title:   cleanText(title),
				Url:     cleanURL(url),
				Content: pageContent,
				Path:    path,
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

func downloadPage(c *colly.Collector, url string, title string) (string, string) {
	var pageContent string
	var path string

	// Create a new collector for downloading the content
	pageCollector := c.Clone()

	pageCollector.OnResponse(func(r *colly.Response) {
		pageContent = string(r.Body)

		// Create a directory for storing the downloaded pages
		dir := "downloaded_pages"
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}

		// Save the content to a file
		filename := filepath.Join(dir, sanitizeFilename(title)+".html")
		err := ioutil.WriteFile(filename, r.Body, 0644)
		if err != nil {
			fmt.Println("Error saving file:", err)
			return
		}

		path = filename
	})

	pageCollector.Visit(url)

	return pageContent, path
}

// Helper function to sanitize filenames
func sanitizeFilename(name string) string {
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(`<>:"/\|?*`, r) || r == ' ' {
			return '_'
		}
		return r
	}, name)
}

// Helper function to clean up text
func cleanText(s string) string {
	return strings.TrimSpace(s)
}

// Helper function to clean up URLs (remove unnecessary Google redirect parts)
func cleanURL(s string) string {
	return strings.TrimPrefix(s, "/url?q=")
}
