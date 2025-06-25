package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

	//topWords := top200Words(scrapedResults)
	//for _, word := range topWords {
	//	fmt.Println(word)
	//}

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

	for index, scrapedResult := range scrapedResults {
		result, errGetCleanCorpus := getCleanCorpus(scrapedResult.Content)
		if errGetCleanCorpus != nil {
			log.Fatal(errGetCleanCorpus)
		}

		//Split
		corpus := strings.Fields(result)

		wordTf := map[string]float64{
			"agresivo":        0,
			"aislamiento":     0,
			"amenaza":         0,
			"ansiedad":        0,
			"ataque":          0,
			"autoestima":      0,
			"ciberbullying":   0,
			"daño":            0,
			"depresión":       0,
			"estrés":          0,
			"hostigar":        0,
			"humillar":        0,
			"insultos":        0,
			"intimidar":       0,
			"manipular":       0,
			"paranoia":        0,
			"ridiculizar":     0,
			"rumor":           0,
			"sufre":           0,
			"suicidio":        0,
			"tristeza":        0,
			"verguenza":       0,
			"violencia":       0,
			"abuso":           0,
			"cambios":         0,
			"ciberacoso":      0,
			"confidencial":    0,
			"cyberbullying":   0,
			"denigrante":      0,
			"divulgar":        0,
			"emoción":         0,
			"espiar":          0,
			"falso":           0,
			"humor":           0,
			"intencional":     0,
			"ira":             0,
			"lastimar":        0,
			"maltrato":        0,
			"poder":           0,
			"reputación":      0,
			"sexual":          0,
			"bullying":        0,
			"venganza":        0,
			"drogas":          0,
			"sustancias":      0,
			"resentimiento":   0,
			"blog":            0,
			"chat":            0,
			"correo":          0,
			"digital":         0,
			"electrónico":     0,
			"facebook":        0,
			"fotografía":      0,
			"grabación":       0,
			"internet":        0,
			"mensaje":         0,
			"movil":           0,
			"pagina":          0,
			"tecnología":      0,
			"teléfono":        0,
			"texto":           0,
			"video":           0,
			"web":             0,
			"youtube":         0,
			"cibernético":     0,
			"foto":            0,
			"imagen":          0,
			"red":             0,
			"twitter":         0,
			"virtual":         0,
			"linea":           0,
			"whatsapp":        0,
			"instagram":       0,
			"tiktok":          0,
			"linkedin":        0,
			"escuela":         0,
			"email":           0,
			"snapchat":        0,
			"foros":           0,
			"mensajes":        0,
			"preparatoria":    0,
			"primaria":        0,
			"secundaria":      0,
			"academia":        0,
			"alumno":          0,
			"alumna":          0,
			"bachillerato":    0,
			"colegio":         0,
			"educación":       0,
			"educativo":       0,
			"escolar":         0,
			"estudiante":      0,
			"facultad":        0,
			"institución":     0,
			"maestro":         0,
			"profesor":        0,
			"universidad":     0,
			"social":          0,
			"trabajo":         0,
			"país":            0,
			"físico ":         0,
			"transporte":      0,
			"centro":          0,
			"instituto":       0,
			"media":           0,
			"acosador":        0,
			"agresor":         0,
			"testigos":        0,
			"victima":         0,
			"atormentador":    0,
			"bully":           0,
			"complice":        0,
			"grupo":           0,
			"matón":           0,
			"matoneo":         0,
			"perpetrador":     0,
			"persona":         0,
			"padre":           0,
			"universitario":   0,
			"trabajador":      0,
			"mujer":           0,
			"madre":           0,
			"hombre":          0,
			"companero":       0,
			"companera":       0,
			"adulto":          0,
			"espia":           0,
			"supervisor":      0,
			"adolescente":     0,
			"joven":           0,
			"niño":            0,
			"chavo":           0,
			"chico":           0,
			"hijo":            0,
			"infantil":        0,
			"menor":           0,
			"muchacho":        0,
			"niña":            0,
			"reiterado":       0,
			"repetitivo":      0,
			"frecuente":       0,
			"año":             0,
			"constante":       0,
			"continuo":        0,
			"creciente":       0,
			"dias":            0,
			"mantiene":        0,
			"menudo":          0,
			"mes":             0,
			"períodico":       0,
			"persecución":     0,
			"perseguir":       0,
			"persistente":     0,
			"recurrente":      0,
			"repeticion":      0,
			"repetido":        0,
			"seguimiento":     0,
			"semanas":         0,
			"tiempo":          0,
			"ocasional":       0,
			"psicoterapia":    0,
			"colaboración":    0,
			"conciencia":      0,
			"equilibrio":      0,
			"identificación":  0,
			"mediación":       0,
			"orientación":     0,
			"prevención":      0,
			"sanación":        0,
			"sensibilización": 0,
			"terapia":         0,
			"autoprotección":  0,
			"establecer":      0,
			"ciberbulling":    0,
			"sextorsion":      0,
			"grooming":        0,
			"ciberviolencia":  0,
			"sexting":         0,
			"invasivo":        0,
			"racial":          0,
			"laboral":         0,
			"pareja":          0,
			"familiar":        0,
			"colectivo":       0,
			"exclusión":       0,
			"suplantación":    0,
			"denigración":     0,
			"sonsacamiento":   0,
			"doxxing":         0,
			"ciberstalking":   0,
			"día":             0,
			"noche":           0,
		}

		wordToTfIdf := map[string]float64{
			"agresivo":        0,
			"aislamiento":     0,
			"amenaza":         0,
			"ansiedad":        0,
			"ataque":          0,
			"autoestima":      0,
			"ciberbullying":   0,
			"daño":            0,
			"depresion":       0,
			"estrés":          0,
			"hostigar":        0,
			"humillar":        0,
			"insultos":        0,
			"intimidar":       0,
			"manipular":       0,
			"paranoia":        0,
			"ridiculizar":     0,
			"rumor":           0,
			"sufre":           0,
			"suicidio":        0,
			"tristeza":        0,
			"verguenza":       0,
			"violencia":       0,
			"abuso":           0,
			"cambios":         0,
			"ciberacoso":      0,
			"confidencial":    0,
			"cyberbullying":   0,
			"denigrante":      0,
			"divulgar":        0,
			"emoción":         0,
			"espiar":          0,
			"falso":           0,
			"humor":           0,
			"intencional":     0,
			"ira":             0,
			"lastimar":        0,
			"maltrato":        0,
			"poder":           0,
			"reputación":      0,
			"sexual":          0,
			"bullying":        0,
			"venganza":        0,
			"drogas":          0,
			"sustancias":      0,
			"resentimiento":   0,
			"blog":            0,
			"chat":            0,
			"correo":          0,
			"digital":         0,
			"electrónico":     0,
			"facebook":        0,
			"fotografía":      0,
			"grabación":       0,
			"internet":        0,
			"mensaje":         0,
			"movil":           0,
			"pagina":          0,
			"tecnología":      0,
			"teléfono":        0,
			"texto":           0,
			"video":           0,
			"web":             0,
			"youtube":         0,
			"cibernético":     0,
			"foto":            0,
			"imagen":          0,
			"red":             0,
			"twitter":         0,
			"virtual":         0,
			"linea":           0,
			"whatsapp":        0,
			"instagram":       0,
			"tiktok":          0,
			"linkedin":        0,
			"escuela":         0,
			"email":           0,
			"snapchat":        0,
			"foros":           0,
			"mensajes":        0,
			"preparatoria":    0,
			"primaria":        0,
			"secundaria":      0,
			"academia":        0,
			"alumno":          0,
			"alumna":          0,
			"bachillerato":    0,
			"colegio":         0,
			"educación":       0,
			"educativo":       0,
			"escolar":         0,
			"estudiante":      0,
			"facultad":        0,
			"institución":     0,
			"maestro":         0,
			"profesor":        0,
			"universidad":     0,
			"social":          0,
			"trabajo":         0,
			"país":            0,
			"físico":          0,
			"transporte":      0,
			"centro":          0,
			"instituto":       0,
			"media":           0,
			"acosador":        0,
			"agresor":         0,
			"testigos":        0,
			"victima":         0,
			"atormentador":    0,
			"bully":           0,
			"complice":        0,
			"grupo":           0,
			"matón":           0,
			"matoneo":         0,
			"perpetrador":     0,
			"persona":         0,
			"padre":           0,
			"universitario":   0,
			"trabajador":      0,
			"mujer":           0,
			"madre":           0,
			"hombre":          0,
			"companero":       0,
			"companera":       0,
			"adulto":          0,
			"espia":           0,
			"supervisor":      0,
			"adolescente":     0,
			"joven":           0,
			"niño":            0,
			"chavo":           0,
			"chico":           0,
			"hijo":            0,
			"infantil":        0,
			"menor":           0,
			"muchacho":        0,
			"niña":            0,
			"reiterado":       0,
			"repetitivo":      0,
			"frecuente":       0,
			"año":             0,
			"constante":       0,
			"continuo":        0,
			"creciente":       0,
			"dias":            0,
			"mantiene":        0,
			"menudo":          0,
			"mes":             0,
			"períodico":       0,
			"persecución":     0,
			"perseguir":       0,
			"persistente":     0,
			"recurrente":      0,
			"repeticion":      0,
			"repetido":        0,
			"seguimiento":     0,
			"semanas":         0,
			"tiempo":          0,
			"ocasional":       0,
			"psicoterapia":    0,
			"colaboración":    0,
			"conciencia":      0,
			"equilibrio":      0,
			"identificación":  0,
			"mediación":       0,
			"orientación":     0,
			"prevención":      0,
			"sanación":        0,
			"sensibilización": 0,
			"terapia":         0,
			"autoprotección":  0,
			"establecer":      0,
			"ciberbulling":    0,
			"sextorsion":      0,
			"grooming":        0,
			"ciberviolencia":  0,
			"sexting":         0,
			"invasivo":        0,
			"racial":          0,
			"laboral":         0,
			"pareja":          0,
			"familiar":        0,
			"colectivo":       0,
			"exclusión":       0,
			"suplantación":    0,
			"denigración":     0,
			"sonsacamiento":   0,
			"doxxing":         0,
			"ciberstalking":   0,
			"día":             0,
			"noche":           0,
		}

		for word, _ := range wordTf {
			wordTf[word] = CountWord(corpus, word)
			wordToTfIdf[word] = calculateTfidfForWordInCorpus(word, corpus)
		}

		isCorrectValue := 1
		if isCorrect(wordTf) == true {
			isCorrectValue = 1
		} else {
			isCorrectValue = 0
		}

		createNewSemanticOntologyTfIdfResult := domain.SemanticOntologyTfIdfResult{
			ProjectID:       scrapedResult.ProjectId,
			Title:           scrapedResult.Title,
			URL:             scrapedResult.Url,
			Content:         scrapedResult.Content,
			Number:          scrapedResult.Number,
			Agresivo:        wordToTfIdf["agresivo"],
			Aislamiento:     wordToTfIdf["aislamiento"],
			Amenaza:         wordToTfIdf["amenaza"],
			Ansiedad:        wordToTfIdf["ansiedad"],
			Ataque:          wordToTfIdf["ataque"],
			Autoestima:      wordToTfIdf["autoestima"],
			Ciberbullying:   wordToTfIdf["ciberbullying"],
			Daño:            wordToTfIdf["daño"],
			Depresion:       wordToTfIdf["depresión"],
			Estres:          wordToTfIdf["estrés"],
			Hostigar:        wordToTfIdf["hostigar"],
			Humillar:        wordToTfIdf["humillar"],
			Insultos:        wordToTfIdf["insultos"],
			Intimidar:       wordToTfIdf["intimidar"],
			Manipular:       wordToTfIdf["manipular"],
			Paranoia:        wordToTfIdf["paranoia"],
			Ridiculizar:     wordToTfIdf["ridiculizar"],
			Rumor:           wordToTfIdf["rumor"],
			Sufre:           wordToTfIdf["sufre"],
			Suicidio:        wordToTfIdf["suicidio"],
			Tristeza:        wordToTfIdf["tristeza"],
			Verguenza:       wordToTfIdf["verguenza"],
			Violencia:       wordToTfIdf["violencia"],
			Abuso:           wordToTfIdf["abuso"],
			Cambios:         wordToTfIdf["cambios"],
			Ciberacoso:      wordToTfIdf["ciberacoso"],
			Confidencial:    wordToTfIdf["confidencial"],
			Cyberbullying:   wordToTfIdf["cyberbullying"],
			Denigrante:      wordToTfIdf["denigrante"],
			Divulgar:        wordToTfIdf["divulgar"],
			Emocion:         wordToTfIdf["emoción"],
			Espiar:          wordToTfIdf["espiar"],
			Falso:           wordToTfIdf["falso"],
			Humor:           wordToTfIdf["humor"],
			Intencional:     wordToTfIdf["intencional"],
			Ira:             wordToTfIdf["ira"],
			Lastimar:        wordToTfIdf["lastimar"],
			Maltrato:        wordToTfIdf["maltrato"],
			Poder:           wordToTfIdf["poder"],
			Reputacion:      wordToTfIdf["reputación"],
			Sexual:          wordToTfIdf["sexual"],
			Bullying:        wordToTfIdf["bullying"],
			Venganza:        wordToTfIdf["venganza"],
			Drogas:          wordToTfIdf["drogas"],
			Sustancias:      wordToTfIdf["sustancias"],
			Resentimiento:   wordToTfIdf["resentimiento"],
			Blog:            wordToTfIdf["blog"],
			Chat:            wordToTfIdf["chat"],
			Correo:          wordToTfIdf["correo"],
			Digital:         wordToTfIdf["digital"],
			Electronico:     wordToTfIdf["electrónico"],
			Facebook:        wordToTfIdf["facebook"],
			Fotografia:      wordToTfIdf["fotografía"],
			Grabacion:       wordToTfIdf["grabación"],
			Internet:        wordToTfIdf["internet"],
			Mensaje:         wordToTfIdf["mensaje"],
			Movil:           wordToTfIdf["movil"],
			Pagina:          wordToTfIdf["pagina"],
			Tecnologia:      wordToTfIdf["tecnología"],
			Telefono:        wordToTfIdf["teléfono"],
			Texto:           wordToTfIdf["texto"],
			Video:           wordToTfIdf["video"],
			Web:             wordToTfIdf["web"],
			Youtube:         wordToTfIdf["youtube"],
			Cibernetico:     wordToTfIdf["cibernético"],
			Foto:            wordToTfIdf["foto"],
			Imagen:          wordToTfIdf["imagen"],
			Red:             wordToTfIdf["red"],
			Twitter:         wordToTfIdf["twitter"],
			Virtual:         wordToTfIdf["virtual"],
			Linea:           wordToTfIdf["linea"],
			Whatsapp:        wordToTfIdf["whatsapp"],
			Instagram:       wordToTfIdf["instagram"],
			Tiktok:          wordToTfIdf["tiktok"],
			Linkedin:        wordToTfIdf["linkedin"],
			Escuela:         wordToTfIdf["escuela"],
			Email:           wordToTfIdf["email"],
			Snapchat:        wordToTfIdf["snapchat"],
			Foros:           wordToTfIdf["foros"],
			Mensajes:        wordToTfIdf["mensajes"],
			Preparatoria:    wordToTfIdf["preparatoria"],
			Primaria:        wordToTfIdf["primaria"],
			Secundaria:      wordToTfIdf["secundaria"],
			Academia:        wordToTfIdf["academia"],
			Alumno:          wordToTfIdf["alumno"],
			Alumna:          wordToTfIdf["alumna"],
			Bachillerato:    wordToTfIdf["bachillerato"],
			Colegio:         wordToTfIdf["colegio"],
			Educacion:       wordToTfIdf["educación"],
			Educativo:       wordToTfIdf["educativo"],
			Escolar:         wordToTfIdf["escolar"],
			Estudiante:      wordToTfIdf["estudiante"],
			Facultad:        wordToTfIdf["facultad"],
			Institucion:     wordToTfIdf["institucion"],
			Maestro:         wordToTfIdf["maestro"],
			Profesor:        wordToTfIdf["profesor"],
			Universidad:     wordToTfIdf["universidad"],
			Social:          wordToTfIdf["social"],
			Trabajo:         wordToTfIdf["trabajo"],
			Pais:            wordToTfIdf["país"],
			Fisico:          wordToTfIdf["físico"],
			Transporte:      wordToTfIdf["transporte"],
			Centro:          wordToTfIdf["centro"],
			Instituto:       wordToTfIdf["instituto"],
			Media:           wordToTfIdf["media"],
			Acosador:        wordToTfIdf["acosador"],
			Agresor:         wordToTfIdf["agresor"],
			Testigos:        wordToTfIdf["testigos"],
			Victima:         wordToTfIdf["victima"],
			Atormentador:    wordToTfIdf["atormentador"],
			Bully:           wordToTfIdf["bully"],
			Complice:        wordToTfIdf["complice"],
			Grupo:           wordToTfIdf["grupo"],
			Maton:           wordToTfIdf["matón"],
			Matoneo:         wordToTfIdf["matoneo"],
			Perpetrador:     wordToTfIdf["perpetrador"],
			Persona:         wordToTfIdf["persona"],
			Padre:           wordToTfIdf["padre"],
			Universitario:   wordToTfIdf["universitario"],
			Trabajador:      wordToTfIdf["trabajador"],
			Mujer:           wordToTfIdf["mujer"],
			Madre:           wordToTfIdf["madre"],
			Hombre:          wordToTfIdf["hombre"],
			Companero:       wordToTfIdf["compañero"],
			Companera:       wordToTfIdf["compañera"],
			Adulto:          wordToTfIdf["adulto"],
			Espia:           wordToTfIdf["espia"],
			Supervisor:      wordToTfIdf["supervisor"],
			Adolescente:     wordToTfIdf["adolescente"],
			Joven:           wordToTfIdf["joven"],
			Niño:            wordToTfIdf["niño"],
			Chavo:           wordToTfIdf["chavo"],
			Chico:           wordToTfIdf["chico"],
			Hijo:            wordToTfIdf["hijo"],
			Infantil:        wordToTfIdf["infantil"],
			Menor:           wordToTfIdf["menor"],
			Muchacho:        wordToTfIdf["muchacho"],
			Nina:            wordToTfIdf["niña"],
			Reiterado:       wordToTfIdf["reiterado"],
			Repetitivo:      wordToTfIdf["repetitivo"],
			Frecuente:       wordToTfIdf["frecuente"],
			Ano:             wordToTfIdf["ano"],
			Constante:       wordToTfIdf["constante"],
			Continuo:        wordToTfIdf["continuo"],
			Creciente:       wordToTfIdf["creciente"],
			Dias:            wordToTfIdf["dias"],
			Mantiene:        wordToTfIdf["mantiene"],
			Menudo:          wordToTfIdf["menudo"],
			Mes:             wordToTfIdf["mes"],
			Periodico:       wordToTfIdf["períodico"],
			Persecucion:     wordToTfIdf["persecucion"],
			Perseguir:       wordToTfIdf["perseguir"],
			Persistente:     wordToTfIdf["persistente"],
			Recurrente:      wordToTfIdf["recurrente"],
			Repeticion:      wordToTfIdf["repeticion"],
			Repetido:        wordToTfIdf["repetido"],
			Seguimiento:     wordToTfIdf["seguimiento"],
			Semanas:         wordToTfIdf["semanas"],
			Tiempo:          wordToTfIdf["tiempo"],
			Ocasional:       wordToTfIdf["ocasional"],
			Psicoterapia:    wordToTfIdf["psicoterapia"],
			Colaboracion:    wordToTfIdf["colaboración"],
			Conciencia:      wordToTfIdf["conciencia"],
			Equilibrio:      wordToTfIdf["equilibrio"],
			Identificacion:  wordToTfIdf["identificación"],
			Mediacion:       wordToTfIdf["mediación"],
			Orientacion:     wordToTfIdf["orientación"],
			Prevencion:      wordToTfIdf["prevención"],
			Sanacion:        wordToTfIdf["sanación"],
			Sensibilizacion: wordToTfIdf["sensibilización"],
			Terapia:         wordToTfIdf["terapia"],
			Autoproteccion:  wordToTfIdf["autoprotección"],
			Establecer:      wordToTfIdf["establecer"],
			Ciberbulling:    wordToTfIdf["ciberbulling"],
			Sextorsion:      wordToTfIdf["sextorsion"],
			Grooming:        wordToTfIdf["grooming"],
			Ciberviolencia:  wordToTfIdf["ciberviolencia"],
			Sexting:         wordToTfIdf["sexting"],
			Invasivo:        wordToTfIdf["invasivo"],
			Racial:          wordToTfIdf["racial"],
			Laboral:         wordToTfIdf["laboral"],
			Pareja:          wordToTfIdf["pareja"],
			Familiar:        wordToTfIdf["familiar"],
			Colectivo:       wordToTfIdf["colectivo"],
			Exclusion:       wordToTfIdf["exclusión"],
			Suplantacion:    wordToTfIdf["suplantación"],
			Denigracion:     wordToTfIdf["denigración"],
			Sonsacamiento:   wordToTfIdf["sonsacamiento"],
			Doxxing:         wordToTfIdf["doxxing"],
			Ciberstalking:   wordToTfIdf["ciberstalking"],
			Dia:             wordToTfIdf["día"],
			Noche:           wordToTfIdf["noche"],
			Correcto:        isCorrectValue,
			DeletedAt:       false,
		}

		createNewSemanticOntologyCountResult := domain.SemanticOntologyCountResult{
			ProjectID:       scrapedResult.ProjectId,
			Title:           scrapedResult.Title,
			URL:             scrapedResult.Url,
			Content:         scrapedResult.Content,
			Number:          scrapedResult.Number,
			Agresivo:        wordTf["agresivo"],
			Aislamiento:     wordTf["aislamiento"],
			Amenaza:         wordTf["amenaza"],
			Ansiedad:        wordTf["ansiedad"],
			Ataque:          wordTf["ataque"],
			Autoestima:      wordTf["autoestima"],
			Ciberbullying:   wordTf["ciberbullying"],
			Daño:            wordTf["daño"],
			Depresion:       wordTf["depresión"],
			Estres:          wordTf["estrés"],
			Hostigar:        wordTf["hostigar"],
			Humillar:        wordTf["humillar"],
			Insultos:        wordTf["insultos"],
			Intimidar:       wordTf["intimidar"],
			Manipular:       wordTf["manipular"],
			Paranoia:        wordTf["paranoia"],
			Ridiculizar:     wordTf["ridiculizar"],
			Rumor:           wordTf["rumor"],
			Sufre:           wordTf["sufre"],
			Suicidio:        wordTf["suicidio"],
			Tristeza:        wordTf["tristeza"],
			Verguenza:       wordTf["verguenza"],
			Violencia:       wordTf["violencia"],
			Abuso:           wordTf["abuso"],
			Cambios:         wordTf["cambios"],
			Ciberacoso:      wordTf["ciberacoso"],
			Confidencial:    wordTf["confidencial"],
			Cyberbullying:   wordTf["cyberbullying"],
			Denigrante:      wordTf["denigrante"],
			Divulgar:        wordTf["divulgar"],
			Emocion:         wordTf["emoción"],
			Espiar:          wordTf["espiar"],
			Falso:           wordTf["falso"],
			Humor:           wordTf["humor"],
			Intencional:     wordTf["intencional"],
			Ira:             wordTf["ira"],
			Lastimar:        wordTf["lastimar"],
			Maltrato:        wordTf["maltrato"],
			Poder:           wordTf["poder"],
			Reputacion:      wordTf["reputación"],
			Sexual:          wordTf["sexual"],
			Bullying:        wordTf["bullying"],
			Venganza:        wordTf["venganza"],
			Drogas:          wordTf["drogas"],
			Sustancias:      wordTf["sustancias"],
			Resentimiento:   wordTf["resentimiento"],
			Blog:            wordTf["blog"],
			Chat:            wordTf["chat"],
			Correo:          wordTf["correo"],
			Digital:         wordTf["digital"],
			Electronico:     wordTf["electrónico"],
			Facebook:        wordTf["facebook"],
			Fotografia:      wordTf["fotografía"],
			Grabacion:       wordTf["grabación"],
			Internet:        wordTf["internet"],
			Mensaje:         wordTf["mensaje"],
			Movil:           wordTf["movil"],
			Pagina:          wordTf["pagina"],
			Tecnologia:      wordTf["tecnología"],
			Telefono:        wordTf["teléfono"],
			Texto:           wordTf["texto"],
			Video:           wordTf["video"],
			Web:             wordTf["web"],
			Youtube:         wordTf["youtube"],
			Cibernetico:     wordTf["cibernético"],
			Foto:            wordTf["foto"],
			Imagen:          wordTf["imagen"],
			Red:             wordTf["red"],
			Twitter:         wordTf["twitter"],
			Virtual:         wordTf["virtual"],
			Linea:           wordTf["linea"],
			Whatsapp:        wordTf["whatsapp"],
			Instagram:       wordTf["instagram"],
			Tiktok:          wordTf["tiktok"],
			Linkedin:        wordTf["linkedin"],
			Escuela:         wordTf["escuela"],
			Email:           wordTf["email"],
			Snapchat:        wordTf["snapchat"],
			Foros:           wordTf["foros"],
			Mensajes:        wordTf["mensajes"],
			Preparatoria:    wordTf["preparatoria"],
			Primaria:        wordTf["primaria"],
			Secundaria:      wordTf["secundaria"],
			Academia:        wordTf["academia"],
			Alumno:          wordTf["alumno"],
			Alumna:          wordTf["alumna"],
			Bachillerato:    wordTf["bachillerato"],
			Colegio:         wordTf["colegio"],
			Educacion:       wordTf["educación"],
			Educativo:       wordTf["educativo"],
			Escolar:         wordTf["escolar"],
			Estudiante:      wordTf["estudiante"],
			Facultad:        wordTf["facultad"],
			Institucion:     wordTf["institucion"],
			Maestro:         wordTf["maestro"],
			Profesor:        wordTf["profesor"],
			Universidad:     wordTf["universidad"],
			Social:          wordTf["social"],
			Trabajo:         wordTf["trabajo"],
			Pais:            wordTf["país"],
			Fisico:          wordTf["físico"],
			Transporte:      wordTf["transporte"],
			Centro:          wordTf["centro"],
			Instituto:       wordTf["instituto"],
			Media:           wordTf["media"],
			Acosador:        wordTf["acosador"],
			Agresor:         wordTf["agresor"],
			Testigos:        wordTf["testigos"],
			Victima:         wordTf["victima"],
			Atormentador:    wordTf["atormentador"],
			Bully:           wordTf["bully"],
			Complice:        wordTf["complice"],
			Grupo:           wordTf["grupo"],
			Maton:           wordTf["matón"],
			Matoneo:         wordTf["matoneo"],
			Perpetrador:     wordTf["perpetrador"],
			Persona:         wordTf["persona"],
			Padre:           wordTf["padre"],
			Universitario:   wordTf["universitario"],
			Trabajador:      wordTf["trabajador"],
			Mujer:           wordTf["mujer"],
			Madre:           wordTf["madre"],
			Hombre:          wordTf["hombre"],
			Companero:       wordTf["compañero"],
			Companera:       wordTf["compañera"],
			Adulto:          wordTf["adulto"],
			Espia:           wordTf["espia"],
			Supervisor:      wordTf["supervisor"],
			Adolescente:     wordTf["adolescente"],
			Joven:           wordTf["joven"],
			Niño:            wordTf["niño"],
			Chavo:           wordTf["chavo"],
			Chico:           wordTf["chico"],
			Hijo:            wordTf["hijo"],
			Infantil:        wordTf["infantil"],
			Menor:           wordTf["menor"],
			Muchacho:        wordTf["muchacho"],
			Nina:            wordTf["niña"],
			Reiterado:       wordTf["reiterado"],
			Repetitivo:      wordTf["repetitivo"],
			Frecuente:       wordTf["frecuente"],
			Ano:             wordTf["ano"],
			Constante:       wordTf["constante"],
			Continuo:        wordTf["continuo"],
			Creciente:       wordTf["creciente"],
			Dias:            wordTf["dias"],
			Mantiene:        wordTf["mantiene"],
			Menudo:          wordTf["menudo"],
			Mes:             wordTf["mes"],
			Periodico:       wordTf["períodico"],
			Persecucion:     wordTf["persecucion"],
			Perseguir:       wordTf["perseguir"],
			Persistente:     wordTf["persistente"],
			Recurrente:      wordTf["recurrente"],
			Repeticion:      wordTf["repeticion"],
			Repetido:        wordTf["repetido"],
			Seguimiento:     wordTf["seguimiento"],
			Semanas:         wordTf["semanas"],
			Tiempo:          wordTf["tiempo"],
			Ocasional:       wordTf["ocasional"],
			Psicoterapia:    wordTf["psicoterapia"],
			Colaboracion:    wordTf["colaboración"],
			Conciencia:      wordTf["conciencia"],
			Identificacion:  wordTf["identificación"],
			Mediacion:       wordTf["mediación"],
			Orientacion:     wordTf["orientación"],
			Prevencion:      wordTf["prevención"],
			Sanacion:        wordTf["sanación"],
			Sensibilizacion: wordTf["sensibilización"],
			Terapia:         wordTf["terapia"],
			Autoproteccion:  wordTf["autoproteccion"],
			Establecer:      wordTf["establecer"],
			Ciberbulling:    wordTf["ciberbulling"],
			Sextorsion:      wordTf["sextorsion"],
			Grooming:        wordTf["grooming"],
			Ciberviolencia:  wordTf["ciberviolencia"],
			Sexting:         wordTf["sexting"],
			Invasivo:        wordTf["invasivo"],
			Racial:          wordTf["racial"],
			Laboral:         wordTf["laboral"],
			Pareja:          wordTf["pareja"],
			Familiar:        wordTf["familiar"],
			Colectivo:       wordTf["colectivo"],
			Exclusion:       wordTf["exclusión"],
			Suplantacion:    wordTf["suplantación"],
			Denigracion:     wordTf["denigración"],
			Sonsacamiento:   wordTf["sonsacamiento"],
			Doxxing:         wordTf["doxxing"],
			Ciberstalking:   wordTf["ciberstalking"],
			Dia:             wordTf["día"],
			Noche:           wordTf["noche"],
			Correcto:        isCorrectValue,
			DeletedAt:       false,
		}
		_, errInsertTf := insertNewSemanticOntologyCount(client, createNewSemanticOntologyCountResult)
		if errInsertTf != nil {
			log.Fatal(errInsertTf)
		}

		_, errInsertTfIdf := insertNewSemanticOntologyTfIdf(client, createNewSemanticOntologyTfIdfResult)
		if errInsertTfIdf != nil {
			log.Fatal(errInsertTfIdf)
		}
		fmt.Printf("Registro numero %d \n", index)
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

// Función para contar las palabras en el corpus
func countWords(corpus string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(corpus)
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}

// Función para obtener las 200 palabras más frecuentes
func top200Words(scrapedResults []domain.ScrapedResult) []string {
	combinedCorpus := ""
	for _, result := range scrapedResults {
		cleanCorpus, _ := getCleanCorpus(result.Content)
		combinedCorpus += " " + cleanCorpus
	}

	wordCount := countWords(combinedCorpus)

	type wordFreq struct {
		word string
		freq int
	}

	// Convertir el mapa de conteo de palabras en una lista de wordFreq
	var freqList []wordFreq
	for word, freq := range wordCount {
		freqList = append(freqList, wordFreq{word, freq})
	}

	// Ordenar por frecuencia (de mayor a menor)
	sort.Slice(freqList, func(i, j int) bool {
		return freqList[i].freq > freqList[j].freq
	})

	// Obtener las 200 palabras más frecuentes
	var topWords []string
	for i := 0; i < len(freqList) && i < 200; i++ {
		topWords = append(topWords, freqList[i].word)
	}

	return topWords
}

func isCategoryPresent(wordTf map[string]float64, categoryWords []string) bool {
	for _, word := range categoryWords {
		if value, exists := wordTf[word]; exists && value != 0 {
			return true
		}
	}
	return false
}

func isCorrect(wordTf map[string]float64) bool {
	// Definir las palabras clave para cada categoría
	factoresRiesgo := []string{
		"agresivo", "aislamiento", "amenaza", "ansiedad", "ataque", "autoestima", "ciberbullying",
		"daño", "depresión", "estrés", "hostigar", "humillar", "insultos", "intimidar", "manipular",
		"paranoia", "ridiculizar", "rumor", "sufre", "suicidio", "tristeza", "verguenza", "violencia",
		"abuso", "cambios", "ciberacoso", "confidencial", "cyberbullying", "denigrante", "divulgar",
		"emoción", "espiar", "falso", "humor", "intencional", "ira", "lastimar", "maltrato", "poder",
		"reputación", "sexual", "bullying", "venganza", "drogas", "sustancias", "resentimiento",
	}
	tics := []string{
		"blog", "chat", "correo", "digital", "electrónico", "facebook", "fotografía", "grabación",
		"internet", "mensaje", "movil", "pagina", "tecnología", "teléfono", "texto", "video", "web",
		"youtube", "cibernético", "foto", "imagen", "red", "twitter", "virtual", "linea", "whatsapp",
		"instagram", "tiktok", "linkedin", "email", "snapchat", "foros", "mensajes",
	}
	partipantes := []string{
		"acosador", "agresor", "testigos", "victima", "atormentador", "bully", "complice", "grupo",
		"matón", "matoneo", "perpetrador", "persona", "padre", "universitario", "trabajador", "mujer",
		"madre", "hombre", "compañero", "compañera", "adulto", "espia", "supervisor", "adolescente",
		"joven", "niño", "chavo", "chico", "hijo", "infantil", "menor", "muchacho", "niña",
	}
	tiposCiberacoso := []string{
		"ciberbulling", "sextorsion", "grooming", "ciberviolencia", "sexting", "invasivo", "racial",
		"laboral", "pareja", "familiar", "colectivo", "exclusión", "suplantación", "denigración",
		"sonsacamiento", "doxxing", "ciberstalking",
	}
	prevencion := []string{
		"psicoterapia", "colaboración", "conciencia", "identificación", "mediación", "orientación", "prevención", "sanación", "sensibilización", "terapia", "autoprotección", "establecer",
	}

	return isCategoryPresent(wordTf, factoresRiesgo) ||
		isCategoryPresent(wordTf, tics) ||
		isCategoryPresent(wordTf, partipantes) ||
		isCategoryPresent(wordTf, tiposCiberacoso) ||
		isCategoryPresent(wordTf, prevencion)
}
