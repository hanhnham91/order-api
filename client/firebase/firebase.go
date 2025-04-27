package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/hanhnham91/order-service/config"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

func GetClient(ctx context.Context) *auth.Client {
	cfg := config.GetConfig()
	opt := option.WithCredentialsFile(cfg.FirebaseAdminSDK)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Error initializing app"))
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Error getting Auth client"))
	}

	return client
}
