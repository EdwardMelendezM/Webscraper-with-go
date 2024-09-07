package data_transformation

import (
	"log"
	"math"
	"regexp"
	"strings"

	"github.com/kljensen/snowball" // Librería para stemming avanzado
)

// Tokenizer function with punctuation handling, normalization, and stop word removal
func tokenize(text string, stopWords []string) []string {
	// Eliminar puntuación y normalizar texto
	re := regexp.MustCompile(`[^\w\s]`)
	text = re.ReplaceAllString(text, "")
	text = strings.ToLower(text)

	// Tokenizar y eliminar stop words
	words := strings.Fields(text)
	tokens := []string{}
	for _, word := range words {
		word = strings.ToLower(word)
		if !contains(stopWords, word) && len(word) > 1 { // También eliminamos palabras muy cortas
			tokens = append(tokens, word)
		}
	}
	return tokens
}

// Stem function improved with a robust stemming algorithm (snowball stemmer)
func stem(word string, lang string) string {
	// Snowball stemmer for more accurate stemming
	stemmed, err := snowball.Stem(word, lang, true)
	if err != nil {
		log.Printf("Error stemming word '%s': %v", word, err)
		return word // Return original word if stemming fails
	}
	return stemmed
}

// TF (Term Frequency) calculation with slight improvements
func calculateTF(term string, document []string) float64 {
	count := 0
	for _, word := range document {
		if word == term {
			count++
		}
	}
	if len(document) == 0 {
		return 0 // Avoid division by zero
	}
	return float64(count) / float64(len(document))
}

// IDF (Inverse Document Frequency) calculation with log smoothing to avoid division by zero
func calculateIDF(term string, documents [][]string) float64 {
	count := 0
	for _, document := range documents {
		if contains(document, term) {
			count++
		}
	}
	// Smoothing to avoid division by zero and handle edge cases
	return math.Log(float64(len(documents)+1) / float64(count+1))
}

// TF-IDF calculation optimized with safety checks
func calculateTFIDF(term string, document []string, documents [][]string) float64 {
	tf := calculateTF(term, document)
	idf := calculateIDF(term, documents)
	return tf * idf
}

// Helper function to check if a term is contained in a document
func contains(document []string, term string) bool {
	for _, word := range document {
		if word == term {
			return true
		}
	}
	return false
}
