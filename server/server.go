package server

import (
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4/middleware"

	uuid "github.com/satori/go.uuid"

	"github.com/labstack/echo/v4"
)

type (
	User struct {
		Login    string `json:"login"`
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Score    uint   `json:"score"`
	}

	RequestUser struct {
		Login    string `json:"login"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginBody struct {
		Login   string `json:"login"`
		Surname string `json:"surname"`
		Name    string `json:"name"`
		Email   string `json:"email"`
		Score   int    `json:"score"`
	}

	GoodLoginResponse struct {
		Status uint      `json:"status"`
		Data   LoginBody `json:"data"`
		Msg    string    `json:"msg"`
	}

	LogoutResponse struct {
		Status     uint   `json:"status"`
		GoodbuyMsg string `json:"goodbuy"`
	}

	SignUpBody struct {
		Login   string `json:"login"`
		Surname string `json:"surname"`
		Name    string `json:"name"`
		Email   string `json:"email"`
		Score   int    `json:"score"`
	}

	RequestSignup struct {
		Login    string `json:"login"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Surname  string `json:"surname"`
	}

	GoodSignupResponse struct {
		Status uint       `json:"status"`
		SBody  SignUpBody `json:"data"`
		Msg    string     `json:"msg"`
	}

	//Представление записи
	NewsRecord struct {
		Id           string   `json:"id"`
		PreviewUrl   string   `json:"previewUrl"`
		Tags         []string `json:"tags"`
		Title        string   `json:"title"`
		Text         string   `json:"text"`
		AuthorUrl    string   `json:"authorUrl"`
		AuthorName   string   `json:"authorName"`
		AuthorAvatar string   `json:"authorAvatar"`
		CommentsUrl  string   `json:"commentsUrl"`
		Comments     uint     `json:"comments"`
		Likes        uint     `json:"likes"`
	}

	//Тело ответа на API-call /getfeed

	RequestChunk struct {
		idLastLoaded string
		login        string
	}

	ChunkResponse struct {
		Status    uint         `json:"status"`
		ChunkData []NewsRecord `json:"schema"`
	}

	ErrorBody struct {
		Status   uint   `json:"status"`
		ErrorMsg string `json:"msg"`
	}

	MyHandler struct {
		sessions map[string]string
		sMu      sync.RWMutex
		users    map[string]User
		uMu      sync.RWMutex
	}
)

var feedSize int = 5

func NewMyHandler() MyHandler {
	return MyHandler{
		sessions: make(map[string]string, 10),
		users: map[string]User{
			"mollenTEST1":     {"mollenTEST1", "mollenTEST1", "mollenTEST1", "mollenTEST1", "mollenTEST1", 123456},
			"dar@exp.ru":      {"dar@exp.ru", "dar@exp.ru", "dar@exp.ru", "dar@exp.ru", "123", 13553},
			"viphania@exp.ru": {"viphania@exp.ru", "viphania@exp.ru", "viphania@exp.ru", "viphania@exp.ru", "123", 120},
		},
	}
}

