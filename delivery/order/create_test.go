package order

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hanhnham91/order-service/entity"
	mockorderuc "github.com/hanhnham91/order-service/mocks/usecase/order"
	"github.com/hanhnham91/order-service/payload"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRouteOrder_create(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		orderUseCase := mockorderuc.NewMockIOrderCreateUseCase(t)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{
			"couponCode": "",
			"items": [
				{"productId": 1, "quantity": 2}
			]
		}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		orderUseCase.EXPECT().Execute(mock.Anything, payload.CreateOrderRequest{
			CouponCode: "",
			Items: []payload.OrderItem{
				{
					ProductID: 1,
					Quantity:  2,
				},
			},
		}).Return(&entity.Order{
			ID:         1,
			CouponCode: "",
			OrderItems: []entity.OrderItem{
				{
					ProductID: 1,
					Quantity:  2,
				},
			},
			Amount: 200,
		}, []entity.Product{}, nil)

		r := NewOrderRoute(orderUseCase)
		err := r.create(c)

		assert.NoError(t, err)
	})
}
