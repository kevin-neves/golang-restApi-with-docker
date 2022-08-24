# Golang REST Api with Postgres and Docker
This is a simple Go Api that exposes only one endpoint `/profiles`. 

My main goal with this project was to learn how to integrate a Gin project with Postgres database and run them together using Docker Compose. I've used Gorm package since it's a great ORM library write in Go, that implements [Pgx](https://github.com/jackc/pgx) driver for postgres.

### Packages used in this project
* [Gin](https://github.com/gin-gonic/gin)
* [Gorm](https://gorm.io/)

### Project folder structure
The project entry point is in **cmd/main.go**.

```
.
└── project/
    ├── cmd
    │   └── main.go
    ├── out/
    │   └── "final app"
    ├── pkg/
    │   ├── common/
    │   │   ├── db/
    │   │   │   └── db.go
    │   │   ├── env/
    │   │   │   └── env.go
    │   │   └── models/
    │   │       └── profile.go
    │   └── profile/
    │       ├── add_profile.go
    │       ├── controller.go
    │       ├── delete_profile.go
    │       ├── get_profile.go
    │       ├── get_profiles.go
    │       └── update_profile.go
    ├── .env
    ├── docker-compose.yml
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── Makefile
    └── README.md
```

### The Dockerfile

Following best practices, I've created the server's Dockerfile to be ran in two steps:

1. Create an image, copying the project's files and building the server application
2. Create a fresh Alpine Linux image and copying the builded project into it.

```
## Dockerfile
FROM golang:1.19-alpine AS build_base

RUN apk add --no-cache git
RUN apk add --update make

WORKDIR /tmp/golang-restApi-with-docker

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make build

# Create a new fresh image, and optmize the size of the application
FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build_base /tmp/golang-restApi-with-docker/out/golang-restApi-with-docker /app/golang-restApi-with-docker

EXPOSE 8080

CMD ["./app/golang-restApi-with-docker"]
```

That way, the final container created by the image has approximately **27 MB**.

### Prerequisites

- [X] make
- [x] [docker](https://docs.docker.com/engine/install/)
- [x] [docker-compose](https://docs.docker.com/compose/install/)

You also need to create a .**env** file at the root directory with some configurations to be loaded when the Docker Compose run, specially **user** and **password** to create the Database. See the following example:
```
## .env
# Server configuration
PORT=":8080"
DB_URL="postgres://USER:PASSOWRD@db:5432/postgres"

# DataBase configuration
POSTGRES_USER="USER"
POSTGRES_PASSWORD="PASSOWRD"
```
The same USER and PASSWORD you use for postgres needs to inserted into the **DB_URL** variable.


### Running the project

Running using make:
```
make docker
```

Running using docker-compose:
```
docker-compose up -d --build
```

In both cases, you should see something like this when running `docker ps` command in terminal:
```
CONTAINER ID   IMAGE                      COMMAND                  CREATED          STATUS         PORTS                                       NAMES
363b3388e55e   go-databases-server        "./app/golang-restAp…"   56 minutes ago   Up 8 seconds   0.0.0.0:8080->8080/tcp, :::8080->8080/tcp   golang-server
4a8bbd2b9e99   postgres:14.5-alpine3.16   "docker-entrypoint.s…"   56 minutes ago   Up 7 seconds   0.0.0.0:5432->5432/tcp, :::5432->5432/tcp   go-databases-db-1
```

### Endpoints

| Method | Endpoint | Description |
| ----------- | ----------- | ----------- |
| GET | "/profiles/" | Get all the profiles |
| GET | "/profiles/:id" | Get profile by ID |
| POST | "/profiles/" | Create new profile |
| PUT | "/profiles/:id" | Update and existing profile |
| DELETE | "/profiles/:id" | Delete the profile using ID |