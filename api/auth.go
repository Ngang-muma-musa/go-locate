package api

import (
	"go-locate/pkg"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	RegisterReq struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	RegisterRes struct {
		ID uint `json:"id"`
	}
)

func addAuthRoutes(c *echo.Group) {
	c.POST("/auth/register", register)
	// c.POST("/auth/login", login)
}

func register(c echo.Context) error {
	var req RegisterReq
	var err error
	if err = c.Bind(&req); err != nil {
		return err
	}

	if err = c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, e := pkg.CreateUser(req.Username, req.Password, req.Email)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusCreated, &RegisterRes{ID: user.ID})
}
