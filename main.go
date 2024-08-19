package main

import (
	"database/sql"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	_ "github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

type SearchResult struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Content string `json:"content"`
}

func main() {
	topics := []string{
		"historias cortas de acoso",
		//"relatos de víctimas de acoso",
		//"historias de acoso laboral",
		//"testimonios de acoso escolar",
		//"experiencias de acoso sexual",
		//"historias de acoso en redes sociales",
		//"foros de víctimas de acoso",
		//"blogs sobre acoso",
		//"artículos periodísticos sobre acoso",
		//"casos reales de acoso",
		//"testimonios de acoso en línea",
		//"historias de acoso psicológico",
		//"narraciones de acoso entre compañeros",
		//"relatos de acoso en el trabajo",
		//"experiencias personales de acoso",
		//"casos de acoso documentados",
		//"historias de bullying en escuelas",
		//"testimonios de acoso en universidades",
		//"experiencias de acoso en el transporte público",
		//"historias de acoso cibernético",
		//"relatos de hostigamiento sexual",
		//"crónicas de acoso por internet",
		//"testimonios de víctimas de stalking",
		//"historias sobre acoso emocional",
		//"experiencias de acoso entre adolescentes",
		//"casos de acoso en la calle",
		//"narraciones de acoso en el deporte",
		//"relatos de acoso en comunidades virtuales",
		//"testimonios de acoso en relaciones de pareja",
		//"historias sobre acoso en centros educativos",
		//"historias de acoso a menores de edad",
		//"relatos de acoso a adultos mayores",
		//"testimonios de acoso en barrios",
		//"experiencias de acoso en plazas públicas",
		//"historias de acoso en centros comerciales",
		//"relatos de acoso en gimnasios",
		//"testimonios de acoso en academias",
		//"experiencias de acoso en universidades",
		//"historias de acoso en colegios",
		//"testimonios de acoso a mujeres adultas",
		//"historias de acoso en centros de trabajo",
		//"casos de acoso en áreas recreativas",
		//"experiencias de acoso en lugares públicos",
		//"relatos de acoso entre adultos en espacios laborales",
		//"historias de acoso en centros culturales",
		//"testimonios de acoso en instituciones educativas",
		//"historias de acoso entre adolescentes en colegios",
		//"relatos de acoso en espacios deportivos",
		//"testimonios de acoso en la comunidad",
		//"historias de acoso en parques",
		//"casos de acoso en gimnasios y centros de fitness",
	}

	var wg sync.WaitGroup
	resultsChan := make(chan SearchResult)

	// Loop through each topic and launch a goroutine for each search
	for _, topic := range topics {
		wg.Add(1)
		go func(topic string) {
			defer wg.Done()
			collectSearchResults(topic, resultsChan)
		}(topic)
	}

	// Espera a que todas las goroutines terminen
	wg.Wait()
	close(resultsChan)
	var results []SearchResult
	for result := range resultsChan {
		results = append(results, SearchResult{
			Title:   result.Title,
			URL:     result.URL,
			Content: result.Content,
		})
	}
	fmt.Printf("Scraped %d results\n", len(results))

	// Insert data into MySQL database
	dsn := "root:secret@tcp(127.0.0.1:3309)/acosoDB"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert data into MySQL database
	for result := range resultsChan {
		var exists bool
		err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM scraped_results WHERE url = ?)",
			result.URL).Scan(&exists)
		if err != nil {
			log.Fatal(err)
		}
		if exists == false {
			var lastNumber *int
			err = db.QueryRow("SELECT MAX(number) FROM scraped_results").Scan(&lastNumber)
			if err != nil {
				log.Fatal(err)
			}
			if lastNumber == nil {
				lastNumber = new(int)
				*lastNumber = 0
			}
			id := uuid.New()
			_, err = db.Exec("INSERT INTO scraped_results (id, title, url, number) VALUES (?, ?, ?, ?)",
				id, result.Title, result.URL, *lastNumber+1)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Inserted result with ID: %s\n", id)
		}
	}
	fmt.Printf("Finished scraping %d topics\n", len(topics))
}

// Función para realizar el scraping de resultados de búsqueda y enviar al canal
func collectSearchResults(topic string, resultsChan chan SearchResult) {
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)
	c := colly.NewCollector(
		colly.AllowedDomains("www.google.com", "google.com"),
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 11.0; Win64; x64) AppleWebKit/537.36 (KHTML, como Gecko) Chrome/116.0.0.0 Safari/537.36"),
	)

	// Configura un tiempo máximo de espera para las solicitudes
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
			resultsChan <- SearchResult{
				Title:   cleanText(title),
				URL:     cleanURL(url),
				Content: cleanText(content),
			}
			fmt.Println("New result: ", title)
		}
	})

	joinedTopic := strings.Join(strings.Fields(topic), "+")
	searchURL := "https://www.google.com/search?q=" + joinedTopic
	c.Visit(searchURL)

}

// Helper function to clean up text
func cleanText(s string) string {
	return strings.TrimSpace(s)
}

// Helper function to clean up URLs (remove unnecessary Google redirect parts)
func cleanURL(s string) string {
	return strings.TrimPrefix(s, "/url?q=")
}
