package api

import (
	"go-locate/model"
	"go-locate/pkg"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	BusinessReq struct {
		Name        string                   `json:"name" validate:"required"`
		Email       string                   `gorm:"unique" json:"email" validate:"required,email"`
		PhoneNumber []model.Contact          `json:"phoneNumber"`
		Category    []model.BusinessCategory `json:"category"`
		Description string                   `json:"description" validate:"required"`
		Location    string                   `json:"location" validate:"required"`
	}

	BusinessRes struct {
		Business *model.Business `json:"business"`
	}
)

func addBusinessRoutes(c *echo.Group) {
	c.POST("/business", createBusiness)
}

func createBusiness(c echo.Context) error {
	var req BusinessReq
	var err error
	if err = c.Bind(&req); err != nil {
		return err
	}
	user := GetUserFromContext(c)
	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, ErrUnauthorized.Error())
	}
	if err = c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	bussines, e := pkg.CreateBusiness(req.Name, req.Description, req.Email, req.Location, user, req.PhoneNumber, req.Category)

	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}

	return echo.NewHTTPError(http.StatusBadRequest, BusinessRes{Business: bussines})
}
