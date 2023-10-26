package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"go-locate/services"
	"log"
	"net/http"
	"os"
	"strings"
)

type Handler interface {
	ProvideRoutes(c *echo.Group)
	Group() string
}

type Server struct {
	e           *echo.Echo
	port        int
	validator   *CustomValidator
	userService *services.User
}

func NewServer(e *echo.Echo, port int, validator *CustomValidator, userService *services.User) *Server {
	e.Validator = validator
	signingKey := os.Getenv("JWT_SECRET")
	jwtConfig := echojwt.Config{
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
		SuccessHandler: func(c echo.Context) {
			claims := ClaimsFromContext(c)
			user := userService.GetByID(claims.UserID)
			if user == nil {
				c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "invalid token"})
				return
			}
			c.Set("user", user)
		},
	}
	e.Use(echojwt.WithConfig(jwtConfig))
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return &Server{
		e:           e,
		port:        port,
		validator:   validator,
		userService: userService,
	}
}

func (s *Server) Use(middlewareFunc ...echo.MiddlewareFunc) {
	s.e.Use(middlewareFunc...)
}

func (s *Server) RegisterEndpoints(endpoints ...Handler) {
	for _, endpoint := range endpoints {
		grp := s.e.Group(endpoint.Group())
		endpoint.ProvideRoutes(grp)
	}
}

func (s *Server) Start() {
	go func() {
		if err := s.e.Start(fmt.Sprintf(":%d", s.port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("shutting down the server")
		}
	}()
}

func (s *Server) Shutdown(ctx context.Context) {
	if err := s.e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
