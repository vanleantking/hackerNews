package subscribers

import (
	"hackerNewsApi/internal/domains/entity"
)

type SubscriberItemDetailRepository interface {
	UpdateItem(items entity.Item) error
}

type SubscriberItemDetailUsc interface {
	SubscriberUpdateDetailTopStory() func(data []byte) error
}
