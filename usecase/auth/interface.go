package auth

import (
	"context"

	"github.com/hanhnham91/order-service/payload"
)

type IAuthLoginWithFirebaseUseCase interface {
	Execute(ctx context.Context, req payload.AuthFirebaseRequest) (string, error)
}
