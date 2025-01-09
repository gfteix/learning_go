
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

- GET /users/{id}/addresses

Update the {id} accordingly

```
curl "http://localhost:8080/users/{id}/addresses" \
-H "Authorization: Bearer <your-auth-token>"
```

- POST /users/{id}/addresses

```
curl -X POST "http://localhost:8080/users/{id}/addresses" \
-H "Authorization: Bearer <your-auth-token>" \
-H "Content-Type: application/json" \
-v -d '{
  "street": "street",
  "country": "country",
  "postalCode": "postalCode",
  "city": "city",
  "state": "state"
}'
```

- POST /cart/checkout

Updates the <your-auth-token> with the token returned from the /users/login

```
curl -X POST "http://localhost:8080/cart/checkout" \
-H "Authorization: Bearer <your-auth-token>" \
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
curl -X PATCH "http://localhost:8080/orders/3/status" \
-H "Authorization: Bearer <your-auth-token>" \
-H "Content-Type: application/json" \
-v -d '{
  "status": "cancelled"
}'
```

- GET /orders

```
curl "http://localhost:8080/orders" \
-H "Authorization: Bearer <your-auth-token>"
```
