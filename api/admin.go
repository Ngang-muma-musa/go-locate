package api

import (
	"go-locate/model"
	"go-locate/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	CategoryReq struct {
		Name string `json:"name" validate:"required"`
	}

	VerifyReq struct {
		ID int `json:"id" validate:"required"`
	}

	VerifyRes struct {
		Business *model.Business
	}
)

type Admin struct {
	businessService *services.Business
	adminService    *services.Admin
	categoryService *services.Category
}

func NewAdmin(businessService *services.Business, categoryService *services.Category, adminService *services.Admin) *Admin {
	return &Admin{businessService: businessService, categoryService: categoryService, adminService: adminService}
}

func (a *Admin) ProvideRoutes(c *echo.Group) {
	c.POST("/admin/category", a.Create)
	c.POST("/admin/verify-business/:id", a.VerifyBusiness)
}

func (a *Admin) Group() string {
	return "/admin"
}

func (a *Admin) Create(c echo.Context) error {
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

	category, err := a.categoryService.Create(req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return echo.NewHTTPError(http.StatusCreated, category)
}

func (a *Admin) VerifyBusiness(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = a.adminService.VerifyBusiness(uint(ID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	business, err := a.businessService.GetByID(uint(ID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return echo.NewHTTPError(http.StatusAccepted, VerifyRes{Business: business})
}
