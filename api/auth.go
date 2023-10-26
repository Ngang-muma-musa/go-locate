package api

import (
	"errors"
	"go-locate/services"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	JWTDefaultTTL = 72 * time.Hour
)

var (
	// ErrUnauthorized is when ther user is not authorized.
	ErrUnauthorized = errors.New("unauthorized")
)

type (
	RegisterReq struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	LoginReq struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	RegisterRes struct {
		ID uint `json:"id"`
	}

	LoginRes struct {
		Token string `json:"token"`
	}
)

type Auth struct {
	userService *services.User
}

func NewAuth(userService *services.User) *Auth {
	return &Auth{userService: userService}
}

func (a *Auth) ProvideRoutes(c *echo.Group) {
	c.POST("/auth/register", a.Register)
	c.POST("/auth/login", a.Login)
}

func (a *Auth) Group() string {
	return "/auth"
}

// Register
// @Summary      Register
// @Description  registers a new user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user body RegisterReq true "req"
// @Success      200  {object}  RegisterRes
// @Router       /auth/register [post]
func (a *Auth) Register(c echo.Context) error {
	var req RegisterReq
	var err error
	if err = c.Bind(&req); err != nil {
		return err
	}

	if err = c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, e := a.userService.Create(req.Username, req.Password, req.Email)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusCreated, &RegisterRes{ID: user.ID})
}

// Login
// @Summary      Login
// @Description  Login a user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user body LoginReq true "req"
// @Success      200  {object}  LoginRes
// @Router       /auth/login [post]
func (a *Auth) Login(c echo.Context) error {
	var req LoginReq
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := a.userService.GetByEmail(req.Email)
	if user == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "one or more details is incorrect")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "one or more details is incorrect")
	}

	signedToken, e := GetToken(user.ID)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}
	return c.JSON(http.StatusOK, &LoginRes{Token: signedToken})
}

func GetToken(userID uint) (string, error) {
	ttl, err := time.ParseDuration(os.Getenv("JWT_TTL"))
	if err != nil {
		ttl = JWTDefaultTTL
	}
	claims := &JwtCustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		},
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
