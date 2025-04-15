package presenter

import (
	"github.com/hanhnham91/order-service/entity"
)

type ImageResponse struct {
	Mobile    string `json:"mobile"`
	Tablet    string `json:"tablet"`
	Desktop   string `json:"desktop"`
	Thumbnail string `json:"thumbnail"`
}

type ProductResponse struct {
	ID       int64         `json:"id"`
	Name     string        `json:"name"`
	Category string        `json:"category"`
	Price    float64       `json:"price"`
	Image    ImageResponse `json:"image"`
}

func FormProductResponse(product entity.Product) ProductResponse {
	resp := ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Category: product.Category,
		Price:    product.Price,
		Image: ImageResponse{
			Mobile:    product.Image.Mobile,
			Tablet:    product.Image.Tablet,
			Desktop:   product.Image.Desktop,
			Thumbnail: product.Image.Thumbnail,
		},
	}

	return resp
}

func FormProductsResponse(products []entity.Product) []ProductResponse {
	resp := make([]ProductResponse, 0, len(products))
	for _, product := range products {
		resp = append(resp, FormProductResponse(product))
	}

	return resp
}
