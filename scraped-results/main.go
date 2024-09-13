package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo/options"
	"webscraper-go/scraped-results/domain"

	"github.com/EdwardMelendezM/api-info-shared/config"
	"github.com/EdwardMelendezM/api-info-shared/db"
	"go.mongodb.org/mongo-driver/mongo"

	ScrapedResultsRepository "webscraper-go/scraped-results/infrastructure/persistence/mysql"
)

func main() {
	cfg := config.Configuration{
		ServerPort:  os.Getenv("SERVER_PORT"),
		StoragePath: os.Getenv("STORAGE_PATH"),
		DB: config.DB{
			DbDatabase: os.Getenv("DB_DATABASE"),
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbUsername: os.Getenv("DB_USERNAME"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
	}

	err := db.InitClients(cfg)
	if err != nil {
		return
	}
	defer func(Client *sql.DB) {
		errClient := Client.Close()
		if errClient != nil {
			fmt.Printf("Error db: %v", errClient)
		}
	}(db.Client)

	scrapedResultRepository := ScrapedResultsRepository.NewScrapedResultRepository()
	scrapedResults, errGetScrapedResult := scrapedResultRepository.GetScrapedResults("91da2ca7-6244-11ef-9d2f-0242ac110002")
	if errGetScrapedResult != nil {
		fmt.Printf("Error: %v", err)
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	for _, scrapedResult := range scrapedResults {
		result, errGetCleanCorpus := getCleanCorpus(scrapedResult.Content)
		if errGetCleanCorpus != nil {
			log.Fatal(errGetCleanCorpus)
		}

		wordToCount := map[string]int{
			"mensajes":       0,
			"redes sociales": 0,
			"chat":           0,
			"video":          0,
			"correo":         0,
			"foros":          0,
			"facebook":       0,
			"instagram":      0,
			"snapchat":       0,
			"whatsapp":       0,
			"twitter":        0,
			"youtube":        0,
			"tiktok":         0,
			"linkedin":       0,
			"blog":           0,
			"email":          0,
			"primaria":       0,
			"secundaria":     0,
			"facultad":       0,
			"bachillerato":   0,
			"universidad":    0,
			"preparatoria":   0,
			"colegio":        0,
			"instituto":      0,
			"media":          0,
			"academia":       0,
			"institucion":    0,
			"acosador":       0,
			"victima":        0,
			"perpetrador":    0,
			"companeros":     0,
			"agresor":        0,
			"testigos":       0,
			"espia":          0,
			"maton":          0,
			"grupo":          0,
			"bully":          0,
			"supervisor":     0,
			"adolescente":    0,
			"joven":          0,
			"niño":           0,
			"infantil":       0,
			"constante":      0,
			"frecuente":      0,
			"persistente":    0,
			"reiterado":      0,
			"ocasional":      0,
			"repetitivo":     0,
			"periodico":      0,
		}

		for word, _ := range wordToCount {
			wordToCount[word] = CountWord(result, word)
		}

		createNewSemanticOntologyCountResult := domain.SemanticOntologyCountResult{
			ProjectID:     scrapedResult.ProjectId,
			Title:         scrapedResult.Title,
			URL:           scrapedResult.Url,
			Content:       scrapedResult.Content,
			Number:        scrapedResult.Number,
			Mensajes:      float64(wordToCount["mensajes"]),
			RedesSociales: float64(wordToCount["redes sociales"]),
			Chat:          float64(wordToCount["chat"]),
			Video:         float64(wordToCount["video"]),
			Correo:        float64(wordToCount["correo"]),
			Foros:         float64(wordToCount["foros"]),
			Facebook:      float64(wordToCount["facebook"]),
			Instagram:     float64(wordToCount["instagram"]),
			SnapChat:      float64(wordToCount["snapchat"]),
			WhatsApp:      float64(wordToCount["whatsapp"]),
			Twitter:       float64(wordToCount["twitter"]),
			YouTube:       float64(wordToCount["youtube"]),
			TikTok:        float64(wordToCount["tiktok"]),
			LinkedIn:      float64(wordToCount["linkedin"]),
			Blog:          float64(wordToCount["blog"]),
			Email:         float64(wordToCount["email"]),
			Primaria:      float64(wordToCount["primaria"]),
			Secundaria:    float64(wordToCount["secundaria"]),
			Facultad:      float64(wordToCount["facultad"]),
			Bachillerato:  float64(wordToCount["bachillerato"]),
			Universidad:   float64(wordToCount["universidad"]),
			Preparatoria:  float64(wordToCount["preparatoria"]),
			Colegio:       float64(wordToCount["colegio"]),
			Instituto:     float64(wordToCount["instituto"]),
			Media:         float64(wordToCount["media"]),
			Academia:      float64(wordToCount["academia"]),
			Institucion:   float64(wordToCount["institucion"]),
			Acosador:      float64(wordToCount["acosador"]),
			Victima:       float64(wordToCount["victima"]),
			Perpetrador:   float64(wordToCount["perpetrador"]),
			Companeros:    float64(wordToCount["companeros"]),
			Agresor:       float64(wordToCount["agresor"]),
			Testigos:      float64(wordToCount["testigos"]),
			Espia:         float64(wordToCount["espia"]),
			Maton:         float64(wordToCount["maton"]),
			Grupo:         float64(wordToCount["grupo"]),
			Bully:         float64(wordToCount["bully"]),
			Supervisor:    float64(wordToCount["supervisor"]),
			Adolescente:   float64(wordToCount["adolescente"]),
			Joven:         float64(wordToCount["joven"]),
			Niño:          float64(wordToCount["niño"]),
			Infantil:      float64(wordToCount["infantil"]),
			Constante:     float64(wordToCount["constante"]),
			Frecuente:     float64(wordToCount["frecuente"]),
			Persistente:   float64(wordToCount["persistente"]),
			Reiterado:     float64(wordToCount["reiterado"]),
			Ocasional:     float64(wordToCount["ocasional"]),
			Repetitivo:    float64(wordToCount["repetitivo"]),
			Periodico:     float64(wordToCount["periodico"]),
			DeletedAt:     false,
		}
		res, errInsert := insertDocument(client, createNewSemanticOntologyCountResult)
		if errInsert != nil {
			log.Fatal(errInsert)
		}
		fmt.Printf("Inserted document with ID %v\n", res.InsertedID)
	}

}

func insertDocument(client *mongo.Client, result domain.SemanticOntologyCountResult) (*mongo.InsertOneResult, error) {
	collection := client.Database("acosoDBMongo").Collection("semantic_ontology_count_result")

	ctx := context.TODO()
	res, err := collection.InsertOne(ctx, result)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func getCleanCorpus(htmlContent string) (string, error) {
	// Crear el cuerpo de la solicitud
	requestBody := RequestBody{
		Content: htmlContent,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error al convertir el cuerpo a JSON: %v", err)
	}

	url := "http://localhost:5000/clean-corpus" // URL del endpoint de Flask
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error al hacer la solicitud POST: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error al leer la respuesta: %v", err)
	}

	var responseBody ResponseBody
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return "", fmt.Errorf("error al decodificar la respuesta: %v", err)
	}

	return responseBody.Corpus, nil
}

func CountWord(corpus string, word string) int {
	count := 0
	words := strings.Fields(corpus)
	for _, w := range words {
		if w == word {
			count++
		}
	}
	return count
}

type RequestBody struct {
	Content string `json:"content"`
}

type ResponseBody struct {
	Corpus string `json:"corpus"`
}
