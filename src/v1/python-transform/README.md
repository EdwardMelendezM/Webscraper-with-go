## Install

```bash
sudo apt update
sudo apt install python3.12-venv
```

## Create new env

- Create new one
```bash
# Crear el entorno limpio
python3 -m venv venv
```

- Activate env
```bash
source venv/bin/activate
```

- Install dependencies
```bash
pip install --upgrade pip setuptools wheel
pip install -r requirements.txt
pip install lxml
python3 -m spacy download es_core_news_sm
```