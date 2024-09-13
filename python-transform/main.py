from flask import Flask, request, jsonify
from bs4 import BeautifulSoup
import spacy
import string

# Cargar el modelo en español
nlp = spacy.load('es_core_news_sm')

app = Flask(__name__)

# Función para limpiar el HTML y extraer solo el texto
def clean_html(html_content):
    soup = BeautifulSoup(html_content, 'html.parser')

    # Eliminar scripts y estilos
    for script_or_style in soup(['script', 'style']):
        script_or_style.extract()

    # Obtener el texto limpio
    text = soup.get_text()

    # Dividir en líneas y unir
    lines = (line.strip() for line in text.splitlines())
    chunks = (phrase.strip() for line in lines for phrase in line.split("  "))
    text = '\n'.join(chunk for chunk in chunks if chunk)

    return text

# Función para procesar el texto: tokenización, eliminación de stopwords y lematización
def process_text(text):
    doc = nlp(text)

    # Tokenización, eliminación de stopwords, puntuación y lematización
    tokens = [token.lemma_.lower() for token in doc if not token.is_stop and token.lemma_ not in string.punctuation]

    # Devolver el corpus como texto limpio
    return ' '.join(tokens)

@app.route('/clean-corpus', methods=['POST'])
def clean_corpus():
    data = request.json
    html_content = data.get('content', '')

    # Limpiar el contenido HTML
    clean_text = clean_html(html_content)

    # Procesar el texto: tokenización, eliminación de stopwords y lematización
    corpus = process_text(clean_text)

    return jsonify({'corpus': corpus})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)