func (api *MyHandler) Login(c echo.Context) error {
	// проверяем активные сессии
	cooke, _ := c.Cookie("session")
	api.sMu.RLock()
	login, ok := api.sessions[cooke.Value]
	api.sMu.RUnlock()
	if ok {
		api.uMu.RLock()
		user, _ := api.users[login]
		api.uMu.RUnlock()

		b := LoginBody{
		Login:   user.Login,
		Name:    user.Email,
		Surname: user.Email,
		Email:   user.Email,
		Score:   12345678, //rand.Int(),
	}
	response := GoodLoginResponse{
		Status: http.StatusOK,
		Data:   b,
		Msg:    "OK",
	}

	return c.JSON(http.StatusOK, response)
	}
	// достаем данные из запроса
	requestUser := new(RequestUser)
	if err := c.Bind(requestUser); err != nil {
		errorJson := ErrorBody{
			Status:   http.StatusInternalServerError,
			ErrorMsg: "Internal server error",
		}
		c.Logger().Printf("Error: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, errorJson)
	}
	c.Logger().Printf("login")
	// тут что-то про передачу bind полей в функции и небезопасность таких операций ¯\_(ツ)_/¯

	// логика логина
	api.uMu.RLock()
	user, ok := api.users[requestUser.Login]
	api.uMu.RUnlock()

	if !ok {
		errorJson := ErrorBody{
			Status:   http.StatusInternalServerError,
			ErrorMsg: "User doesnt exist",
		}
		return c.JSON(http.StatusInternalServerError, errorJson)
	}
	if user.Password != requestUser.Password {
		errorJson := ErrorBody{
			Status:   http.StatusInternalServerError,
			ErrorMsg: "Wrong password",
		}
		return c.JSON(http.StatusInternalServerError, errorJson)
	}
	// ставим куку на сутки
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = uuid.NewV4().String()
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(10 * time.Hour)
	c.SetCookie(cookie)
	//
	// добавляем пользователя в активные сессии
	api.sMu.Lock()
	api.sessions[cookie.Value] = user.Login
	api.sMu.Unlock()

	// формируем ответ
	b := LoginBody{
		Login:   user.Login,
		Name:    user.Email,
		Surname: user.Email,
		Email:   user.Email,
		Score:   12345678, //rand.Int(),
	}
	response := GoodLoginResponse{
		Status: http.StatusOK,
		Data:   b,
		Msg:    "OK",
	}

	return c.JSON(http.StatusOK, response)
}

func (api *MyHandler) Register(c echo.Context) error {
	// достаем данные из запроса
	newUser := new(RequestSignup)
	if err := c.Bind(newUser); err != nil {
		errorJson := ErrorBody{
			Status:   http.StatusInternalServerError,
			ErrorMsg: "Internal server error",
		}
		c.Logger().Printf("Error: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, errorJson)
	}
	api.uMu.RLock()
	_, exists := api.users[newUser.Email]
	api.uMu.RUnlock()

	if exists {
		errorJson := ErrorBody{
			Status:   http.StatusInternalServerError,
			ErrorMsg: "User already exists",
		}
		return c.JSON(http.StatusInternalServerError, errorJson)
	}

	cc, err := c.Cookie("session")
	if err == nil {
		api.sMu.RLock()
		_, exists = api.sessions[cc.Value]
		api.sMu.RUnlock()

		if exists {
			errorJson := ErrorBody{
				Status:   http.StatusInternalServerError,
				ErrorMsg: "Already authorised",
			}
			return c.JSON(http.StatusInternalServerError, errorJson)
		}
	}
	// логика регистрации,  добавляем юзера в мапу
	user := User{newUser.Login, newUser.Name, newUser.Surname, newUser.Email, newUser.Password, 12345}
	api.uMu.Lock()
	api.users[newUser.Login] = user
	api.uMu.Unlock()

	// ставим куку на сутки
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = uuid.NewV4().String()
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	//
	// добавляем пользователя в активные сессии
	api.sMu.Lock()
	api.sessions[cookie.Value] = user.Login
	api.sMu.Unlock()

	// формируем ответ
	s := SignUpBody{
		Login:   user.Login,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Score:   12345678, //rand.Int(),
	}
	response := GoodSignupResponse{
		Status: http.StatusOK,
		SBody:  s,
		Msg:    "OK",
	}

	return c.JSON(http.StatusOK, response)
}

func (api *MyHandler) Logout(c echo.Context) error {
	// удаляем пользователя из активных сессий
	cookie, err := c.Cookie("session")
	if err != nil {
		return err
	}
	api.sMu.Lock()
	delete(api.sessions, cookie.Value)
	api.sMu.Unlock()

	// ставим протухшую куку
	cookie.Expires = time.Now().Local().Add(-1 * time.Hour)
	c.SetCookie(cookie)
	// формируем ответ
	response := LogoutResponse{
		Status:     http.StatusOK,
		GoodbuyMsg: "Goodbuy, friend!",
	}
	return c.JSON(http.StatusOK, response)
}

func (api *MyHandler) Root(c echo.Context) error {
	b := LoginBody{
		Login:   "user.Email",
		Name:    "user.Email",
		Surname: "user.Email",
		Email:   "user.Email",
		Score:   rand.Int(),
	}
	u := GoodLoginResponse{
		Status: 54,
		Data:   b,
		Msg:    "OK",
	}
	return c.JSON(http.StatusOK, u)
}

func (api *MyHandler) Getfeed(c echo.Context) error {
	// достаем данные из запроса
	// requestChunk := new(RequestChunk)
	// if err := c.Bind(requestChunk); err != nil {
	// 	errorJson := ErrorBody{
	// 		Status:   http.StatusNotFound,
	// 		ErrorMsg: "Wrong request",
	// 	}
	// 	c.Logger().Printf("Error: %s", err.Error())
	// 	return c.JSON(http.StatusNotFound, errorJson)
	// }
	rec := c.QueryParam("idLastLoaded")
	from, err := strconv.Atoi(rec)
	if err != nil {
		errorJson := ErrorBody{
			Status:   http.StatusNotFound,
			ErrorMsg: err.Error() + "%%" + rec,
		}
		c.Logger().Printf("Error: %s", err.Error())
		return c.JSON(http.StatusNotFound, errorJson)
	}
	to := from + 5
	var ChunkData []NewsRecord
	//Возвращаем записи
	if from >= 0 && to < len(testData) {
		api.sMu.RLock()
		ChunkData = testData[from:to]
		api.sMu.RUnlock()
	} else {
		api.sMu.RLock()
		start := 0
		if len(testData) > 6 {
			start = len(testData) - 6
		}
		ChunkData = testData[start : len(testData)-1]
		api.sMu.RUnlock()
	}
	// формируем ответ
	response := ChunkResponse{
		http.StatusOK,
		ChunkData,
	}
	return c.JSON(http.StatusOK, response)
}

var testData = [...]NewsRecord{
	{"1", "static/img/computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1001,
	},
	{"2", "static/img/computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1002,
	},
	{"3", "static/img/computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1003,
	},
	{"4", "static/img/computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1004,
	},
	{"5", "static/img/computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1005,
	},
	{"end", "static/img/endOfFeed.png", []string{"IT-News", "Study"},
		"А всё, а раньше надо было", "", "#", "Tester-ender",
		"static/img/loader-1-HorizontalBalls.gif", "#", 0, 0,
	},
}

func Run() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	api := NewMyHandler()

	e.POST("/login", api.Login)
	e.POST("/signup", api.Register)
	e.POST("api/v1/user/logout", api.Logout)
	e.GET("/feed", api.Getfeed)
	e.GET("/", api.Root)

	e.Logger.Fatal(e.Start("192.168.0.104:8081"))
}
