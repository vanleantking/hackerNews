package entity

import "encoding/json"

type Item struct {
	ID          uint            `gorm:"column:id;primaryKey"`
	HNItemID    uint            `gorm:"column:hn_item_id"`
	URL         string          `gorm:"column:item_url"`
	Text        string          `gorm:"column:item_content"`
	Title       string          `gorm:"column:item_title"`
	By          string          `gorm:"column:item_by"`
	ItemType    string          `gorm:"column:item_type"`
	Category    uint8           `gorm:"column:category_id"`
	Score       int             `gorm:"column:item_score"`
	CreatedTime int64           `gorm:"column:created_time"`
	DescenDants int             `gorm:"column:descendants"`
	Kids        json.RawMessage `gorm:"column:kids"`
	ItemStatus  int             `gorm:"column:item_status"`
	UpdatedAt   int64           `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	CreatedAt   int64           `gorm:"created_at;autoCreateTime:milli"`
}

func (a *Item) TableName() string {
	return "hn_item"
}
