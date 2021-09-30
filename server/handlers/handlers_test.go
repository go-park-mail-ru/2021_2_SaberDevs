package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/server/data"
	"server/server/models"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// Setup
	e := echo.New()
	loginReq := new(models.RequestUser)
	loginReq.Login = "mollenTEST1"
	loginReq.Password = "mollenTEST1"

	loginRequest, _ := json.Marshal(loginReq)
	loginData := new(models.LoginData)
	loginData.Login = "mollenTEST1"
	loginData.Surname = "mollenTEST1"
	loginData.Name = "mollenTEST1"
	loginData.Email = "mollenTEST1"
	loginData.Score = 0

	loginResponse := new(models.LoginResponse)
	loginResponse.Status = 200
	loginResponse.Data = *loginData
	loginResponse.Msg = "OK"

	loginRes, _ := json.Marshal(loginResponse)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(loginRequest)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")

	h := NewMyHandler()

	// Assertions
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(loginRes)+"\n", rec.Body.String())
	}
}

func TestSignUp(t *testing.T) {
	// Setup
	e := echo.New()
	newUser := new(models.RequestSignup)
	newUser.Login = "yura"
	newUser.Email = "yuramail.ru"
	newUser.Password = "yuramail.ru"
	newUser.Name = "yura"
	newUser.Surname = "Lyubsk"
	signrec, _ := json.Marshal(newUser)
	data := new(models.SignUpData)
	data.Login = newUser.Login
	data.Surname = newUser.Surname
	data.Name = newUser.Name
	data.Email = newUser.Email
	data.Score = 0
	res := new(models.SignupResponse)
	res.Data = *data
	res.Status = 200
	res.Msg = "OK"
	signres, _ := json.Marshal(res)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(signrec)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/signup")
	h := NewMyHandler()
	// Assertions
	if assert.NoError(t, h.Register(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(signres)+"\n", rec.Body.String())
	}
}

func TestFeed(t *testing.T) {
	// Setup
	e := echo.New()
	request := "?idLastLoaded=1&login=all"
	testData := data.TestData
	chunkData := testData[0:5]
	// формируем ответ
	response := models.ChunkResponse{
		Status:    http.StatusOK,
		ChunkData: chunkData,
	}
	res, _ := json.Marshal(response)
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(request))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/feed")
	h := NewMyHandler()
	// Assertions
	if assert.NoError(t, h.Getfeed(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res)+"\n", rec.Body.String())
	}
}

func TestLogout(t *testing.T) {
	// Setup
	e := echo.New()
	loginReq := new(models.RequestUser)
	loginReq.Login = "mollenTEST1"
	loginReq.Password = "mollenTEST1"
	loginRequest, _ := json.Marshal(loginReq)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(loginRequest)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")
	h := NewMyHandler()
	// Assertions
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	answer := rec.Result()
	cookies := answer.Cookies()
	res := new(models.LogoutResponse)
	res.Status = 200
	res.GoodbyeMsg = "Goodbye, friend!"
	response, _ := json.Marshal(res)
	rec = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.AddCookie(cookies[0])
	c = e.NewContext(req, rec)
	c.SetPath("/logout")
	// Assertions
	if assert.NoError(t, h.Logout(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(response)+"\n", rec.Body.String())
	}
}
