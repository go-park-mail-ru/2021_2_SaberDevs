package server

import (
	"fmt"
	"math/rand"
	"net/http"
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
		Title        string   `json:"title"`
		Text         string   `json:"text"`
		AuthorUrl    string   `json:"authorUrl"`
		AuthorName   string   `json:"authorName"`
		AuthorAvatar string   `json:"authorAvatar"`
		CommentsUrl  string   `json:"commentsUrl"`
		Comments     uint     `json:"comments"`
		Likes        uint     `json:"likes"`
		Tags         []string `json:"tags"`
	}

	//Тело ответа на API-call /getfeed
	Chunk struct {
		From      string       `json:"from"`
		To        string       `json:"to"`
		ChunkData []NewsRecord `json:"chunk"`
	}

	RequestChunk struct {
		From int `json:"from"`
		To   int `json:"to"`
	}

	ChunkResponse struct {
		Status          uint  `json:"status"`
		NewsRecordChunk Chunk `json:"body"`
	}

	ErrorBody struct {
		Status   uint   `json:"status"`
		ErrorMsg string `json:"error"`
	}

	MyHandler struct {
		sessions map[string]string
		sMu      sync.RWMutex
		users    map[string]User
		uMu      sync.RWMutex
	}
)

var endOfFeed = NewsRecord{"endOfFeedMarkerID", "static/img/endOfFeed.png",
	"А всё, а раньше надо было", "", "#", "Tester-ender",
	"static/img/loader-1-HorizontalBalls.gif", "#", 0, 0, []string{"Bottom"},
}

func NewMyHandler() MyHandler {
	return MyHandler{
		sessions: make(map[string]string, 10),
		users: map[string]User{
			"mollenTEST1":     {"mollenTEST1", "mollenTEST1", "mollenTEST1", "mollenTEST1", "123", 123456},
			"dar@exp.ru":      {"dar@exp.ru", "dar@exp.ru", "dar@exp.ru", "dar@exp.ru", "123", 13553},
			"viphania@exp.ru": {"viphania@exp.ru", "viphania@exp.ru", "viphania@exp.ru", "viphania@exp.ru", "123", 120},
		},
	}
}

func (api *MyHandler) Login(c echo.Context) error {
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
	cookie.Expires = time.Now().Add(10 * time.Second)
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
	api.users[newUser.Email] = user
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
	requestChunk := new(RequestChunk)
	if err := c.Bind(requestChunk); err != nil {
		errorJson := ErrorBody{
			Status:   http.StatusNotFound,
			ErrorMsg: "Wrong request",
		}
		c.Logger().Printf("Error: %s", err.Error())
		return c.JSON(http.StatusNotFound, errorJson)
	}

	f := requestChunk.From
	t := requestChunk.To

	//Возвращаем записи
	chunk := Chunk{}
	chunk.From = fmt.Sprint(requestChunk.From)
	chunk.To = fmt.Sprint(requestChunk.To)
	if f >= 0 && t >= 0 && t >= f && t < len(testData) {
		api.sMu.RLock()
		chunk.ChunkData = testData[f:t]
		api.sMu.RUnlock()
	} else {
		api.sMu.RLock()
		chunk.ChunkData = []NewsRecord{endOfFeed}
		api.sMu.RUnlock()
	}
	// формируем ответ
	response := ChunkResponse{
		http.StatusOK,
		chunk,
	}
	return c.JSON(http.StatusOK, response)
}

var testData = [...]NewsRecord{
	{"1", "static/img/computer.png", "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1001, []string{"IT-News", "Study"},
	},
	{"2", "static/img/computer.png", "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1002, []string{"IT-News", "Study"},
	},
	{"3", "static/img/computer.png", "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1003, []string{"IT-News", "Study"},
	},
	{"4", "static/img/computer.png", "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1004, []string{"IT-News", "Study"},
	},
	{"5", "static/img/computer.png", "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1005, []string{"IT-News", "Study"},
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
	e.POST("api/v1/user/getfeed", api.Getfeed)
	e.GET("/", api.Root)

	e.Logger.Fatal(e.Start("192.168.0.104:8081"))
}
