package product

import (
	"context"
	"errors"
	"testing"

	"github.com/hanhnham91/order-service/entity"
	mockproduct "github.com/hanhnham91/order-service/mocks/product"
	"github.com/hanhnham91/order-service/repository/product"
	pkgerror "github.com/hanhnham91/pkg/error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewProductFindUseCase(t *testing.T) {
	productRepo := mockproduct.NewMockRepository(t)

	type args struct {
		productRepo product.Repository
	}

	tests := []struct {
		name string
		args args
		want IProductFindUseCase
	}{
		{
			name: "success",
			args: args{
				productRepo: productRepo,
			},
			want: &productFindUseCase{
				productRepo: productRepo,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewProductFindUseCase(tt.args.productRepo)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_productFindUseCase_Execute(t *testing.T) {
	t.Parallel()

	t.Run("success - returns multiple products", func(t *testing.T) {
		t.Parallel()
		productRepo := mockproduct.NewMockRepository(t)
		expectedProducts := []entity.Product{
			{ID: 1, Name: "Product A", Price: 10},
			{ID: 2, Name: "Product B", Price: 20},
		}

		productRepo.EXPECT().
			FindAll().
			Return(expectedProducts, nil)

		uc := productFindUseCase{
			productRepo: productRepo,
		}

		gotProducts, gotErr := uc.Execute(context.Background())
		require.NoError(t, gotErr)
		require.Len(t, gotProducts, len(expectedProducts))
		assert.Equal(t, expectedProducts, gotProducts)
	})

	t.Run("success - returns no products", func(t *testing.T) {
		t.Parallel()
		productRepo := mockproduct.NewMockRepository(t)
		productRepo.EXPECT().
			FindAll().
			Return([]entity.Product{}, nil)

		uc := productFindUseCase{
			productRepo: productRepo,
		}

		gotProducts, gotErr := uc.Execute(context.Background())
		require.NoError(t, gotErr)
		require.Empty(t, gotProducts)
	})

	t.Run("failure - product repository returns error", func(t *testing.T) {
		t.Parallel()
		productRepo := mockproduct.NewMockRepository(t)
		expectedErr := errors.New("database error")

		productRepo.EXPECT().
			FindAll().
			Return(nil, expectedErr)

		uc := productFindUseCase{
			productRepo: productRepo,
		}

		gotProducts, gotErr := uc.Execute(context.Background())
		require.Error(t, gotErr)
		assert.Nil(t, gotProducts)
		assert.ErrorIs(t, gotErr, pkgerror.ErrInternalServerError(expectedErr))
	})
}
