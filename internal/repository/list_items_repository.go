package repository

import (
	"hackerNewsApi/internal/entity"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

type listItemRepository struct {
	Repository[entity.Item]
	Log *logrus.Logger
}

type ListItemRepository interface {
	UpsertListItems(items []entity.Item) []error
	UpsertBulkItems(items []entity.Item) error
}

func NewListItemRepository(log *logrus.Logger) ListItemRepository {
	return &listItemRepository{}
}

func (listItems *listItemRepository) UpsertListItems(items []entity.Item) []error {
	var errs = make([]error, 0)
	for _, item := range items {
		tx := listItems.DB.Where("hn_item_id = ?", item.HNItemID).Save(item)
		if tx.Error != nil {
			errs = append(errs, tx.Error)
		}
		time.Sleep(100 * time.Millisecond)
	}
	return errs
}

func (listItems *listItemRepository) UpsertBulkItems(items []entity.Item) error {
	tx := listItems.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "hn_item_id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"score",
			"updated_at"}),
	}).Create(&items)
	return tx.Error
}
