package adapters

import (
	"github.com/kitpk/go-unittest106/core"

	"gorm.io/gorm"
)

// Secondary adapter
type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) core.OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) Save(order core.Order) error {
	if result := r.db.Create(&order); result.Error != nil {
		return result.Error
	}

	return nil
}
