package mysql

import "webscraper-go/topics/domain"

type TopicsMysqlRepo struct {
}

func NewTopicsRepository() domain.TopicsRepository {
	return &TopicsMysqlRepo{}
}
