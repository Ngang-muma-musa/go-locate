package api

import (
	"encoding/json"
	"go-locate/model"
	"go-locate/pkg"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	userJSON = `{"username":"Jon Snow","email":"jon@labstack.com","password":"12345678"}`
)

func TestRegister(t *testing.T) {
	pkg.SetupTest(t)
	defer pkg.ShutdownTest(func() {
		model.ClearTables(&model.User{})
	})

	// Setup
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := register(c)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var res RegisterRes
		err = json.Unmarshal(rec.Body.Bytes(), &res)
		require.NoError(t, err)
		assert.Equal(t, uint(1), res.ID)
	}
}

func TestLogin(t *testing.T) {
	pkg.SetupTest(t)
	defer pkg.ShutdownTest(func() {
		model.ClearTables(&model.User{})
	})

	_, err := pkg.CreateUser("muma", "muma1234", "ngangmusa0@gmail.com")
	require.NoError(t, err)

	u := &LoginReq{
		Email:    "ngangmusa0@gmail.com",
		Password: "muma1234",
	}
	d, _ := json.Marshal(u)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(d)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = login(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var res LoginRes
	err = json.Unmarshal(rec.Body.Bytes(), &res)
	require.NoError(t, err)
	assert.NotEqual(t, "", res.Token)
}
