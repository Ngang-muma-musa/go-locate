package api

import (
	"go-locate/model"
	"go-locate/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	BusinessReq struct {
		Name        string                   `json:"name" validate:"required"`
		Email       string                   `gorm:"unique" json:"email" validate:"required,email"`
		PhoneNumber []model.Contact          `json:"phoneNumber"`
		Category    []model.BusinessCategory `json:"category"`
		Description string                   `json:"description" validate:"required"`
		Location    []model.BusinessLocation `json:"location" validate:"required"`
	}

	BusinessRes struct {
		Business *model.Business `json:"business"`
	}
)

type Business struct {
	businessService *services.Business
}

func NewBusiness(businessService *services.Business) *Business {
	return &Business{businessService: businessService}
}

func (b *Business) ProvideRoutes(c *echo.Group) {
	c.POST("/business", b.Create)
	c.GET("/business", b.Find)
}

func (b *Business) Group() string {
	return "/business"
}

// Create
// @Summary      Create Business
// @Description  Creates a new business
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        business body BusinessReq false "req"
// @Success      200  {object}  BusinessRes
// @Router       /business [post]
func (b *Business) Create(c echo.Context) error {
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
	business, e := b.businessService.Create(req.Name, req.Description, req.Email, req.Location, user, req.PhoneNumber, req.Category)

	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}

	return echo.NewHTTPError(http.StatusCreated, BusinessRes{Business: business})
}

// Find
// @Summary      Find business
// @Description  Finds a business using category and location
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        location   path  string  true  "Business location"
// @Success      200  {object}  BusinessesRes
// @Router       /business [get]
func (b *Business) Find(c echo.Context) error {
	location := c.QueryParam("location")
	category, err := strconv.Atoi(c.QueryParam("category"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	business, err := b.businessService.Find(location, category)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return echo.NewHTTPError(http.StatusAccepted, business)
}
