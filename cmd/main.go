package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kevin-neves/golang-restApi-with-docker/pkg/common/db"
	"github.com/kevin-neves/golang-restApi-with-docker/pkg/common/env"
	"github.com/kevin-neves/golang-restApi-with-docker/pkg/profile"
)

func main() {
	if os.Getenv("MODE") != "production" {
		env.GetEnvVar("../.env")
	}

	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	r := gin.Default()
	h := db.Init(dbUrl)

	profile.RegisterRoutes(r, h)

	r.Run(port)
}
