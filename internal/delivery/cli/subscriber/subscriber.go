package subscriber

import (
	"hackerNewsApi/internal/domains/pubsub"
	"hackerNewsApi/internal/domains/usecases"
	event "hackerNewsApi/internal/infrastructure/pubsub"
)

type subscribeProcessing struct {
	ItemDetailUsc usecases.ItemDetailUseCase
	PubSub        pubsub.SubsriberBus
}

type SubscribeProcessing interface {
	ProcessSubscribeHandlers() func(event.PubSubEvent) error
}
