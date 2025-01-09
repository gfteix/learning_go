
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



## Executing the requests

- /users/register

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


- /users/login

```
curl -X POST http://localhost:8080/users/login \
-v -H "Content-Type: application/json" \
-d '{
  "email": "john.doe@example.com",
  "password": "password123"
}'
```


- /products

```
curl http://localhost:8080/products -v
```


- /cart/checkout

Updates the <your-auth-token> with the token returned from the /users/login

```
curl -X POST "http://localhost:8080/cart/checkout" \
-H "Authorization: <your-auth-token>" \
-H "Content-Type: application/json" \
-v -d '{
  "items": [
    {
      "ProductID": 1,
      "Quantity": 2
    },
    {
      "ProductID": 2,
      "Quantity": 1
    }
  ]
}'
```