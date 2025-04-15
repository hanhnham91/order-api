package product

import (
	"github.com/hanhnham91/order-service/entity"
	"github.com/hanhnham91/order-service/repository/specifications"
)

type Repository interface {
	FindAll() ([]entity.Product, error)
	Find(spec specifications.I) ([]entity.Product, error)
	Get(spec specifications.I) (entity.Product, error)
}
