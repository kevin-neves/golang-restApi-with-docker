# Golang REST Api with Postgres and Docker
This is a simple Go Api that exposes only one endpoint `/profiles`. 

My main goal with this project was to learn how to integrate a Gin project with Postgres database and run them togheter using Docker Compose. I've used Gorm package since it's a great ORM library write in Go, that implements Pgx driver for postgres.

### Packages used in this project
* [Gin](https://github.com/gin-gonic/gin)
* [Gorm](https://gorm.io/)

### Project folder structure
Our project entry point is in `cmd/main.go`.

```
.
└── project/
    ├── cmd/
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

### Prerequisites

- [X] make
- [x] [docker](https://docs.docker.com/engine/install/)
- [x] [docker-compose](https://docs.docker.com/compose/install/)

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
4a8bbd2b9e99   postgres:14.5-alpine3.16   "docker-entrypoint.s…"   56 minutes ago   Up 7 seconds   0.0.0.0:5432->5432/tcp, :::5432->5432/tcp                                                                                                                 
```

