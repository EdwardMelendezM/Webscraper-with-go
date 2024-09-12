package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
	// Conectar a Neo4j
	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "password", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	// Definir el corpus y la palabra clave
	titleCorpus := "deb,fortalec,atencion,cas,acos,sexual"
	workKey := "puebl"

	// Ejecutar la consulta
	result, err := session.Run(`
        MATCH (t:Title {corpus: $titleCorpus})
        MATCH (k:Keyword {keyword: $workKey})
        MATCH (t)-[:RELACIONADO_CON]->(categoria:Categoria)
        RETURN categoria
    `, map[string]interface{}{
		"titleCorpus": titleCorpus,
		"workKey":     workKey,
	})

	// Procesar resultados
	for result.Next() {
		fmt.Println(result.Record().GetByIndex(0)) // Imprime la categoría encontrada
	}
	if err := result.Err(); err != nil {
		panic(err)
	}
}
