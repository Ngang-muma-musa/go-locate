package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go-locate/api"
	"go-locate/model"
	"go-locate/repository"
	"go-locate/services"
	"os"
	"os/signal"
	"strconv"
	"time"

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
	db, err := model.InitDB(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatalf("Error instantiation database %v\n", err)
	}

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("invalid port: %v", err)
	}

	customValidator := &api.CustomValidator{Validator: validator.New()}

	userRepo := repository.NewUser(db)
	userService := services.NewUser(userRepo)
	server := api.NewServer(e, port, customValidator, userService)

	businessRepo := repository.NewBusiness(db)
	businessService := services.NewBusiness(businessRepo)

	categoryRepo := repository.NewCategory(db)
	categoryService := services.NewCategory(categoryRepo)

	adminService := services.NewAdmin(businessRepo)
	adminEndpoint := api.NewAdmin(businessService, categoryService, adminService)
	server.RegisterEndpoints(adminEndpoint)

	authEndpoint := api.NewAuth(userService)
	server.RegisterEndpoints(authEndpoint)

	businessEndpoint := api.NewBusiness(businessService)
	server.RegisterEndpoints(businessEndpoint)

	server.Start()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
