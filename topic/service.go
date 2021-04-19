package topic

type TopicService struct {
	topicRepository *TopicRepository
}

func NewTopicService(topicRepository *TopicRepository) *TopicService {
	return &TopicService{
		topicRepository: topicRepository,
	}
}
