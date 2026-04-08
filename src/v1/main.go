package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	TopicsRepository "webscraper-go/topics/infrastructure/persistence/mysql"
	WebScrapingRepository "webscraper-go/web-scraping/infrastructure/persistence/mysql"

	WebScrapingCollectRepository "webscraper-go/web-scraping/infrastructure/scraping/collect"
	webScraperUseCase "webscraper-go/web-scraping/usecase"

	"github.com/EdwardMelendezM/api-info-shared/config"
	"github.com/EdwardMelendezM/api-info-shared/db"

	// El driver de pgx debe estar presente para que sql.Open lo encuentre
	_ "github.com/jackc/pgx/v5/stdlib"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	// 1. Mapeo de configuración usando las llaves de tu .env
	cfg := config.Configuration{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		StoragePath: getEnv("STORAGE_PATH", "./storage"),
		DB: config.DB{
			// Usamos los nombres exactos de tu archivo .env
			DbDatabase: getEnv("POSTGRES_DB", "acoso-db"),
			DbHost:     getEnv("POSTGRES_HOST", "127.0.0.1"),
			DbPort:     getEnv("POSTGRES_PORT", "5111"),
			DbUsername: getEnv("POSTGRES_USER", "postgres"),
			DbPassword: getEnv("POSTGRES_PASSWORD", "secret"),
		},
	}

	// 2. Inicializar la conexión
	err := db.InitClients(cfg)
	if err != nil {
		log.Fatalf("No se pudo iniciar la DB: %v", err)
	}

	// 3. Cerrar la conexión al terminar (Usa ClientV2 que definimos antes)
	defer func(client *sql.DB) {
		if client != nil {
			if err := client.Close(); err != nil {
				fmt.Printf("Error cerrando db: %v\n", err)
			}
		}
	}(db.ClientV2)

	// 4. Inyección de dependencias
	// Nota: Asegúrate de que estos constructores usen db.ClientV2 internamente
	topicsRepository := TopicsRepository.NewTopicsRepository()
	webScrapingRepository := WebScrapingRepository.NewWebScrapingRepository()
	webScrapingCollectRepository := WebScrapingCollectRepository.NewWebScrapingCollectRepository()

	instance := webScraperUseCase.NewWebScrapingFuncUseCase(
		webScrapingRepository,
		webScrapingCollectRepository,
		topicsRepository,
	)

	// 5. Ejecución del caso de uso
	fmt.Println("Iniciando proceso de scraping...")
	value, errExtract := instance.ExtractSearchResults()
	if errExtract != nil {
		fmt.Printf("Error durante el proceso: %v\n", errExtract)
	}

	if value {
		fmt.Println("Scraping exitoso.")
	} else {
		fmt.Println("Scraping no exitoso.")
	}
	fmt.Println("Proceso finalizado.")
}
