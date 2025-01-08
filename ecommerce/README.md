
Learning how to create an ecommerce api with go https://youtu.be/7VLmLOiQ3ck (30:00)

---

Endpoints

- GET /users/:id
- POST /register
- POST /login
- GET /products
- POST /cart/checkout


---


Examples:

- /register

```
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{
  "firstName": "John",
  "lastName": "Doe",
  "email": "john.doe@example.com",
  "password": "password123"
}' -v
```

---

To create a DB migration run `make migration name_of_the_migration`

---