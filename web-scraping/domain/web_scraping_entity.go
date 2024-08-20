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
