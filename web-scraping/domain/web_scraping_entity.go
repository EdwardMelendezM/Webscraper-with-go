package domain

type SearchResult struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Content string `json:"content"`
}

type CreateRecordWebScraping struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
