package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/hanhnham91/order-service/client/sql"
	"github.com/hanhnham91/order-service/config"
	serviceHTTP "github.com/hanhnham91/order-service/delivery"
	"github.com/hanhnham91/order-service/entity"
	"github.com/soheilhy/cmux"
)

func init() {
	db := sql.GetClient

	if !db().Migrator().HasTable(&entity.User{}) {
		err := db().Migrator().CreateTable(&entity.User{})
		if err != nil {
			log.Fatal(err)
		}
	}

	if !db().Migrator().HasTable(&entity.Product{}) {
		err := db().Migrator().CreateTable(&entity.Product{})
		if err != nil {
			log.Fatal(err)
		}

		err = db().Migrator().CreateTable(&entity.Image{})
		if err != nil {
			log.Fatal(err)
		}

		images := []entity.Image{
			{
				ID:        1,
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-waffle-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-waffle-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-waffle-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-waffle-desktop.jpg",
			},
			{
				ID:        2,
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-desktop.jpg",
			},
			{
				ID:        3,
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-macaron-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-macaron-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-macaron-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-macaron-desktop.jpg",
			},
			{
				ID:        4,
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-tiramisu-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-tiramisu-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-tiramisu-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-tiramisu-desktop.jpg",
			},
			{
				ID:        5,
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-baklava-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-baklava-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-baklava-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-baklava-desktop.jpg",
			},
			{
				ID:        6,
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-meringue-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-meringue-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-meringue-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-meringue-desktop.jpg",
			},
			{
				ID:        7,
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-cake-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-cake-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-cake-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-cake-desktop.jpg",
			},
			{
				ID:        8,
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-brownie-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-brownie-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-brownie-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-brownie-desktop.jpg",
			},
			{
				ID:        9,
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-desktop.jpg",
			},
		}
		// Create initial records
		products := []entity.Product{
			{
				ID:       1,
				Name:     "Waffle with Berries",
				Category: "Waffle",
				Price:    6.5,
				ImageID:  1,
			},
			{
				ID:       2,
				Name:     "Vanilla Bean Crème Brûlée",
				Category: "Crème Brûlée",
				Price:    7,
				ImageID:  2,
			},
			{
				ID:       3,
				Name:     "Macaron Mix of Five",
				Category: "Macaron",
				Price:    8,
				ImageID:  3,
			},
			{
				ID:       4,
				ImageID:  4,
				Name:     "Classic Tiramisu",
				Category: "Tiramisu",
				Price:    5.5,
			},
			{
				ID:       5,
				ImageID:  5,
				Name:     "Pistachio Baklava",
				Category: "Baklava",
				Price:    4,
			},
			{
				ID:       6,
				ImageID:  6,
				Name:     "Lemon Meringue Pie",
				Category: "Pie",
				Price:    5,
			},
			{
				ID:       7,
				ImageID:  7,
				Name:     "Red Velvet Cake",
				Category: "Cake",
				Price:    4.5,
			},
			{
				ID:       8,
				ImageID:  8,
				Name:     "Salted Caramel Brownie",
				Category: "Brownie",
				Price:    4.5,
			},
			{
				ID:       9,
				ImageID:  9,
				Name:     "Vanilla Panna Cotta",
				Category: "Panna Cotta",
				Price:    6.5,
			},
		}

		// Use the Create function to insert the records
		if err = db().Create(&products).Error; err != nil {
			log.Fatal(err)
		}

		if err = db().Create(&images).Error; err != nil {
			log.Fatal(err)
		}
	}

	if !db().Migrator().HasTable(&entity.Order{}) {
		err := db().Migrator().CreateTable(&entity.Order{})
		if err != nil {
			log.Fatal(err)
		}

		err = db().Migrator().CreateTable(&entity.OrderItem{})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	cfg := config.GetConfig()

	l, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatal(err)
	}

	var (
		errs = make(chan error)
		m    = cmux.New(l)
		h    = serviceHTTP.NewHTTPHandler()
	)

	go func() {
		h.Listener = m.Match(cmux.HTTP1Fast(http.MethodPatch))

		log.Printf("Server is running on http://localhost:%s", cfg.Port)

		errs <- h.Start("")
	}()
	go func() {
		errs <- m.Serve()
	}()

	// graceful
	{
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-c

		log.Print("The server is stopping ...")

		_ = h.Shutdown(context.Background())

		close(c)
	}

	log.Printf("exit - %s", (<-errs).Error())
}
