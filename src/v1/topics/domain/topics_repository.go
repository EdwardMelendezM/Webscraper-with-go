package domain

type TopicsRepository interface {
	GetTopics(projectId string) ([]Topic, error)
}
