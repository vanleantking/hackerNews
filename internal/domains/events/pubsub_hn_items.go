package events

import (
	"hackerNewsApi/internal/infrastructure/service/hn_api/common"
	"time"
)

type PubSubHNItem struct {
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}

func (pbItem *PubSubHNItem) GetName() string {
	return common.PB_EVENT_HN_ITEM_DETAIL
}
