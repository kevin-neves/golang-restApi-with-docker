package env

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEnvVar(path string) {
	if err := godotenv.Load(path); err != nil {
		log.Println(err)
		return
	}
}
