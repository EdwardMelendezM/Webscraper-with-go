package mysql

type WebScrapingResult struct {
	Id    string `db:"scraped_result_id"`
	Title string `db:"scraped_result_title"`
	Url   string `db:"scraped_result_url"`
}
