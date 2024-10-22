package collect

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gocolly/colly"

	"webscraper-go/web-scraping/domain"
)

func (r WebScrapingCollectRepository) CollectSearchResultsAndReturn(
	topic string,
) (results []domain.SearchResult,
	err error) {
	c := colly.NewCollector(
		//colly.AllowedDomains("www.google.com", "google.com", "www.search.brave.com", "search.brave.com"),
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

	// Manejo para Brave
	c.OnHTML("div.snippet", func(e *colly.HTMLElement) {
		title := e.ChildText(".title")
		url := e.ChildAttr("a", "href")

		if title != "" && url != "" {
			cleanedURL := cleanURL(url)
			client := &http.Client{
				Timeout: 8 * time.Second,
			}
			resp, errGet := client.Get(cleanedURL)
			if errGet != nil {
				fmt.Printf("Error descargando contenido de %s: %v\n", cleanedURL, errGet)
				return
			}
			defer resp.Body.Close()

			body, errRead := ioutil.ReadAll(resp.Body)
			if errRead != nil {
				fmt.Printf("Error leyendo contenido de %s: %v\n", cleanedURL, errRead)
				return
			}

			newResult := domain.SearchResult{
				Title:   cleanText(title),
				Url:     cleanedURL,
				Content: cleanText(string(body)),
				Path:    "", // Puedes modificar este valor si lo necesitas
			}

			results = append(results, newResult)
		}
	})

	// Extract search results
	c.OnHTML("div.g", func(e *colly.HTMLElement) {
		title := e.ChildText("h3")
		url := e.ChildAttr("a", "href")

		if title != "" && url != "" {
			cleanedURL := cleanURL(url)

			client := &http.Client{
				Timeout: 8 * time.Second,
			}
			// Fetch the content from the URL
			resp, errClient := client.Get(cleanedURL)
			if errClient != nil {
				fmt.Printf("Error downloading content from %s: %v\n", cleanedURL, errClient)
				return
			}
			defer resp.Body.Close()

			body, errBody := ioutil.ReadAll(resp.Body)
			if errBody != nil {
				fmt.Printf("Error reading content from %s: %v\n", cleanedURL, errBody)
				return
			}

			newResult := domain.SearchResult{
				Title:   cleanText(title),
				Url:     cleanedURL,
				Content: cleanText(string(body)),
				Path:    "", // You can modify this if needed
			}

			results = append(results, newResult)
		}
	})

	c.OnHTML("li.b_algo", func(e *colly.HTMLElement) {
		// Extraer el enlace desde el elemento "a" con la clase "tilk b_lario"
		url := e.ChildAttr("a.tilk.b_lario", "href")
		// Extraer el título desde el elemento "h2 > a"
		title := e.ChildText("h2 > a")

		if title != "" && url != "" {
			cleanedURL := cleanURL(url)

			client := &http.Client{
				Timeout: 8 * time.Second,
			}
			// Fetch the content from the URL
			resp, errClient := client.Get(cleanedURL)
			if errClient != nil {
				fmt.Printf("Error downloading content from %s: %v\n", cleanedURL, errClient)
				return
			}
			defer resp.Body.Close()

			body, errBody := ioutil.ReadAll(resp.Body)
			if errBody != nil {
				fmt.Printf("Error reading content from %s: %v\n", cleanedURL, errBody)
				return
			}

			newResult := domain.SearchResult{
				Title:   cleanText(title),
				Url:     cleanedURL,
				Content: cleanText(string(body)),
				Path:    "", // Puedes modificar este valor si lo necesitas
			}

			results = append(results, newResult)
		}
	})

	joinedTopic := strings.Join(strings.Fields(topic), "+")
	//searchURL := "https://www.google.com/search?q=" + joinedTopic
	searchURL := "https://search.brave.com/search?q=" + joinedTopic + "&source=desktop"
	//searchURL := "https://www.bing.com/search?q=" + joinedTopic
	err = c.Visit(searchURL)
	if err != nil {
		fmt.Println("Error visiting URL: ", searchURL)
	}
	return results, err
}

func (r WebScrapingCollectRepository) CollectSearchResults(
	topic string,
	results *[]domain.SearchResult,
) {
	c := colly.NewCollector(
		//colly.AllowedDomains("www.google.com", "google.com", "www.search.brave.com", "search.brave.com"),
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

	// Manejo para Brave
	c.OnHTML("div.snippet", func(e *colly.HTMLElement) {
		title := e.ChildText(".title")
		url := e.ChildAttr("a", "href")

		if title != "" && url != "" {
			cleanedURL := cleanURL(url)
			client := &http.Client{
				Timeout: 8 * time.Second,
			}
			resp, err := client.Get(cleanedURL)
			if err != nil {
				fmt.Printf("Error descargando contenido de %s: %v\n", cleanedURL, err)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Error leyendo contenido de %s: %v\n", cleanedURL, err)
				return
			}

			newResult := domain.SearchResult{
				Title:   cleanText(title),
				Url:     cleanedURL,
				Content: cleanText(string(body)),
				Path:    "", // Puedes modificar este valor si lo necesitas
			}

			*results = append(*results, newResult)
		}
	})

	// Extract search results
	c.OnHTML("div.g", func(e *colly.HTMLElement) {
		title := e.ChildText("h3")
		url := e.ChildAttr("a", "href")

		if title != "" && url != "" {
			cleanedURL := cleanURL(url)

			client := &http.Client{
				Timeout: 8 * time.Second,
			}
			// Fetch the content from the URL
			resp, err := client.Get(cleanedURL)
			if err != nil {
				fmt.Printf("Error downloading content from %s: %v\n", cleanedURL, err)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Error reading content from %s: %v\n", cleanedURL, err)
				return
			}

			newResult := domain.SearchResult{
				Title:   cleanText(title),
				Url:     cleanedURL,
				Content: cleanText(string(body)),
				Path:    "", // You can modify this if needed
			}

			*results = append(*results, newResult)
		}
	})

	joinedTopic := strings.Join(strings.Fields(topic), "+")
	//searchURL := "https://www.google.com/search?q=" + joinedTopic
	searchURL := "https://search.brave.com/search?q=" + joinedTopic
	err := c.Visit(searchURL)
	if err != nil {
		fmt.Println("Error visiting URL: ", searchURL)
	}
}

func extractContent(link string) {
	contentCollector := colly.NewCollector()

	// User-Agent para evitar bloqueos
	contentCollector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	})

	// Extraer el contenido de la página
	contentCollector.OnHTML("body", func(e *colly.HTMLElement) {
		content := e.Text
		fmt.Printf("Contenido de la página: %s\n", content)
	})

	err := contentCollector.Visit(link)
	if err != nil {
		log.Printf("Error al visitar: %s", link)
	}

	// Pausa para evitar bloqueos
	time.Sleep(2 * time.Second)
}

// Helper function to clean up text
func cleanText(s string) string {
	return strings.TrimSpace(s)
}

// Helper function to clean up URLs (remove unnecessary Google redirect parts)
func cleanURL(s string) string {
	return strings.TrimPrefix(s, "/url?q=")
}
