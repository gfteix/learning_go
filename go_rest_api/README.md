Learning how to build an rest api with go https://www.youtube.com/watch?v=2JNUmzuBNV0

- Go
- Authentication with JWT
- MySql 
- Docker

Docker command to start db

`sudo docker run --name mysql_docker -p 33060:3306 -e MYSQL_ROOT_PASSWORD=mysql123 -e MYSQL_DATABASE=projectmanager -d mysql`

TODO: 
- Add Dockerfile and Docker compose
- Unit tests for users and projects services
