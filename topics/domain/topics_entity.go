package domain

import "time"

type Topic struct {
	//Description: the id of the topic
	Id string `json:"id" example:"739bbbc9-7e93-11ee-89fd-0242ac110010"`
	//Description: the title of the topic
	Title string `json:"title" example:"historias cortas de acoso"`
	//Description: the created at of the topic
	CreatedAt *time.Time `json:"created_at" example:"2022-01-01T00:00:00Z"`
}
