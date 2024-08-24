package migrate_topics

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func migrateTopic() {
	topics := []string{
		"historias cortas de acoso",
		"relatos de víctimas de acoso",
		"historias de acoso laboral",
		"testimonios de acoso escolar",
		"experiencias de acoso sexual",
		"historias de acoso en redes sociales",
		"foros de víctimas de acoso",
		"blogs sobre acoso",
		"artículos periodísticos sobre acoso",
		"casos reales de acoso",
		"testimonios de acoso en línea",
		"historias de acoso psicológico",
		"narraciones de acoso entre compañeros",
		"relatos de acoso en el trabajo",
		"experiencias personales de acoso",
		"casos de acoso documentados",
		"historias de bullying en escuelas",
		"testimonios de acoso en universidades",
		"experiencias de acoso en el transporte público",
		"historias de acoso cibernético",
		"relatos de hostigamiento sexual",
		"crónicas de acoso por internet",
		"testimonios de víctimas de stalking",
		"historias sobre acoso emocional",
		"experiencias de acoso entre adolescentes",
		"casos de acoso en la calle",
		"narraciones de acoso en el deporte",
		"relatos de acoso en comunidades virtuales",
		"testimonios de acoso en relaciones de pareja",
		"historias sobre acoso en centros educativos",
		"historias de acoso a menores de edad",
		"relatos de acoso a adultos mayores",
		"testimonios de acoso en barrios",
		"experiencias de acoso en plazas públicas",
		"historias de acoso en centros comerciales",
		"relatos de acoso en gimnasios",
		"testimonios de acoso en academias",
		"experiencias de acoso en universidades",
		"historias de acoso en colegios",
		"testimonios de acoso a mujeres adultas",
		"historias de acoso en centros de trabajo",
		"casos de acoso en áreas recreativas",
		"experiencias de acoso en lugares públicos",
		"relatos de acoso entre adultos en espacios laborales",
		"historias de acoso en centros culturales",
		"testimonios de acoso en instituciones educativas",
		"historias de acoso entre adolescentes en colegios",
		"relatos de acoso en espacios deportivos",
		"testimonios de acoso en la comunidad",
		"historias de acoso en parques",
		"casos de acoso en gimnasios y centros de fitness",
	}
	dsn := "root:secret@tcp(127.0.0.1:3309)/acosoDB"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Verificar la conexión
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexión exitosa a la base de datos.")
	now := time.Now()
	for _, topic := range topics {
		//Insert into database
		id := uuid.New().String()
		projectId := "91da2ca7-6244-11ef-9d2f-0242ac110002"
		_, errEx := db.Exec(
			"INSERT INTO scraped_topics (id, project_id,title,created_at) VALUES (?,?,?,?)",
			id,
			projectId,
			topic,
			now,
		)

		if errEx != nil {
			fmt.Println("Error: ", errEx)
			panic(err)
		}

	}
}
