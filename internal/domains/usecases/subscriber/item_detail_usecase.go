package subscriber

import (
	"hackerNewsApi/internal/domains/entity"
	"hackerNewsApi/internal/infrastructure/pubsub"
)

type SubscriberItemDetailRepository interface {
	UpdateItem(items entity.Item) error
}

type SubscriberItemDetailUseCase interface {
	SubscriberUpdateDetailTopStory() func(data pubsub.PubSubEvent) error
}
