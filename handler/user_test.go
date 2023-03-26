package handler

import (
	"company-keeper-go/router/middleware"
	"company-keeper-go/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSignUpCaseSuccess(t *testing.T) {
	setup()
	var (
		reqJSON = `{"username":"bob","email":"bob@gmail.com","password":"secret"}`
	)
	req := httptest.NewRequest(echo.POST, "/api/users", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, h.SignUp(c))
	if assert.Equal(t, http.StatusCreated, rec.Code) {
		m := responseMap(rec.Body.Bytes())
		assert.Equal(t, "bob", m["username"])
		assert.Equal(t, "bob@gmail.com", m["email"])
	}
}

func TestLoginCaseSuccess(t *testing.T) {
	setup()
	var (
		reqJSON = `{"email":"jake@gmail.com","password":"secret"}`
	)
	req := httptest.NewRequest(echo.POST, "/api/users/login", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, h.Login(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		m := responseMap(rec.Body.Bytes())
		assert.Equal(t, "Jake", m["username"])
		assert.Equal(t, "jake@gmail.com", m["email"])
		assert.NotEmpty(t, m["token"])
	}
}

func TestLoginCaseFailed(t *testing.T) {
	setup()
	var (
		reqJSON = `{"email":"userx@gmail.com","password":"secret"}`
	)
	req := httptest.NewRequest(echo.POST, "/api/users/login", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, h.Login(c))
	assert.Equal(t, http.StatusForbidden, rec.Code)
}

func TestCurrentUserCaseSuccess(t *testing.T) {
	setup()
	jwtMiddleware := middleware.JWT(utils.GetJWTSecret())
	req := httptest.NewRequest(echo.GET, "/api/user", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return h.CurrentUser(c)
	})(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusOK, rec.Code) {
		m := responseMap(rec.Body.Bytes())
		assert.Equal(t, "Jake", m["username"])
		assert.Equal(t, "jake@gmail.com", m["email"])
		assert.NotEmpty(t, m["token"])
	}
}

func TestCurrentUserCaseInvalid(t *testing.T) {
	setup()
	jwtMiddleware := middleware.JWT(utils.GetJWTSecret())
	req := httptest.NewRequest(echo.GET, "/api/user", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(100)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return h.CurrentUser(c)
	})(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}
