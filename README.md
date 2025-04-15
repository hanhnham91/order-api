# Simple order product api using go
    
## This service included: 
- api method `GET` `/product` to get list product
- api method `GET` `/product/:id` to get product by id
- api method `POST``/order` to order list product with their quantity, and return orderId
```json
## Order request
{
    "couponCode": "",
    "items": [
        {
            "productId": 1,
            "quantity": 1
        }
    ]
}

## Order response
{
    "id": 2,
    "couponCode": "",
    "amount": 6.5,
    "orderItems": [
        {
            "productId": 1,
            "quantity": 1
        }
    ],
    "products": [
        {
            "id": 1,
            "name": "Waffle with Berries",
            "category": "Waffle",
            "price": 6.5,
            "image": {
                "mobile": "",
                "tablet": "",
                "desktop": "",
                "thumbnail": ""
            }
        }
    ]
}
```
## Design:
- using go echo framework.
- relational database with sqlite, can easily replace with mysql, postgre.
- using go validator v10 to validate input
- return error with status code & message: `400` bad request, `404` no found, `500` internal error

## Usage
```bash
## Run
make run

## Built 
make build

## Linter
make lint

## Sec
make sec
````
