package main

import (
	"fmt"
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
)

func extractContent(link string) {
	// Iniciar Playwright
	//err := playwright.Install()
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
	defer pw.Stop()

	// Lanzar el navegador Chromium
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch Chromium: %v", err)
	}
	defer browser.Close()

	// Crear una nueva página
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	// Navegar a la URL
	_, err = page.Goto(link)
	if err != nil {
		log.Fatalf("could not go to page: %v", err)
	}

	// Esperar a que la página esté completamente cargada
	time.Sleep(5 * time.Second) // Puedes ajustar el tiempo o usar `WaitForSelector()` para esperar elementos específicos.

	// Extraer el contenido de la página
	content, err := page.Content()
	if err != nil {
		log.Fatalf("could not get content: %v", err)
	}

	fmt.Printf("Contenido de la página: %s\n", content)

	// Cerrar la página
	page.Close()
}

func main() {
	// URL que deseas visitar
	link := "https://www.unicef.org/es/end-violence/ciberacoso-que-es-y-como-detenerlo"
	extractContent(link)
}
