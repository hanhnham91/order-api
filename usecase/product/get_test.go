package product

import (
	"context"
	"errors"
	"testing"

	"github.com/hanhnham91/order-service/entity"
	mockproduct "github.com/hanhnham91/order-service/mocks/product"
	"github.com/hanhnham91/order-service/repository/product"
	"github.com/hanhnham91/order-service/repository/specifications"
	pkgerror "github.com/hanhnham91/pkg/error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestNewProductGetByIDUseCase(t *testing.T) {
	productRepo := mockproduct.NewMockRepository(t)

	type args struct {
		productRepo product.Repository
	}

	tests := []struct {
		name string
		args args
		want IProductGetUseCase
	}{
		{
			name: "success",
			args: args{
				productRepo: productRepo,
			},
			want: &productGetByIDUseCase{
				productRepo: productRepo,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewProductGetByIDUseCase(tt.args.productRepo)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_productGetByIDUseCase_Execute(t *testing.T) {
	t.Parallel()

	t.Run("success - product found", func(t *testing.T) {
		t.Parallel()

		productRepo := mockproduct.NewMockRepository(t)
		expectedProduct := entity.Product{ID: 1, Name: "Test Product", Price: 25.0}

		productRepo.EXPECT().
			Get(specifications.ProductByID(int64(1))).
			Return(expectedProduct, nil)

		uc := productGetByIDUseCase{
			productRepo: productRepo,
		}

		gotProduct, gotErr := uc.Execute(context.Background(), 1)
		require.NoError(t, gotErr)
		assert.Equal(t, expectedProduct, gotProduct)
	})

	t.Run("failure - product not found", func(t *testing.T) {
		t.Parallel()

		productRepo := mockproduct.NewMockRepository(t)

		productRepo.EXPECT().
			Get(specifications.ProductByID(int64(99))).
			Return(entity.Product{}, gorm.ErrRecordNotFound)

		uc := productGetByIDUseCase{
			productRepo: productRepo,
		}

		gotProduct, gotErr := uc.Execute(context.Background(), 99)
		require.Error(t, gotErr)
		assert.Empty(t, gotProduct)
		assert.ErrorIs(t, gotErr, pkgerror.ErrNotFound())
	})

	t.Run("failure - product repository returns other error", func(t *testing.T) {
		t.Parallel()

		productRepo := mockproduct.NewMockRepository(t)
		expectedErr := errors.New("database error")

		productRepo.EXPECT().
			Get(specifications.ProductByID(int64(1))).
			Return(entity.Product{}, expectedErr)

		uc := productGetByIDUseCase{
			productRepo: productRepo,
		}

		gotProduct, gotErr := uc.Execute(context.Background(), 1)
		require.Error(t, gotErr)
		assert.Empty(t, gotProduct)
		assert.ErrorIs(t, gotErr, pkgerror.ErrInternalServerError(expectedErr))
	})
}
