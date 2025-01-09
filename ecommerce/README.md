
Learning how to create an ecommerce api with GO https://youtu.be/7VLmLOiQ3ck

---

## Endpoints

- POST /users/register
- POST /users/login
- GET /users/:id
- GET /products
- POST /cart/checkout


---

## Starting the server

`make run`

## Executing the requests

- POST /users/register

```
curl -X POST http://localhost:8080/users/register \
-H "Content-Type: application/json" \
-d '{
  "firstName": "John",
  "lastName": "Doe",
  "email": "john.doe@example.com",
  "password": "password123"
}' -v
```


- POST /users/login

```
curl -X POST http://localhost:8080/users/login \
-v -H "Content-Type: application/json" \
-d '{
  "email": "john.doe@example.com",
  "password": "password123"
}'
```


- GET /products

```
curl http://localhost:8080/products -v
```


- POST /cart/checkout

Updates the <your-auth-token> with the token returned from the /users/login

```
curl -X POST "http://localhost:8080/cart/checkout" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkQXQiOjE3MzY1NDQ0MzcsInVzZXJJRCI6IjEifQ.RheIuNCQbv0qlAEo4ABco32gQCZriJkafbcu1Du3e1s" \
-H "Content-Type: application/json" \
-v -d '{
  "items": [
    {
      "productID": 1,
      "quantity": 2
    },
    {
      "productID": 2,
      "quantity": 1
    }
  ],
  "address": {
    "street": "street",
    "country": "country",
    "postalCode": "postalCode",
    "city": "city",
    "state": "state"
  }
}'
```

- PATCH /orders/{orderId}/status

Update the {orderId} accordingly

```
curl -X POST "http://localhost:8080/orders/{orderId}/status" \
-H "Content-Type: application/json" \
-v -d '{
  "status": "cancelled"
}'
```

- GET /orders

```
curl "http://localhost:8080/orders" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkQXQiOjE3MzY1NDQ0MzcsInVzZXJJRCI6IjEifQ.RheIuNCQbv0qlAEo4ABco32gQCZriJkafbcu1Du3e1s"
```