package domain

type TopicsRepository interface {
	GetTopics() ([]Topic, error)
}
