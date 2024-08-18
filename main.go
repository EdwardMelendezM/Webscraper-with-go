package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
)

var db *sql.DB

// Inicializar la conexión a la base de datos
func initDB() {
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/nombre_base_de_datos")
	if err != nil {
		log.Fatal(err)
	}
}

// Verifica si la URL ya existe en la base de datos
func urlExists(url string) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM web_pages WHERE url = ?)", url).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}
	return exists
}

// Almacena la página en la base de datos si no existe
func savePage(url, title, content string) {
	if !urlExists(url) {
		_, err := db.Exec("INSERT INTO web_pages (url, title, content) VALUES (?, ?, ?)", url, title, content)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Página almacenada: %s\n", title)
	} else {
		fmt.Printf("URL ya existente: %s\n", url)
	}
}

// Scraper concurrente usando goroutines
func scrapePage(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	c := colly.NewCollector()

	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := e.Text
		content := e.DOM.Text() // Extraer todo el texto de la página
		savePage(url, title, content)
	})

	c.Visit(url)
}

func main() {
	//initDB()
	//defer db.Close()
	//
	//var wg sync.WaitGroup
	//urls := []string{
	//	"https://example.com/page1",
	//	"https://example.com/page2",
	//}
	//
	//for _, url := range urls {
	//	wg.Add(1)
	//	go scrapePage(url, &wg)
	//}
	//
	//wg.Wait()
	//url := "https://kidshealth.org/es/teens/cyberbullying.html"
	//titles := make([]string, 0)
	//c := colly.NewCollector()
	//c.OnHTML("title", func(e *colly.HTMLElement) {
	//	title := e.Text
	//	titles = append(titles, title)
	//	//content := e.DOM.Text() // Extraer todo el texto de la página
	//	fmt.Println("---------------")
	//})
	//c.Visit(url)
	//fmt.Println(titles)
	// instantiate a new collector object
	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com"),
	)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	// triggered when the scraper encounters an error
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	// fired when the server responds
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	// triggered when a CSS selector matches an element
	c.OnHTML("a", func(e *colly.HTMLElement) {
		// printing all URLs associated with the <a> tag on the page
		fmt.Println("%v", e.Attr("href"))
	})

	// triggered once scraping is done (e.g., write the data to a CSV file)
	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})
	fmt.Println("Scraping completado.")
}
