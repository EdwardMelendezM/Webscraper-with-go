package mysql

import (
	"database/sql"
	_ "database/sql"
	_ "embed"

	"github.com/EdwardMelendezM/api-info-shared/db"
	"github.com/jackskj/carta"
	"github.com/stroiman/go-automapper"

	"webscraper-go/topics/domain"
)

//go:embed sql/get_topics.sql
var QueryGetTopics string

func (r TopicsMysqlRepo) GetTopics(projectId string) (topics []domain.Topic, err error) {
	results, err := db.Client.Query(QueryGetTopics, projectId)
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			return
		}
	}(results)

	topicsTmp := make([]Topic, 0)
	err = carta.Map(results, &topicsTmp)
	if err != nil {
		return nil, err
	}
	automapper.Map(topicsTmp, &topics)
	return topics, nil

}
