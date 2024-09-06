package mysql

import "time"

type Topic struct {
	Id        string     `db:"topic_id"`
	Title     string     `db:"topic_title"`
	CreatedAt *time.Time `db:"topic_created_at"`
}
