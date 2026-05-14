package mysql

import (
	"database/sql"
	_ "database/sql"
	_ "embed"
	// El driver de pgx debe estar presente para que sql.Open lo encuentre
	_ "github.com/jackc/pgx/v5/stdlib"

	"webscraper-go/topics/domain"

	"github.com/jackskj/carta"
	"github.com/stroiman/go-automapper"

	"webscraper-go/infrastructure/persistence/postgres/db"
)

//go:embed sql/get_topics.sql
var QueryGetTopics string

func (r TopicsMysqlRepo) GetTopics(projectId string) (
	topics []domain.Topic,
	err error,
) {
	results, err := db.ClientV2.Query(
		QueryGetTopics,
		projectId,
	)
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
