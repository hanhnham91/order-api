package order

import (
	"github.com/hanhnham91/order-service/entity"
)

type Repository interface {
	Create(data *entity.Order) error
}
