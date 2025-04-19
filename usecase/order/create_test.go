package order

import (
	"context"
	"errors"
	"testing"

	"github.com/hanhnham91/order-service/entity"
	mockorder "github.com/hanhnham91/order-service/mocks/order"
	mockproduct "github.com/hanhnham91/order-service/mocks/product"
	"github.com/hanhnham91/order-service/payload"
	"github.com/hanhnham91/order-service/repository/order"
	"github.com/hanhnham91/order-service/repository/product"
	"github.com/hanhnham91/order-service/repository/specifications"
	pkgerror "github.com/hanhnham91/pkg/error"
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

func Test_orderCreateUseCase_Execute(t *testing.T) {
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

		productRepo := mockproduct.NewMockRepository(t)
		orderRepo := mockorder.NewMockRepository(t)

		products := []entity.Product{
			{
				ID:    1,
				Name:  "Name",
				Price: 10,
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

		uc := orderCreateUseCase{
			productRepo: productRepo,
			orderRepo:   orderRepo,
		}
		gotOrder, gotProducts, gotErr := uc.Execute(context.Background(), req)

		require.NoError(t, gotErr)
		require.NotNil(t, gotOrder)
		require.Len(t, gotProducts, 1)
		assert.Equal(t, req.CouponCode, gotOrder.CouponCode)
		assert.Equal(t, float64(10), gotOrder.Amount)
		assert.Len(t, gotOrder.OrderItems, 1)
		assert.Equal(t, int64(1), gotOrder.OrderItems[0].ProductID)
		assert.Equal(t, 1, gotOrder.OrderItems[0].Quantity)
		assert.Equal(t, products, gotProducts)
	})

	t.Run("failure - find product error", func(t *testing.T) {
		t.Parallel()

		req := payload.CreateOrderRequest{
			Items: []payload.OrderItem{
				{ProductID: 1, Quantity: 1},
			},
		}

		productIDs := []int64{1}

		productRepo := mockproduct.NewMockRepository(t)
		orderRepo := mockorder.NewMockRepository(t)

		expectedErr := errors.New("errors")
		productRepo.EXPECT().
			Find(specifications.ProductByIDs(productIDs)).
			Return(nil, expectedErr)

		uc := orderCreateUseCase{
			productRepo: productRepo,
			orderRepo:   orderRepo,
		}
		gotOrder, gotProducts, gotErr := uc.Execute(context.Background(), req)

		require.Error(t, gotErr)
		assert.Nil(t, gotOrder)
		assert.Nil(t, gotProducts)
		assert.ErrorIs(t, gotErr, pkgerror.ErrInternalServerError(expectedErr))
	})

	t.Run("failure - product not found", func(t *testing.T) {
		t.Parallel()

		req := payload.CreateOrderRequest{
			Items: []payload.OrderItem{
				{ProductID: 99, Quantity: 1},
			},
		}

		productIDs := []int64{99}

		productRepo := mockproduct.NewMockRepository(t)
		orderRepo := mockorder.NewMockRepository(t)

		productRepo.EXPECT().
			Find(specifications.ProductByIDs(productIDs)).
			Return([]entity.Product{}, nil)

		uc := orderCreateUseCase{
			productRepo: productRepo,
			orderRepo:   orderRepo,
		}
		gotOrder, gotProducts, gotErr := uc.Execute(context.Background(), req)

		require.Error(t, gotErr)
		assert.Nil(t, gotOrder)
		assert.Nil(t, gotProducts)
		assert.ErrorIs(t, gotErr, pkgerror.ErrNotFound("Product not found"))
	})

	t.Run("failure - order repository returns error on create", func(t *testing.T) {
		t.Parallel()

		req := payload.CreateOrderRequest{
			Items: []payload.OrderItem{
				{ProductID: 1, Quantity: 1},
			},
		}

		productIDs := []int64{1}

		productRepo := mockproduct.NewMockRepository(t)
		orderRepo := mockorder.NewMockRepository(t)

		products := []entity.Product{
			{
				ID:    1,
				Name:  "Name",
				Price: 10,
			},
		}
		productRepo.EXPECT().
			Find(specifications.ProductByIDs(productIDs)).
			Return(products, nil)

		expectedErr := errors.New("database error")
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
			Return(expectedErr)

		uc := orderCreateUseCase{
			productRepo: productRepo,
			orderRepo:   orderRepo,
		}
		gotOrder, gotProducts, gotErr := uc.Execute(context.Background(), req)

		require.Error(t, gotErr)
		assert.Nil(t, gotOrder)
		assert.Nil(t, gotProducts)
		assert.ErrorIs(t, gotErr, pkgerror.ErrInternalServerError(expectedErr))
	})

	t.Run("success - multiple items", func(t *testing.T) {
		t.Parallel()

		req := payload.CreateOrderRequest{
			Items: []payload.OrderItem{
				{ProductID: 1, Quantity: 2},
				{ProductID: 2, Quantity: 1},
			},
		}

		productIDs := []int64{1, 2}

		productRepo := mockproduct.NewMockRepository(t)
		orderRepo := mockorder.NewMockRepository(t)

		products := []entity.Product{
			{ID: 1, Name: "Product A", Price: 5},
			{ID: 2, Name: "Product B", Price: 10},
		}
		productRepo.EXPECT().
			Find(specifications.ProductByIDs(productIDs)).
			Return(products, nil)

		orderRepo.EXPECT().
			Create(&entity.Order{
				CouponCode: req.CouponCode,
				Amount:     20, // (2 * 5) + (1 * 10)
				OrderItems: []entity.OrderItem{
					{ProductID: 1, Quantity: 2},
					{ProductID: 2, Quantity: 1},
				},
			}).
			Return(nil)

		uc := orderCreateUseCase{
			productRepo: productRepo,
			orderRepo:   orderRepo,
		}
		gotOrder, gotProducts, gotErr := uc.Execute(context.Background(), req)

		require.NoError(t, gotErr)
		require.NotNil(t, gotOrder)
		require.Len(t, gotProducts, 2)
		assert.Equal(t, float64(20), gotOrder.Amount)
		assert.Len(t, gotOrder.OrderItems, 2)
		assert.Contains(t, gotOrder.OrderItems, entity.OrderItem{ProductID: 1, Quantity: 2})
		assert.Contains(t, gotOrder.OrderItems, entity.OrderItem{ProductID: 2, Quantity: 1})
		assert.ElementsMatch(t, products, gotProducts)
	})
}
