# tdd-online-store
Online Store Backend using Fiber framework and MongoDB

## Project outline

- users -> buy products from an online store
- Authentication and authorisation -> JWT tokens
- Store
- Products
- Scripts -> database management

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
Docker run --name mongodb -d mongo:latest -p 27017:27017
...