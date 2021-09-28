package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var userJSON = `{"login":"mollenTEST1","email":"mollenTEST1", "password":"123"}`
var userJSON2 = "{\"status\":200,\"data\":{\"login\":\"mollenTEST1\",\"surname\":\"mollenTEST1\",\"name\":\"mollenTEST1\",\"email\":\"mollenTEST1\",\"score\":12345678},\"msg\":\"OK\"}\n"

func TestLogin(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")

	h := NewMyHandler()

	// Assertions
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON2, rec.Body.String())
	}
}
