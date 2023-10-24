package api

import (
	"go-locate/model"
	"go-locate/pkg"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	CategoryReq struct {
		Category string `json:"category" validate:"required"`
	}

	CategoryRes struct {
		Category *model.Category `json:"category"`
	}
)

func addCategoryRoutes(c *echo.Group) {
	c.POST("/category", createBusiness)
}

func createCategory(c echo.Context) error {
	var req CategoryReq
	var err error
	if err = c.Bind(&req); err != nil {
		return err
	}
	user := GetUserFromContext(c)

	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, ErrUnauthorized.Error())
	}

	if user.IsAdmin == false {
		return echo.NewHTTPError(http.StatusUnauthorized, ErrUnauthorized.Error())
	}

	if err = c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	category, err := pkg.CreateCategory(req.Category)

	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return echo.NewHTTPError(http.StatusCreated, CategoryRes{Category: category})
}
