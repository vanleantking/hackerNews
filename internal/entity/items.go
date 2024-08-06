package entity

import "time"

type Item struct {
	ID        string    `gorm:"column:id;primaryKey"`
	HNItemID  int       `gorm:"column:hn_item_id"`
	URL       string    `gorm:"column:url"`
	Text      string    `gorm:"column:city"`
	Title     string    `gorm:"column:title"`
	By        string    `gorm:"column:item_author"`
	ItemType  string    `gorm:"column:item_type"`
	Category  string    `gorm:"column:string"`
	Score     int8      `gorm:"column:score"`
	UpdatedAt int64     `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	CreatedAt time.Time `gorm:"created_at;autoCreateTime:milli"`
}

func (a *Item) TableName() string {
	return "hn_items"
}
