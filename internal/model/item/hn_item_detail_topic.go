package item

import (
	"encoding/json"
	"hackerNewsApi/internal/infrastructure/service/hn_api/common"
	"hackerNewsApi/internal/model"
	"time"
)

type PubsubHNItemDetail struct {
	TopicName   string          `json:"topic_name"`
	ID          uint            `json:"id"`
	HNID        uint            `json:"hn_item_id"`
	Score       int             `json:"item_score"`
	By          string          `json:"item_by"`
	Title       string          `json:"item_title"`
	URL         string          `json:"item_url"`
	ItemText    string          `json:"item_content"`
	DescenDants int             `json:"descendants"`
	Kids        json.RawMessage `json:"kids"`
	ItemType    string          `json:"item_type"`
	ItemDeleted bool            `json:"deleted"`
	ItemStatus  int             `json:"item_status"`
	UpdatedAt   int64           `json:"updated_at"`
}

func (e PubsubHNItemDetail) GetTopicName() string {
	return e.TopicName
}

func MapperItemToPubsubItem(item *model.Item) *PubsubHNItemDetail {
	currentTime := time.Now()
	return &PubsubHNItemDetail{
		TopicName:   common.PB_EVENT_HN_ITEM_DETAIL,
		HNID:        item.ItemID,
		ItemStatus:  common.ITEM_STATUS_PROCESS_TITLE,
		By:          item.ItemBy,
		Score:       item.ItemScore,
		ItemText:    item.ItemText,
		Title:       item.ItemTitle,
		ItemType:    item.ItemType,
		URL:         item.ItemURL,
		DescenDants: int(item.DescenDants),
		Kids:        item.Kids,
		UpdatedAt:   currentTime.Unix(),
		ItemDeleted: item.ItemDelete,
	}
}

func (item *PubsubHNItemDetail) BytePresenter() ([]byte, error) {
	return json.Marshal(item)
}
