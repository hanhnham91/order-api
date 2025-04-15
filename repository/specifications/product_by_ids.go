package specifications

import (
	"github.com/hanhnham91/order-service/entity"
	"gorm.io/gorm"
)

type productByIDs struct {
	ids []int64
}

func ProductByIDs(ids []int64) I {
	return &productByIDs{
		ids: ids,
	}
}

func (p *productByIDs) GormQuery(db *gorm.DB) *gorm.DB {
	return db.Model(&entity.Product{}).
		Where("id in ?", p.ids)
}
