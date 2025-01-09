
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

To stop delete and prune the volumes from docker:

docker container stop mysql_docker && docker container rm mysql_docker && docker volume prune

---

## Env variables

The project is using [godotenv](https://pkg.go.dev/github.com/joho/godotenv@v1.5.1) to load env variables from a .env file


## Migrations

---

- To create a DB migration run `make migration name_of_the_migration`

- Then update the "up" file with the sql command to implement a new change to the database

- And update the "down" file with the sql command to revent the change

- Run make migrate-up to execute the migration or make migration-down to revert it

---


- [ ] Start using transactions

- [X] Use Bearer in the authorization header, and makes the api parses that accodingly

- [X] Implement the user addresses feature. This way we can store the user's address and use it in the checkout instead of an hardcoded value.

- [X]  Implement the user's order history. This way we can store the user's orders and show them in the user's profile.

- [X] Implement the cancel order endpoint. This way we can allow the user to cancel an order if it's not yet shipped.

- [ ] Add tests for new files


/*

For each productId in the payload, retrieve item_ids from the database:

SELECT id 
FROM items 
WHERE product_id = ? AND status = 'available' 
LIMIT ? 
FOR UPDATE;

Validate that the number of available item_ids matches the quantity.
Reserve the required items by updating their status to sold during the same transaction.
Create Order



- Get Products should have a SUM of available items subquery on it
- Accept address for orders

With addressId:

    Validate that the addressId belongs to the customer.
    Fetch the address details from the database to associate it with the order.

With a New Address:

    Validate the address fields.
    Save the new address to the database (optional: associate it with the customer for future use).
    Use the saved address ID for order creation.

Error Handling:

    If neither addressId nor address is provided, return an error.
    If both are provided, prioritize one based on business logic (e.g., prefer addressId or reject the request as ambiguous).

Validate that the addressId provided belongs to the current user


*/