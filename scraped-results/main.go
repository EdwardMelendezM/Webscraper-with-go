package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math"
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

type RequestBody struct {
	Content string `json:"content"`
}

type ResponseBody struct {
	Corpus string `json:"corpus"`
}

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

		//Split
		corpus := strings.Fields(result)

		wordTf := map[string]float64{
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

		wordToTfIdf := map[string]float64{
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

		for word, _ := range wordTf {
			wordTf[word] = CountWord(corpus, word)
			wordToTfIdf[word] = calculateTfidfForWordInCorpus(word, corpus)
		}

		createNewSemanticOntologyTfIdfResult := domain.SemanticOntologyTfIdfResult{
			ProjectID:     scrapedResult.ProjectId,
			Title:         scrapedResult.Title,
			URL:           scrapedResult.Url,
			Content:       scrapedResult.Content,
			Number:        scrapedResult.Number,
			Mensajes:      wordToTfIdf["mensajes"],
			RedesSociales: wordToTfIdf["redes sociales"],
			Chat:          wordToTfIdf["chat"],
			Video:         wordToTfIdf["video"],
			Correo:        wordToTfIdf["correo"],
			Foros:         wordToTfIdf["foros"],
			Facebook:      wordToTfIdf["facebook"],
			Instagram:     wordToTfIdf["instagram"],
			SnapChat:      wordToTfIdf["snapchat"],
			WhatsApp:      wordToTfIdf["whatsapp"],
			Twitter:       wordToTfIdf["twitter"],
			YouTube:       wordToTfIdf["youtube"],
			TikTok:        wordToTfIdf["tiktok"],
			LinkedIn:      wordToTfIdf["linkedin"],
			Blog:          wordToTfIdf["blog"],
			Email:         wordToTfIdf["email"],
			Primaria:      wordToTfIdf["primaria"],
			Secundaria:    wordToTfIdf["secundaria"],
			Facultad:      wordToTfIdf["facultad"],
			Bachillerato:  wordToTfIdf["bachillerato"],
			Universidad:   wordToTfIdf["universidad"],
			Preparatoria:  wordToTfIdf["preparatoria"],
			Colegio:       wordToTfIdf["colegio"],
			Instituto:     wordToTfIdf["instituto"],
			Media:         wordToTfIdf["media"],
			Academia:      wordToTfIdf["academia"],
			Institucion:   wordToTfIdf["institucion"],
			Acosador:      wordToTfIdf["acosador"],
			Victima:       wordToTfIdf["victima"],
			Perpetrador:   wordToTfIdf["perpetrador"],
			Companeros:    wordToTfIdf["companeros"],
			Agresor:       wordToTfIdf["agresor"],
			Testigos:      wordToTfIdf["testigos"],
			Espia:         wordToTfIdf["espia"],
			Maton:         wordToTfIdf["maton"],
			Grupo:         wordToTfIdf["grupo"],
			Bully:         wordToTfIdf["bully"],
			Supervisor:    wordToTfIdf["supervisor"],
			Adolescente:   wordToTfIdf["adolescente"],
			Joven:         wordToTfIdf["joven"],
			Niño:          wordToTfIdf["niño"],
			Infantil:      wordToTfIdf["infantil"],
			Constante:     wordToTfIdf["constante"],
			Frecuente:     wordToTfIdf["frecuente"],
			Persistente:   wordToTfIdf["persistente"],
			Reiterado:     wordToTfIdf["reiterado"],
			Ocasional:     wordToTfIdf["ocasional"],
			Repetitivo:    wordToTfIdf["repetitivo"],
			Periodico:     wordToTfIdf["periodico"],
			DeletedAt:     false,
		}

		createNewSemanticOntologyCountResult := domain.SemanticOntologyCountResult{
			ProjectID:     scrapedResult.ProjectId,
			Title:         scrapedResult.Title,
			URL:           scrapedResult.Url,
			Content:       scrapedResult.Content,
			Number:        scrapedResult.Number,
			Mensajes:      wordTf["mensajes"],
			RedesSociales: wordTf["redes sociales"],
			Chat:          wordTf["chat"],
			Video:         wordTf["video"],
			Correo:        wordTf["correo"],
			Foros:         wordTf["foros"],
			Facebook:      wordTf["facebook"],
			Instagram:     wordTf["instagram"],
			SnapChat:      wordTf["snapchat"],
			WhatsApp:      wordTf["whatsapp"],
			Twitter:       wordTf["twitter"],
			YouTube:       wordTf["youtube"],
			TikTok:        wordTf["tiktok"],
			LinkedIn:      wordTf["linkedin"],
			Blog:          wordTf["blog"],
			Email:         wordTf["email"],
			Primaria:      wordTf["primaria"],
			Secundaria:    wordTf["secundaria"],
			Facultad:      wordTf["facultad"],
			Bachillerato:  wordTf["bachillerato"],
			Universidad:   wordTf["universidad"],
			Preparatoria:  wordTf["preparatoria"],
			Colegio:       wordTf["colegio"],
			Instituto:     wordTf["instituto"],
			Media:         wordTf["media"],
			Academia:      wordTf["academia"],
			Institucion:   wordTf["institucion"],
			Acosador:      wordTf["acosador"],
			Victima:       wordTf["victima"],
			Perpetrador:   wordTf["perpetrador"],
			Companeros:    wordTf["companeros"],
			Agresor:       wordTf["agresor"],
			Testigos:      wordTf["testigos"],
			Espia:         wordTf["espia"],
			Maton:         wordTf["maton"],
			Grupo:         wordTf["grupo"],
			Bully:         wordTf["bully"],
			Supervisor:    wordTf["supervisor"],
			Adolescente:   wordTf["adolescente"],
			Joven:         wordTf["joven"],
			Niño:          wordTf["niño"],
			Infantil:      wordTf["infantil"],
			Constante:     wordTf["constante"],
			Frecuente:     wordTf["frecuente"],
			Persistente:   wordTf["persistente"],
			Reiterado:     wordTf["reiterado"],
			Ocasional:     wordTf["ocasional"],
			Repetitivo:    wordTf["repetitivo"],
			Periodico:     wordTf["periodico"],
			DeletedAt:     false,
		}
		res1, errInsert := insertNewSemanticOntologyCount(client, createNewSemanticOntologyCountResult)
		if errInsert != nil {
			log.Fatal(errInsert)
		}
		fmt.Printf("Inserted document with ID %v\n", res1.InsertedID)

		res2, errInsert := insertNewSemanticOntologyTfIdf(client, createNewSemanticOntologyTfIdfResult)
		if errInsert != nil {
			log.Fatal(errInsert)
		}
		fmt.Printf("Inserted document with ID %v\n", res2.InsertedID)
	}

}

