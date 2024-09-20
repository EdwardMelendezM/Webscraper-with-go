# Web Scraper with Go for Cyberbullying Dataset Creation

![Project Status](https://img.shields.io/badge/Status-Completed-brightgreen)

## 🚀 Project Overview
This project focuses on creating a **Cyberbullying Dataset** by utilizing **advanced web scraping techniques** and **semantic ontologies** to enrich the collected data. The goal is to extract valuable information from the web, process it, and enhance it using ontological concepts to address modern cyberbullying patterns.

## 🛠 Technologies Used
- **Go**: For performing web scraping and handling data extraction.
- **Python**: For data processing (lemmatization) via API.
- **MySQL**: To store raw and processed data.
- **MongoDB**: For storing enriched data with semantic ontologies.
- **Semantic Ontologies**: To classify and enhance data related to cyberbullying.

## 📑 Project Workflow
1. **Web Scraping**:
   - Developed a custom web scraper in Go to gather relevant information about cyberbullying from various websites.
   
2. **Data Cleaning**:
   - Removed unnecessary elements like HTML, CSS, JavaScript tags, and advertisements.
   - Tokenization and stop word removal were applied to the raw text.

3. **Lemmatization**:
   - Implemented lemmatization via an API developed in Python to process the cleaned data.

4. **Ontology Creation**:
   - Built semantic ontologies to categorize and enhance the extracted data based on cyberbullying-related terms.

5. **Data Enrichment**:
   - Enriched the cleaned and processed records using the ontologies, then stored them in a NoSQL MongoDB database for further analysis.

## 💡 Key Features
- **Real-time Web Data Extraction**: Continuously extracts up-to-date cyberbullying information from the web.
- **Data Cleaning Pipeline**: Automated pipeline to clean and process raw web data.
- **Semantic Enrichment**: Enhances data with ontologies to add context and value, making the dataset more relevant for research.

## 🚧 How to Run the Project
### Prerequisites:
- Go 1.16+ installed
- Python 3.x installed
- MySQL and MongoDB instances

### Step-by-Step Instructions:
1. **Clone the repository**:
   ```bash
   git clone https://github.com/EdwardMelendezM/Webscraper-with-go.git
   cd Webscraper-with-go
