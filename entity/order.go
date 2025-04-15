package entity

import "time"

type Order struct {
	ID         int64
	CouponCode string
	Amount     float64
	OrderItems []OrderItem

	CreatedAt time.Time
	UpdatedAt time.Time
}
