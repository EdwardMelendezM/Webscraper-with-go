package domain

type SearchResult struct {
	Title   string `json:"title"`
	Url     string `json:"url"`
	Content string `json:"content"`
}

type CreateRecordWebScraping struct {
	Title  string `json:"title"`
	Url    string `json:"url"`
	Number int    `json:"number"`
}

type UpdateRecordWebScraping struct {
	Content string `json:"content"`
}

type WebScrapingResult struct {
	Id  string `json:""`
	Url string `json:"https://google.com.pe"`
}
