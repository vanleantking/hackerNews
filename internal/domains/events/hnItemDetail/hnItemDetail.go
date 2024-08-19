package hnitemdetail

import "hackerNewsApi/internal/infrastructure/service/hn_api/common"

type EventHNItemDetail struct {
	ID     uint `json:"id"`
	HNID   uint `json:"hn_item_id"`
	Status int  `json:"status"`
}

func (e *EventHNItemDetail) GetName() string {
	return common.PB_EVENT_HN_ITEM_DETAIL
}
