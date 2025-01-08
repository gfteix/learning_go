
It's a convention to store the entry points of the application in the cmd folder.


## Docker

### Useful docker commands

- To build a mysql container:

```
docker run -d \
  --name mysql_docker \
  -e MYSQL_ROOT_PASSWORD=mypassword \
  -e MYSQL_DATABASE=ecom \
  -p 3306:3306 \
  mysql
```


- To stop the container:

`docker container stop mysql_docker`

- To remove the container:

`docker container remove mysql_docker`

- To go inside the container:

`docker exec -it mysql_docker mysql -uroot -p` (where “root” is the username for MySQL database.)

Select the database

`USE ecom`

Get List of tables

`show tables;`

Run any query

`SELECT * FROM users;`

---

## Env variables

The project is using [godotenv](https://pkg.go.dev/github.com/joho/godotenv@v1.5.1) to load env variables from a .env file