package api

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func inti() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	e = echo.New()
	fmt.Println("Hello, world!")
}
