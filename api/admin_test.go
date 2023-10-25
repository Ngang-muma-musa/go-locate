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

func TestCreateCategory(t *testing.T) {
	pkg.SetupTest(t)
	defer pkg.ShutdownTest(func() {
		model.ClearTables(&model.User{})
	})

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	_, err := pkg.CreateCategory("TestCategory")

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var res CategoryReq
		err = json.Unmarshal(rec.Body.Bytes(), &res)
		require.NoError(t, err)
		assert.Equal(t, uint(1), res.Category)
	}
}
