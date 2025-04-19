package order

import (
	"context"
	"testing"

	"github.com/hanhnham91/order-service/entity"
	mockorder "github.com/hanhnham91/order-service/mocks/order"
	mockproduct "github.com/hanhnham91/order-service/mocks/product"
	"github.com/hanhnham91/order-service/payload"
	"github.com/hanhnham91/order-service/repository/order"
	"github.com/hanhnham91/order-service/repository/product"
	"github.com/hanhnham91/order-service/repository/specifications"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewOrderCreateUseCase(t *testing.T) {
	productRepo := mockproduct.NewMockRepository(t)
	orderRepo := mockorder.NewMockRepository(t)

	type args struct {
		productRepo product.Repository
		orderRepo   order.Repository
	}

	tests := []struct {
		name string
		args args
		want IOrderCreateUseCase
	}{
		{
			name: "success",
			args: args{
				productRepo: productRepo,
				orderRepo:   orderRepo,
			},
			want: &orderCreateUseCase{
				productRepo: productRepo,
				orderRepo:   orderRepo,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewOrderCreateUseCase(
				tt.args.productRepo,
				tt.args.orderRepo,
			)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_orderCreateUseCase(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		req := payload.CreateOrderRequest{
			CouponCode: "",
			Items: []payload.OrderItem{
				{
					ProductID: 1,
					Quantity:  1,
				},
			},
		}

		productIDs := make([]int64, 0)
		for i := range req.Items {
			productIDs = append(productIDs, req.Items[i].ProductID)
		}

		// Init mocks
		productRepo := mockproduct.NewMockRepository(t)
		orderRepo := mockorder.NewMockRepository(t)

		products := []entity.Product{
			{
				ID:       1,
				Name:     "Name",
				Category: "",
				Price:    10,
			},
		}
		productRepo.EXPECT().
			Find(specifications.ProductByIDs(productIDs)).
			Return(products, nil)

		orderRepo.EXPECT().
			Create(&entity.Order{
				CouponCode: req.CouponCode,
				Amount:     10,
				OrderItems: []entity.OrderItem{
					{
						ProductID: 1,
						Quantity:  1,
					},
				},
			}).
			Return(nil)

		// execute
		uc := orderCreateUseCase{
			productRepo: productRepo,
			orderRepo:   orderRepo,
		}
		_, _, gotErr := uc.Execute(context.Background(), req)

		require.NoError(t, gotErr)
	})
}
