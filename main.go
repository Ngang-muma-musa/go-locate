package main

import (
	"go-locate/api"
	"go-locate/model"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

// @title Go-Locate
// @version 1.0
// @description This applications helps find businesses by location and category
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	err = model.InitDB(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatalf("Error instantiation database %v\n", err)
	}
	api.Start()
}
