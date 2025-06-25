package domain

type SearchResult struct {
	Title   string `json:"title"`
	Url     string `json:"url"`
	Content string `json:"content"`
	Path    string `json:"path"`
}

type CreateRecordWebScraping struct {
	Title         string `json:"title"`
	Url           string `json:"url"`
	Content       string `json:"content"`
	Number        int    `json:"number"`
	TitleCorpus   string `json:"title_corpus"`
	ContentCorpus string `json:"content_corpus"`
	WordKey       string `json:"word_key"`
}

type UpdateRecordWebScraping struct {
	Content string `json:"content"`
}

type WebScrapingResult struct {
	Id    string `json:"c68a81dc-623e-11ef-9d2f-0242ac110002"`
	Title string `json:"title"`
	Url   string `json:"https://google.com.pe"`
}

type NewRecordWebScraping struct {
	Title   string `json:"title"`
	Url     string `json:"url"`
	Content string `json:"content"`
}
