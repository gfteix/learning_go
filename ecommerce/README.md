
Learning how to create an ecommerce api with go https://youtu.be/7VLmLOiQ3ck (31:24)


Endpoints

- GET /users/:id
- POST /register
- POST /login
- GET /products
- POST /cart/checkout




Docker command to build a mysql container

```
docker run -d \
  --name mysql_docker \
  -e MYSQL_ROOT_PASSWORD=mypassword \
  -e MYSQL_DATABASE=ecom \
  -p 3306:3306 \
  mysql
```
