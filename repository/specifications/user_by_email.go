package specifications

import (
	"github.com/hanhnham91/order-service/entity"
	"gorm.io/gorm"
)

type userByEmail struct {
	email    string
	unscoped []bool
}

func UserByEmail(email string, unscoped ...bool) I {
	return &userByEmail{
		email:    email,
		unscoped: unscoped,
	}
}

func (p *userByEmail) GormQuery(db *gorm.DB) *gorm.DB {
	if len(p.unscoped) > 0 && p.unscoped[0] {
		db = db.Unscoped()
	}

	return db.Model(&entity.User{}).
		Where("email = ?", p.email)
}
