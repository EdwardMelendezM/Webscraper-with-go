package usecase

import (
	"math"
	"regexp"
	"strings"
)

// Calcular TF-IDF para un conjunto de documentos
func calculateTFIDF(term string, document []string, documents [][]string) float64 {
	tf := calculateTF(term, document)
	idf := calculateIDF(term, documents)
	return tf * idf
}

// Obtener la palabra clave más relevante de un texto
func getKeyword(text string, documents [][]string, stopWords []string) (string, error) {
	// Tokenizar el texto
	tokens := tokenize(text, stopWords)

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

// Helper function to check if a term is contained in a document
func contains(document []string, term string) bool {
	for _, word := range document {
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
a
abajo
acerca
adelante
además
adentro
aquel
aquella
aquello
aquí
arriba
así
atras
bajo
bien
cada
cerca
cierto
como
con
cuál
cuya
dentro
de
del
desde
donde
el
ella
ellos
en
encima
ese
esa
eso
este
esta
esto
fue
gran
hacia
hasta
hay
la
lo
los
me
mi
mis
muy
nada
ni
no
nos
nuestra
nuestro
o
para
pero
por
que
quien
se
sin
sobre
su
sus
también
tan
te
ti
todo
todos
tu
tus
un
una
uno
usted
van
varios
ve
vez
yo
abajo
además
al
algún
alguna
alguno
algunos
amigo
antes
bajo
bien
cada
cerca
cierto
como
cuánto
de
debe
dicha
dicen
donde
estado
estoy
evidentemente
hasta
igual
llevar
luego
más
me
mí
nunca
poco
por
puede
sabemos
sabe
ser
si
sobre
tan
tanto
todas
todo
tu
usted
vale
ver
vosotros
y
ya
acerca
al
algun
alguna
alguno
algunos
bajo
bien
cada
con
cuál
cuál
dentro
donde
es
este
final
fue
grande
gente
hemos
igual
largo
menos
nuestra
nuestro
poca
poco
por
pues
que
quien
según
ser
si
solo
toda
todo
tu
usted
vosotros
y
ya
adicional
agregar
ahora
algunos
apenas
aún
bastante
cada
casi
como
constantemente
debería
después
donde
durante
en
entonces
esos
hasta
independientemente
información
incluso
inmediato
interiormente
nuevamente
últimamente
principalmente
próximamente
pues
rápidamente
realmente
recientemente
repetidamente
resumidamente
sus
tampoco
temprano
tendencia
último
varias
varios
visible
voluntariamente
`
)