func insertNewSemanticOntologyCount(client *mongo.Client, result domain.SemanticOntologyCountResult) (*mongo.InsertOneResult, error) {
	collection := client.Database("acosoDBMongo").Collection("semantic_ontology_count_result")
	ctx := context.TODO()
	res, err := collection.InsertOne(ctx, result)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func insertNewSemanticOntologyTfIdf(client *mongo.Client, result domain.SemanticOntologyTfIdfResult) (*mongo.InsertOneResult, error) {
	collection := client.Database("acosoDBMongo").Collection("semantic_ontology_tfidf_result")
	ctx := context.TODO()
	res, err := collection.InsertOne(ctx, result)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func getCleanCorpus(htmlContent string) (string, error) {
	requestBody := RequestBody{
		Content: htmlContent,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error al convertir el cuerpo a JSON: %v", err)
	}

	url := "http://localhost:5000/clean-corpus"
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

func CountWord(corpus []string, word string) float64 {
	count := 0
	for _, doc := range corpus {
		count += strings.Count(strings.ToLower(doc), strings.ToLower(word))
	}
	return float64(count)
}

// Función para calcular el TF-IDF de una palabra en el corpus
func calculateTfidfForWordInCorpus(word string, corpus []string) float64 {
	tfIdf := 0.0
	docCount := len(corpus)
	wordDocCount := 0

	// Calcular la frecuencia de documentos que contienen la palabra
	for _, doc := range corpus {
		if strings.Contains(strings.ToLower(doc), strings.ToLower(word)) {
			wordDocCount++
		}
	}

	// Calcular IDF
	idf := math.Log(float64(docCount) / float64(wordDocCount+1)) // +1 para evitar log(0)

	// Calcular TF y TF-IDF para cada documento y tomar el máximo
	for _, doc := range corpus {
		words := strings.Fields(doc)
		totalWords := len(words)
		wordCount := 0

		for _, w := range words {
			if strings.ToLower(w) == strings.ToLower(word) {
				wordCount++
			}
		}

		tf := float64(wordCount) / float64(totalWords)
		tfIdfDoc := tf * idf

		if tfIdfDoc > tfIdf {
			tfIdf = tfIdfDoc
		}
	}

	return tfIdf
}
