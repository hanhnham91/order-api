package entity

import (
	"time"

	"github.com/hanhnham91/order-service/codetype"
	"gorm.io/gorm"
)

type User struct {
	ID         int64
	Email      string
	GGID       *string
	FullName   string
	Role       int8
	Gender     int64
	Password   string
	Avatar     string
	IsVerified bool
	Status     codetype.Status
	CreatedBy  int64
	UpdatedBy  *int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt
}
