package usecase

import (
	"github.com/bbalet/stopwords"
	"math"
	"regexp"
	"strings"

	"github.com/kljensen/snowball"
)

// Calcular TF-IDF para un conjunto de documentos
func calculateTFIDF(term string, document []string, documents [][]string) float64 {
	tf := calculateTF(term, document)
	idf := calculateIDF(term, documents)
	return tf * idf
}

// Obtener la palabra clave más relevante de un texto
func getKeyword(text string, documents [][]string) (string, error) {
	// Tokenizar el texto
	tokens := tokenize(text)

	if len(tokens) == 0 {
		return "", nil
	}

	// Crear un mapa para almacenar el puntaje TF-IDF de cada palabra
	tfIdfScores := make(map[string]float64)
	for _, token := range tokens {
		score := calculateTFIDF(token, tokens, documents)
		tfIdfScores[token] = score
	}

	// Encontrar la palabra con el mayor puntaje TF-IDF
	var keyword string
	maxScore := math.Inf(-1)
	for word, score := range tfIdfScores {
		if score > maxScore {
			maxScore = score
			keyword = word
		}
	}

	return keyword, nil
}

// Tokenizer function with punctuation handling, normalization, stop word removal, and stemming
func tokenize(text string) []string {
	cleanStopWords := strings.TrimSpace(stopWords)
	stopWordsList := strings.Split(cleanStopWords, "\n")

	textWithoutStopWords := stopwords.CleanString(text, "es", true)

	re := regexp.MustCompile(`[^\p{L}\p{N}\s]`) // Solo mantener letras, números y espacios
	textWithoutStopWords = re.ReplaceAllString(textWithoutStopWords, "")
	textWithoutStopWords = strings.ToLower(textWithoutStopWords)

	words := strings.Fields(textWithoutStopWords)
	tokens := []string{}
	for _, word := range words {
		word = strings.TrimSpace(word)
		// No agregar palabras vacías ni stop words
		if word != "" && !contains(stopWordsList, word) {
			// Aplicar stemming
			stemmed, err := snowball.Stem(word, "spanish", true)
			if err != nil {
				return nil
			}
			tokens = append(tokens, stemmed)
		}
	}
	return tokens
}

// Helper function to check if a term is contained in a list
func contains(list []string, term string) bool {
	for _, word := range list {
		if word == term {
			return true
		}
	}
	return false
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

const (
	removeWords = `\b(descuento|oferta|cupón|outlet|precios|rebajas|promo|publicidad|subscripción|inscripción|newsletter|suscríbete|registrado|iniciar sesión|compartir|copiar enlace|ahora|compra|llama ahora|oferta especial|solicitar información|ganar|comprobar|ver|conocer más|información|al instante)\b|
        \b(publicidad|promoción|anuncio|código|cupón|rebaja|voucher|descuento|venta|compra|anunciar|especial|ahorra)\b|
        \b(facebook|instagram|whatsapp|twitter|email|youtube|redes sociales|linkedin|tiktok|pinterest|snapchat|telegram)\b|
        \b(¡|!|"|\')?\b(esta funcionalidad es sólo para registrados|este contenido está disponible para suscriptores|política de privacidad|contacto|acerca de|sorteo|lotería|solo para suscriptores|contenido exclusivo|solicita ahora|nueva oferta|elige tu plan|suscríbete ahora|aplicación disponible|descarga gratuita)\b|
        \b(premium|suscripción|programa de afiliados|membresía|premium|acceso completo|proveedor|patrocinado|socio|beneficio)\b|
        \b(categoría|noticias|deportes|ocio|cultura|moda|tecnología|salud|economía|finanzas|entretención|actualidad|eventos|artículos|editorial|trending|lo último)\b|
        \b(reporte|error|informar|problema|informe|aviso)\b|
        \b(descargar|app|aplicación|programa|software|actualizar|instalar|abrir|obtener|actualización)\b|
        \b(eventos|concursos|premios|sorteos|juegos|actividades|fiesta|celebración|campaña|sorteo|desafío)\b|
        \b(ingresar|registrarse|iniciar sesión|cerrar sesión|salir|cuenta|perfil|configuración|ajustes|opciones|cuenta personal|administrar|acceder|login|logout)\b|
        \b(ayuda|soporte|preguntas frecuentes|faq|términos y condiciones|condiciones de uso|política de privacidad|asistencia|guía|manual|tutorial)\b|`

	stopWords = `
y
o
n
quotte
accesibilidad
saltar
era
antildeos
en
nº
`
)
