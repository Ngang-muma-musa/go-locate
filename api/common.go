package api

import (
	"go-locate/model"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type (
	// ErrorResponse is a generic error response.
	ErrorResponse struct {
		Error string `json:"error,omitempty"`
	}

	// CustomValidator is a custom validator.
	CustomValidator struct {
		validator *validator.Validate
	}

	// JwtCustomClaims are custom claims extending default ones.
	// See https://github.com/golang-jwt/jwt for more examples
	JwtCustomClaims struct {
		AuthLevel uint `json:"authLevel"`
		UserID    uint `json:"userId"`
		jwt.RegisteredClaims
	}
)

type agggregatedLogger struct {
	inFoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

// Validate validates the interface.
func (c *CustomValidator) Validate(i interface{}) error {
	return c.validator.Struct(i)
}

// ClaimsFromContext extracts the claims from the http request context.
func ClaimsFromContext(c echo.Context) *JwtCustomClaims {
	token := c.Get("user").(*jwt.Token)
	return token.Claims.(*JwtCustomClaims)
}

// GetUserFromContext returns the full user from the request context.
func GetUserFromContext(c echo.Context) *model.User {
	claims := ClaimsFromContext(c)
	return model.GetUserByID(claims.UserID)
}

func (l *agggregatedLogger) Info(v ...interface{}) {
	l.inFoLogger.Println(v...)
}

func (l *agggregatedLogger) Warn(v ...interface{}) {
	l.warnLogger.Println(v...)
}

func (l *agggregatedLogger) Error(v ...interface{}) {
	l.errorLogger.Println(v...)
}
