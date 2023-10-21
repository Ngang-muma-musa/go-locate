package main

import (
	"go-locate/api"
	"go-locate/model"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	err = model.InitDB(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	if err != nil {
		println("Error")
		// log.WithError(err).
		// 	Fatal("failed to open db")
	}
	api.Start()
}
