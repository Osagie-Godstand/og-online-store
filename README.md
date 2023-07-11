# online-shop-apiv1 
Online Store Backend using Fiber framework and MongoDB

## This version does not include the category, cart and order models.

## Project outline version1
- Models: user and product
- Controllers: user_handler, product_handler
- Authentication and Authorisation -> JWT tokens
- Scripts -> database management -> seeding, migration
- Testing for user and auth handlers

# Project environment variables
...
- HTTP_LISTEN_ADDRESS=:8080
- JWT_SECRET=
- MONGO_DB_NAME=online-store-apiv1
- MONGO_DB_URL=mongodb://localhost:27017
- MONGO_DB_URL_TEST=mongodb://localhost:27017
...

## Automating Go Application with Makefile
- use command: make build 
- 'make build' to build target
- 'make run' to run target
- 'make seed' to seed target
- 'make docker' to build target api with docker
- 'make test' to test target 


## Resources
### Mongodb driver
Documantation
...

https://mongodb.com/docs/drivers/go/current/quick-start
...

Installing mongodb client

go get go.mongodb.org/mongo-driver/mongo
...

### gofiber
Documentation
...

https://gofiber.io
...

Installing gofiber
...

go get github.com/gofiber/fiber/v2
...

## Docker 
### Installing mongodb as a Docker container
...
docker run --name mongodb -d mongo:latest -p 27017:27017
...

## godotenv (.env) file
...
go get github.com/joho/godotenv
...