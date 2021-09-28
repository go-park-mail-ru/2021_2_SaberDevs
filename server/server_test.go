package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// RequestSignup struct {
// 	Login    string `json:"login"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// 	Name     string `json:"name"`
// 	Surname  string `json:"surname"`
// }

var loginJson = `{"login":"mollenTEST1","email":"mollenTEST1", "password":"123"}`
var answerLogin = "{\"status\":200,\"data\":{\"login\":\"mollenTEST1\",\"surname\":\"mollenTEST1\",\"name\":\"mollenTEST1\",\"email\":\"mollenTEST1\",\"score\":12345678},\"msg\":\"OK\"}\n"
var signupJson = "{\"login\":\"yura\",\"email\":\"yura@ya.ru\",\"password\":12345678\",\"name\":\"yura\",\"surname\":\"yura\"}"
var answerSignup = "{\"login\":\"yura\",\"email\":\"yura@ya.ru\",\"password\":12345678\",\"name\":\"yura\",\"surname\":\"yura\"}"

func TestLogin(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(loginJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")

	h := NewMyHandler()

	// Assertions
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, answerLogin, rec.Body.String())
	}
}

func TestSignUp(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(signupJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/signup")

	h := NewMyHandler()

	// Assertions
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, answerSignup, rec.Body.String())
	}
}
