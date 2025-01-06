
Learning how to create an ecommerce api with go https://youtu.be/7VLmLOiQ3ck (30:00)

---

Endpoints

- GET /users/:id
- POST /register
- POST /login
- GET /products
- POST /cart/checkout


---

Docker command to build a mysql container

```
docker run -d \
  --name mysql_docker \
  -e MYSQL_ROOT_PASSWORD=mypassword \
  -e MYSQL_DATABASE=ecom \
  -p 3306:3306 \
  mysql
```

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
}'
```