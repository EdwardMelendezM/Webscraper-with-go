package mysql

import "time"

type ScrapedResult struct {
	Id            string     `db:"id"`
	ProjectId     string     `db:"project_id"`
	Title         string     `db:"title"`
	Url           string     `db:"url"`
	Content       string     `db:"content"`
	Number        int        `db:"number"`
	TitleCorpus   *string    `db:"title_corpus"`
	ContentCorpus *string    `db:"content_corpus"`
	WordKey       *string    `db:"word_key"`
	CreatedAt     *time.Time `db:"created_at"`
}
