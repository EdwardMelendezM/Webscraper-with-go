package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"webscraper-go/scraped-results/domain"
)

func main() {
	// Conexión a la base de datos
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Pinging MongoDB para verificar conexión
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado a MongoDB!")

	// Selecciona la base de datos y la colección
	collection := client.Database("acosoDBMongo").Collection("semantic_ontology_count_result")

	// Abre el archivo CSV
	file, err := os.Create("semantic_ontology_count_result.csv")
	if err != nil {
		log.Fatal("No se pudo crear el archivo CSV:", err)
	}
	defer file.Close()

	// Crea un escritor CSV con separador de punto y coma
	writer := csv.NewWriter(file)
	writer.Comma = ';' // Configura el separador a punto y coma
	defer writer.Flush()

	// Escribe la cabecera del CSV
	header := []string{
		"Title",
		"URL",
		"Content",
		"agresivo",
		"aislamiento",
		"amenaza",
		"ansiedad",
		"ataque",
		"autoestima",
		"ciberbullying",
		"daño",
		"depresión",
		"estrés",
		"hostigar",
		"humillar",
		"insultos",
		"intimidar",
		"manipular",
		"paranoia",
		"ridiculizar",
		"rumor",
		"sufre",
		"suicidio",
		"tristeza",
		"verguenza",
		"violencia",
		"abuso",
		"cambios",
		"ciberacoso",
		"confidencial",
		"cyberbullying",
		"denigrante",
		"divulgar",
		"emoción",
		"espiar",
		"falso",
		"humor",
		"intencional",
		"ira",
		"lastimar",
		"maltrato",
		"poder",
		"reputación",
		"sexual",
		"bullying",
		"venganza",
		"drogas",
		"sustancias",
		"resentimiento",
		"blog",
		"chat",
		"correo",
		"digital",
		"electrónico",
		"facebook",
		"fotografía",
		"grabación",
		"internet",
		"mensaje",
		"movil",
		"pagina",
		"tecnología",
		"teléfono",
		"texto",
		"video",
		"web",
		"youtube",
		"cibernético",
		"foto",
		"imagen",
		"red",
		"twitter",
		"virtual",
		"linea",
		"whatsapp",
		"instagram",
		"tiktok",
		"linkedin",
		"escuela",
		"email",
		"snapchat",
		"foros",
		"mensajes",
		"preparatoria",
		"primaria",
		"secundaria",
		"academia",
		"alumno",
		"alumna",
		"bachillerato",
		"colegio",
		"educación",
		"educativo",
		"escolar",
		"estudiante",
		"facultad",
		"institucion",
		"maestro",
		"profesor",
		"universidad",
		"social",
		"trabajo",
		"país",
		"físico",
		"transporte",
		"centro",
		"instituto",
		"media",
		"acosador",
		"agresor",
		"testigos",
		"victima",
		"atormentador",
		"bully",
		"complice",
		"grupo",
		"matón",
		"matoneo",
		"perpetrador",
		"persona",
		"padre",
		"universitario",
		"trabajador",
		"mujer",
		"madre",
		"hombre",
		"compañero",
		"compañera",
		"adulto",
		"espia",
		"supervisor",
		"adolescente",
		"joven",
		"niño",
		"chavo",
		"chico",
		"hijo",
		"infantil",
		"menor",
		"muchacho",
		"niña",
		"reiterado",
		"repetitivo",
		"frecuente",
		"ano",
		"constante",
		"continuo",
		"creciente",
		"dias",
		"mantiene",
		"menudo",
		"mes",
		"períodico",
		"persecucion",
		"perseguir",
		"persistente",
		"recurrente",
		"repeticion",
		"repetido",
		"seguimiento",
		"semanas",
		"tiempo",
		"ocasional",
		"psicoterapia",
		"colaboración",
		"conciencia",
		"equilibrio",
		"identificación",
		"mediación",
		"orientación",
		"prevención",
		"sanación",
		"sensibilización",
		"terapia",
		"autoprotección",
		"establecer",
		"ciberbulling",
		"sextorsion",
		"grooming",
		"ciberviolencia",
		"sexting",
		"invasivo",
		"racial",
		"laboral",
		"pareja",
		"familiar",
		"colectivo",
		"exclusión",
		"suplantación",
		"denigración",
		"sonsacamiento",
		"doxxing",
		"ciberstalking",
		"día",
		"noche",
		"correcto",
	}
	err = writer.Write(header)
	if err != nil {
		log.Fatal("No se pudo escribir la cabecera CSV:", err)
	}

	// Define un filtro si deseas filtrar los datos, en este caso no aplicamos filtro
	filter := bson.D{}

	// Encuentra los datos
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())

	// Itera sobre los resultados y escribe en el CSV
	for cur.Next(context.TODO()) {
		var result domain.SemanticOntologyTfIdfResult
		errDecode := cur.Decode(&result)
		if errDecode != nil {
			log.Fatal(errDecode)
		}

		content := strings.ReplaceAll(result.Content, ";", " ")

		// Convierte el resultado a un slice de strings para escribir en CSV
		record := []string{
			result.Title,
			result.URL,
			content,
			fmt.Sprintf("%f", result.Agresivo),
			fmt.Sprintf("%f", result.Aislamiento),
			fmt.Sprintf("%f", result.Amenaza),
			fmt.Sprintf("%f", result.Ansiedad),
			fmt.Sprintf("%f", result.Ataque),
			fmt.Sprintf("%f", result.Autoestima),
			fmt.Sprintf("%f", result.Ciberbullying),
			fmt.Sprintf("%f", result.Daño),
			fmt.Sprintf("%f", result.Depresion),
			fmt.Sprintf("%f", result.Estres),
			fmt.Sprintf("%f", result.Hostigar),
			fmt.Sprintf("%f", result.Humillar),
			fmt.Sprintf("%f", result.Insultos),
			fmt.Sprintf("%f", result.Intimidar),
			fmt.Sprintf("%f", result.Manipular),
			fmt.Sprintf("%f", result.Paranoia),
			fmt.Sprintf("%f", result.Ridiculizar),
			fmt.Sprintf("%f", result.Rumor),
			fmt.Sprintf("%f", result.Sufre),
			fmt.Sprintf("%f", result.Suicidio),
			fmt.Sprintf("%f", result.Tristeza),
			fmt.Sprintf("%f", result.Verguenza),
			fmt.Sprintf("%f", result.Violencia),
			fmt.Sprintf("%f", result.Abuso),
			fmt.Sprintf("%f", result.Cambios),
			fmt.Sprintf("%f", result.Ciberacoso),
			fmt.Sprintf("%f", result.Confidencial),
			fmt.Sprintf("%f", result.Cyberbullying),
			fmt.Sprintf("%f", result.Denigrante),
			fmt.Sprintf("%f", result.Divulgar),
			fmt.Sprintf("%f", result.Emocion),
			fmt.Sprintf("%f", result.Espiar),
			fmt.Sprintf("%f", result.Falso),
			fmt.Sprintf("%f", result.Humor),
			fmt.Sprintf("%f", result.Intencional),
			fmt.Sprintf("%f", result.Ira),
			fmt.Sprintf("%f", result.Lastimar),
			fmt.Sprintf("%f", result.Maltrato),
			fmt.Sprintf("%f", result.Poder),
			fmt.Sprintf("%f", result.Reputacion),
			fmt.Sprintf("%f", result.Sexual),
			fmt.Sprintf("%f", result.Bullying),
			fmt.Sprintf("%f", result.Venganza),
			fmt.Sprintf("%f", result.Drogas),
			fmt.Sprintf("%f", result.Sustancias),
			fmt.Sprintf("%f", result.Resentimiento),
			fmt.Sprintf("%f", result.Blog),
			fmt.Sprintf("%f", result.Chat),
			fmt.Sprintf("%f", result.Correo),
			fmt.Sprintf("%f", result.Digital),
			fmt.Sprintf("%f", result.Electronico),
			fmt.Sprintf("%f", result.Facebook),
			fmt.Sprintf("%f", result.Fotografia),
			fmt.Sprintf("%f", result.Grabacion),
			fmt.Sprintf("%f", result.Internet),
			fmt.Sprintf("%f", result.Mensaje),
			fmt.Sprintf("%f", result.Movil),
			fmt.Sprintf("%f", result.Pagina),
			fmt.Sprintf("%f", result.Tecnologia),
			fmt.Sprintf("%f", result.Telefono),
			fmt.Sprintf("%f", result.Texto),
			fmt.Sprintf("%f", result.Video),
			fmt.Sprintf("%f", result.Web),
			fmt.Sprintf("%f", result.Youtube),
			fmt.Sprintf("%f", result.Cibernetico),
			fmt.Sprintf("%f", result.Foto),
			fmt.Sprintf("%f", result.Imagen),
			fmt.Sprintf("%f", result.Red),
			fmt.Sprintf("%f", result.Twitter),
			fmt.Sprintf("%f", result.Virtual),
			fmt.Sprintf("%f", result.Linea),
			fmt.Sprintf("%f", result.Whatsapp),
			fmt.Sprintf("%f", result.Instagram),
			fmt.Sprintf("%f", result.Tiktok),
			fmt.Sprintf("%f", result.Linkedin),
			fmt.Sprintf("%f", result.Escuela),
			fmt.Sprintf("%f", result.Email),
			fmt.Sprintf("%f", result.Snapchat),
			fmt.Sprintf("%f", result.Foros),
			fmt.Sprintf("%f", result.Mensajes),
			fmt.Sprintf("%f", result.Preparatoria),
			fmt.Sprintf("%f", result.Primaria),
			fmt.Sprintf("%f", result.Secundaria),
			fmt.Sprintf("%f", result.Academia),
			fmt.Sprintf("%f", result.Alumno),
			fmt.Sprintf("%f", result.Alumna),
			fmt.Sprintf("%f", result.Bachillerato),
			fmt.Sprintf("%f", result.Colegio),
			fmt.Sprintf("%f", result.Educacion),
			fmt.Sprintf("%f", result.Educativo),
			fmt.Sprintf("%f", result.Escolar),
			fmt.Sprintf("%f", result.Estudiante),
			fmt.Sprintf("%f", result.Facultad),
			fmt.Sprintf("%f", result.Institucion),
			fmt.Sprintf("%f", result.Maestro),
			fmt.Sprintf("%f", result.Profesor),
			fmt.Sprintf("%f", result.Universidad),
			fmt.Sprintf("%f", result.Social),
			fmt.Sprintf("%f", result.Trabajo),
			fmt.Sprintf("%f", result.Pais),
			fmt.Sprintf("%f", result.Fisico),
			fmt.Sprintf("%f", result.Transporte),
			fmt.Sprintf("%f", result.Centro),
			fmt.Sprintf("%f", result.Instituto),
			fmt.Sprintf("%f", result.Media),
			fmt.Sprintf("%f", result.Acosador),
			fmt.Sprintf("%f", result.Agresor),
			fmt.Sprintf("%f", result.Testigos),
			fmt.Sprintf("%f", result.Victima),
			fmt.Sprintf("%f", result.Atormentador),
			fmt.Sprintf("%f", result.Bully),
			fmt.Sprintf("%f", result.Complice),
			fmt.Sprintf("%f", result.Grupo),
			fmt.Sprintf("%f", result.Maton),
			fmt.Sprintf("%f", result.Matoneo),
			fmt.Sprintf("%f", result.Perpetrador),
			fmt.Sprintf("%f", result.Persona),
			fmt.Sprintf("%f", result.Padre),
			fmt.Sprintf("%f", result.Universitario),
			fmt.Sprintf("%f", result.Trabajador),
			fmt.Sprintf("%f", result.Mujer),
			fmt.Sprintf("%f", result.Madre),
			fmt.Sprintf("%f", result.Hombre),
			fmt.Sprintf("%f", result.Companero),
			fmt.Sprintf("%f", result.Companera),
			fmt.Sprintf("%f", result.Adulto),
			fmt.Sprintf("%f", result.Espia),
			fmt.Sprintf("%f", result.Supervisor),
			fmt.Sprintf("%f", result.Adolescente),
			fmt.Sprintf("%f", result.Joven),
			fmt.Sprintf("%f", result.Niño),
			fmt.Sprintf("%f", result.Chavo),
			fmt.Sprintf("%f", result.Chico),
			fmt.Sprintf("%f", result.Hijo),
			fmt.Sprintf("%f", result.Infantil),
			fmt.Sprintf("%f", result.Menor),
			fmt.Sprintf("%f", result.Muchacho),
			fmt.Sprintf("%f", result.Nina),
			fmt.Sprintf("%f", result.Reiterado),
			fmt.Sprintf("%f", result.Repetitivo),
			fmt.Sprintf("%f", result.Frecuente),
			fmt.Sprintf("%f", result.Ano),
			fmt.Sprintf("%f", result.Constante),
			fmt.Sprintf("%f", result.Continuo),
			fmt.Sprintf("%f", result.Creciente),
			fmt.Sprintf("%f", result.Dias),
			fmt.Sprintf("%f", result.Mantiene),
			fmt.Sprintf("%f", result.Menudo),
			fmt.Sprintf("%f", result.Mes),
			fmt.Sprintf("%f", result.Periodico),
			fmt.Sprintf("%f", result.Persecucion),
			fmt.Sprintf("%f", result.Perseguir),
			fmt.Sprintf("%f", result.Persistente),
			fmt.Sprintf("%f", result.Recurrente),
			fmt.Sprintf("%f", result.Repeticion),
			fmt.Sprintf("%f", result.Repetido),
			fmt.Sprintf("%f", result.Seguimiento),
			fmt.Sprintf("%f", result.Semanas),
			fmt.Sprintf("%f", result.Tiempo),
			fmt.Sprintf("%f", result.Ocasional),
			fmt.Sprintf("%f", result.Psicoterapia),
			fmt.Sprintf("%f", result.Colaboracion),
			fmt.Sprintf("%f", result.Conciencia),
			fmt.Sprintf("%f", result.Equilibrio),
			fmt.Sprintf("%f", result.Identificacion),
			fmt.Sprintf("%f", result.Mediacion),
			fmt.Sprintf("%f", result.Orientacion),
			fmt.Sprintf("%f", result.Prevencion),
			fmt.Sprintf("%f", result.Sanacion),
			fmt.Sprintf("%f", result.Sensibilizacion),
			fmt.Sprintf("%f", result.Terapia),
			fmt.Sprintf("%f", result.Autoproteccion),
			fmt.Sprintf("%f", result.Establecer),
			fmt.Sprintf("%f", result.Ciberbulling),
			fmt.Sprintf("%f", result.Sextorsion),
			fmt.Sprintf("%f", result.Grooming),
			fmt.Sprintf("%f", result.Ciberviolencia),
			fmt.Sprintf("%f", result.Sexting),
			fmt.Sprintf("%f", result.Invasivo),
			fmt.Sprintf("%f", result.Racial),
			fmt.Sprintf("%f", result.Laboral),
			fmt.Sprintf("%f", result.Pareja),
			fmt.Sprintf("%f", result.Familiar),
			fmt.Sprintf("%f", result.Colectivo),
			fmt.Sprintf("%f", result.Exclusion),
			fmt.Sprintf("%f", result.Suplantacion),
			fmt.Sprintf("%f", result.Denigracion),
			fmt.Sprintf("%f", result.Sonsacamiento),
			fmt.Sprintf("%f", result.Doxxing),
			fmt.Sprintf("%f", result.Ciberstalking),
			fmt.Sprintf("%f", result.Dia),
			fmt.Sprintf("%f", result.Noche),
			fmt.Sprintf("%d", result.Correcto),
		}

		// Escribe los datos al CSV
		err = writer.Write(record)
		if err != nil {
			log.Fatal("No se pudo escribir el registro en el CSV:", err)
		}
	}

	if errCur := cur.Err(); errCur != nil {
		log.Fatal(errCur)
	}

	fmt.Println("Exportación completada.")
}
