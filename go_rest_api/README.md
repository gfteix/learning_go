Learning how to build an rest api with go https://www.youtube.com/watch?v=2JNUmzuBNV0

- Go
- Authentication with JWT
- MySql 
- Docker

How to run:
- Install docker and docker compose
- Run: `docker-compose up -d`

Endpoints:

- POST http://localhost:3000/api/v1/users/login
- POST http://localhost:3000/api/v1/users/register
- POST http://localhost:3000/api/v1/tasks
- GET http://localhost:3000/api/v1/tasks/{id}
- POST http://localhost:3000/api/v1/projects
- GET http://localhost:3000/api/v1/projects/{id}

TODO: 
- Stop using network_mode as host
- Unit tests for users and projects services
