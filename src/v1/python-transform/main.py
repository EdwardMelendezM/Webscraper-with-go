from flask import Flask, request, jsonify
from bs4 import BeautifulSoup
import spacy
import string

# Optimización: Cargamos solo lo necesario (tokenizer y lemmatizer)
# Excluimos el parser y ner para ahorrar memoria y CPU
nlp = spacy.load('es_core_news_sm', exclude=['parser', 'ner'])

app = Flask(__name__)

def clean_html(html_content):
    if not html_content:
        return ""

    soup = BeautifulSoup(html_content, 'html.parser')

    for script_or_style in soup(['script', 'style', 'header', 'footer', 'nav']):
        script_or_style.decompose() # decompose es ligeramente más eficiente que extract

    # Usamos separator para evitar que palabras queden pegadas
    text = soup.get_text(separator=' ')

    # Limpieza de espacios en blanco usando join/split (más pythonic)
    return " ".join(text.split())

def process_text(text):
    # nlp.pipe es más rápido para textos grandes, pero para strings cortos
    # nlp(text) está bien. Aquí usamos disable para estar seguros.
    doc = nlp(text)

    # Lematización y filtrado en una sola pasada
    # Añadimos limpieza de saltos de línea y espacios extras
    tokens = [
        token.lemma_.lower()
        for token in doc
        if not token.is_stop
           and not token.is_punct
           and not token.is_space
           and token.lemma_.strip() not in string.punctuation
    ]

    return ' '.join(tokens)

@app.route('/clean-corpus', methods=['POST'])
def clean_corpus():
    data = request.get_json(silent=True)
    if not data or 'content' not in data:
        return jsonify({'error': 'Falta el campo content'}), 400

    html_content = data.get('content', '')

    # Pipeline de procesamiento
    clean_text = clean_html(html_content)
    corpus = process_text(clean_text)

    return jsonify({
        'corpus': corpus,
        'length_original': len(html_content),
        'length_clean': len(corpus)
    })

if __name__ == '__main__':
    # host='0.0.0.0' está perfecto para Docker
    app.run(host='0.0.0.0', port=5000, debug=False)