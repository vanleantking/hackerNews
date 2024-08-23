package pubsub

type PubSubEvent struct {
	TopicName string      `json:"topic_name"`
	Data      interface{} `json:"data"`
}

func (e PubSubEvent) GetTopicName() string {
	return e.TopicName
}

/**
 * Subcripbe for list
 * map[topic]handler : handler func(data []byte) error
 */
type HandlerFunc func(data []byte) error
