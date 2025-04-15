package specifications

import "gorm.io/gorm"

type I interface {
	GormQuery(db *gorm.DB) *gorm.DB
}
