package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/hanhnham91/order-service/codetype"
	"github.com/hanhnham91/order-service/config"
	"github.com/hanhnham91/order-service/entity"
	"github.com/hanhnham91/order-service/payload"
	"github.com/hanhnham91/order-service/repository/specifications"
	"github.com/hanhnham91/order-service/repository/user"
	"github.com/hanhnham91/order-service/util"
	pkgerror "github.com/hanhnham91/pkg/error"
	"gorm.io/gorm"
)

type authLoginWithFirebaseUseCase struct {
	firebase *auth.Client
	userRepo user.Repository
}

func NewAuthLoginWithFirebaseUseCase(firebase *auth.Client, userRepo user.Repository) IAuthLoginWithFirebaseUseCase {
	return &authLoginWithFirebaseUseCase{
		firebase: firebase,
		userRepo: userRepo,
	}
}

func (u *authLoginWithFirebaseUseCase) Execute(
	ctx context.Context,
	req payload.AuthFirebaseRequest,
) (string, error) {
	cfg := config.GetConfig()

	dataToken, err := u.firebase.VerifyIDToken(ctx, req.TokenID)
	if err != nil {
		return "", pkgerror.ErrBadRequest("Invalid token")
	}

	if dataToken.Claims["email_verified"] == false {
		return "", pkgerror.ErrBadRequest("The email was not verified")
	}

	userEmail := fmt.Sprintf("%v", dataToken.Claims["email"])

	existingUser, err := u.userRepo.Get(ctx, specifications.UserByEmail(userEmail, true))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// create new user
			existingUser = entity.User{
				Email:     userEmail,
				FullName:  fmt.Sprintf("%v", dataToken.Claims["name"]),
				Avatar:    fmt.Sprintf("%v", dataToken.Claims["picture"]),
				Role:      0, // / 0: User; 1: Admin
				CreatedBy: 0,
				Status:    codetype.StatusActive,
			}
		} else {
			return "", pkgerror.ErrInternalServerError(err)
		}
	}

	if existingUser.Status != codetype.StatusActive {
		return "", pkgerror.ErrBadRequest("The user account is inactive")
	}

	if err := u.userRepo.CreateOrDoUpdate(ctx, &existingUser, "full_name", "avatar", "status"); err != nil {
		return "", pkgerror.ErrInternalServerError(err)
	}

	userRole := "user"
	if existingUser.Role == 1 {
		userRole = "admin"
	}

	// create token
	tokenString, err := util.NewJwt().
		EncodeToken(
			util.CustomClaims{
				UserID: existingUser.ID,
				Email:  existingUser.Email,
				Role:   userRole,
			},
			cfg.TokenSecretKey,
			time.Now().Add(24*time.Hour),
		)
	if err != nil {
		return "", pkgerror.ErrInternalServerError(err)
	}

	return tokenString, nil
}
