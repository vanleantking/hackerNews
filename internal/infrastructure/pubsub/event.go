package pubsub

type PubSubEvent interface {
	GetTopicName() string
}
