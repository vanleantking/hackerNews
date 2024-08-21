package hnitemdetail

import "hackerNewsApi/internal/infrastructure/service/hn_api/common"

type PubsubHNItemDetail struct {
	ID     uint `json:"id"`
	HNID   uint `json:"hn_item_id"`
	Status int  `json:"status"`
}

func (e *PubsubHNItemDetail) GetTopicName() string {
	return common.PB_EVENT_HN_ITEM_DETAIL
}
