package specifications

import (
	"github.com/hanhnham91/order-service/entity"
	"gorm.io/gorm"
)

type productByID struct {
	id       int64
	unscoped []bool
}

func ProductByID(id int64, unscoped ...bool) I {
	return &productByID{
		id:       id,
		unscoped: unscoped,
	}
}

func (p *productByID) GormQuery(db *gorm.DB) *gorm.DB {
	if len(p.unscoped) > 0 && p.unscoped[0] {
		db = db.Unscoped()
	}

	return db.Model(&entity.Product{}).
		Where("id = ?", p.id)
}
