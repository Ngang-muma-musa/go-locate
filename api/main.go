package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

var e *echo.Echo

var (
	signingKey string
	JWTConfig  echojwt.Config
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	e = echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	signingKey = os.Getenv("JWT_SECRET")
	JWTConfig = echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: []byte(signingKey),
		Skipper: func(c echo.Context) bool {
			freeRoutes := []string{"/auth/login", "/auth/register", "/swagger"}
			for _, route := range freeRoutes {
				if strings.Contains(c.Request().RequestURI, route) {
					return true
				}
			}
			return false
		},
	}
	fmt.Println("Hello, world!")
}

func Start() {
	var wg sync.WaitGroup
	e.Use(echojwt.WithConfig(JWTConfig))
	base := e.Group("")
	addAuthRoutes(base)
	addBusinessRoutes(base)
	addAdminRoutes(base)
	e.GET("/business", findBusiness)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	wg.Add(1)
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
			wg.Done()
		}
	}()
	wg.Wait()
}
