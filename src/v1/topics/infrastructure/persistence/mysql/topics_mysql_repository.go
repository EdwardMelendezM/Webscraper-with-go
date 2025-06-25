package mysql

import (
	"webscraper-go/v1/topics/domain"
)

type TopicsMysqlRepo struct {
}

func NewTopicsRepository() domain.TopicsRepository {
	return &TopicsMysqlRepo{}
}
