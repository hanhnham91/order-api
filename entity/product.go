package entity

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID        int64
	Thumbnail string
	Mobile    string
	Tablet    string
	Desktop   string
}

type Product struct {
	ID       int64
	Name     string
	Category string
	Price    float64
	ImageID  int64
	Image    Image

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt
}
