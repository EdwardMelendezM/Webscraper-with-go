package main

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

func main() {
	// Configurar conexión
	uri := "neo4j://localhost:7687"
	username := "neo4j"
	password := "testpassword"

	ctx := context.Background()

	driver, err := neo4j.NewDriverWithContext(
		uri,
		neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Fatal("Error al crear el driver:", err)
	}
	defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		log.Fatal("Error de conectividad:", err)
	}
	fmt.Println("Conexión establecida.")

	// Iniciar sesión
	session := driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: "neo4j", // Si tienes una base de datos específica, reemplázala aquí
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close(ctx)

	// Parámetros
	titleCorpus := "Euforia,Deseo de no haber nacido"
	workKey := "Acoso"

	// Consulta
	query := `
	MATCH (s:Signal)-[:ASSOCIATED_WITH]->(r:RiskFactor)
	WHERE s.name CONTAINS $titleCorpus OR s.name CONTAINS $workKey
	RETURN s.name AS signal, r.name AS riskFactor
	`

	result, err := session.Run(ctx, query, map[string]any{
		"titleCorpus": titleCorpus,
		"workKey":     workKey,
	})
	if err != nil {
		log.Fatal("Error en la consulta:", err)
	}

	lastResult, err := result.Single(ctx)
	if err != nil {
		log.Fatal("Error al obtener el resultado:", err)
	}

	fmt.Printf("%v está asociado con %v.\n", lastResult.Keys)
	fmt.Printf("%v está asociado con %v.\n", lastResult.Values)
}
