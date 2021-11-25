package server

import (
	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	ahandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/handler"
	arepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/repository"
	chandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/handler"
	crepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/repository"
	cusecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/usecase"
	ihandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/image/handler"
	irepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/image/repository"
	iusecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/image/usecase"
	krepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/keys/repository"
	syberMiddleware "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/middleware"
	shandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/handler"
	srepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/repository"
	susecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/usecase"
	uhandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/handler"
	urepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/repository"
	uusecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/usecase"
	commentWS "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/ws/commentStream"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tarantool/go-tarantool"
	"google.golang.org/grpc"
	"net/http"
)

func TarantoolConnect() (*tarantool.Connection, error) {
	user, pass, addr, err := TarantoolConfig()
	if err != nil {
		return nil, err
	}

	opts := tarantool.Opts{User: user, Pass: pass}
	conn, err := tarantool.Connect(addr, opts)
	if err != nil {
		return nil, err
	}

	_, err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func DbConnect() (*sqlx.DB, error) {
	connStr, err := DbConfig()
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return db, err
	}
	err = db.Ping()
	if err != nil {
		return db, err
	}
	return db, err
}

func DbClose(db *sqlx.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return err
}

func router(e *echo.Echo, db *sqlx.DB, sessionsDbConn *tarantool.Connection, a *app.ArticleDeliveryClient) {
	//e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	//	TokenLookup: "header:X-XSRF-TOKEN",
	//}))

	publisher := commentWS.NewPublisher()
	go publisher.Run()
	commentWSAPI := commentWS.NewCommentStreamHandler(publisher)

	userRepo := urepo.NewUserRepository(db)
	sessionRepo := srepo.NewSessionRepository(sessionsDbConn)
	keyRepo := krepo.NewKeyRepository(sessionsDbConn)
	articleRepo := arepo.NewArticleRepository(db)
	imageRepo := irepo.NewImageRepository()
	commentsRepo := crepo.NewCommentRepository(db)

	streamCommetChecker := commentWS.NewRepoChecker(publisher, commentsRepo)
	go streamCommetChecker.Run()

	userUsecase := uusecase.NewUserUsecase(userRepo, sessionRepo, keyRepo, articleRepo)
	userAPI := uhandler.NewUserHandler(userUsecase)

	sessionUsecase := susecase.NewsessionUsecase(userRepo, sessionRepo)
	sessionAPI := shandler.NewSessionHandler(sessionUsecase)

	//articlesUsecase := ausecase.NewArticleUsecase(articleRepo, sessionRepo)
	articlesAPI := ahandler.NewArticlesHandler(*a)

	imageUsecase := iusecase.NewImageUsecase(imageRepo)
	imageAPI := ihandler.NewImageHandler(imageUsecase)

	commentUsecase := cusecase.NewCommentUsecase(userRepo, sessionRepo, commentsRepo)
	commentsAPi := chandler.NewCommentHandler(commentUsecase)

	articles := e.Group("/api/v1/articles")
	articles.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{}))
	search := e.Group("/api/v1/search")
	search.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{}))
	authMiddleware := syberMiddleware.NewAuthMiddleware(sessionRepo)

	e.Use(syberMiddleware.ValidateRequestBody)

	//Logger.SetOutput() //to file
	e.Logger.SetLevel(log.INFO)
	// e.Logger.SetLevel(log.ERROR)

	e.HTTPErrorHandler = syberMiddleware.ErrorHandler
	e.Use(syberMiddleware.AccessLogger)
	e.Use(syberMiddleware.AddId)

	e.GET("api/v1/ws", commentWSAPI.HandleWS)

	e.GET("api/v1/img/:name", imageAPI.GetImage)
	e.POST("api/v1/img/upload", imageAPI.SaveImage, authMiddleware.CheckAuth)

	e.POST("api/v1/comments/create", commentsAPi.CreateComment, authMiddleware.CheckAuth)
	e.POST("api/v1/comments/update", commentsAPi.UpdateComment, authMiddleware.CheckAuth)
	e.GET("api/v1/comments", commentsAPi.GetCommentsByArticleID)

	e.POST("api/v1/user/login", userAPI.Login)
	e.POST("api/v1/user/signup", userAPI.Register)
	e.POST("api/v1/user/logout", userAPI.Logout, authMiddleware.CheckAuth)
	e.POST("api/v1/", sessionAPI.CheckSession)
	e.POST("api/v1/user/profile/update", userAPI.UpdateProfile, authMiddleware.CheckAuth)
	e.GET("api/v1/user/profile", userAPI.UserProfile, authMiddleware.CheckAuth)
	e.GET("api/v1/user", userAPI.AuthorProfile)

	articles.GET("/feed", articlesAPI.GetFeed)
	articles.GET("", articlesAPI.GetByID)
	articles.GET("/author", articlesAPI.GetByAuthor)
	articles.GET("/category", articlesAPI.GetByCategory)
	articles.GET("/tags", articlesAPI.GetByTag)
	articles.POST("/create", articlesAPI.Create, authMiddleware.CheckAuth)
	articles.POST("/update", articlesAPI.Update, authMiddleware.CheckAuth)
	articles.POST("/delete", articlesAPI.Delete, authMiddleware.CheckAuth)

	search.GET("/articles", articlesAPI.FindArticles)
	search.GET("/author", articlesAPI.FindAuthors)
	search.GET("/tags", articlesAPI.FindByTag)
}

func Run(address string) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080", "http://87.228.2.178:8080", "http://89.208.197.247:8080", "ws://localhost:8080"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
	}))

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{http.MethodGet, http.MethodPost},
	// 	AllowCredentials: true,
	// }))

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	db, err := DbConnect()
	if err != nil {
		e.Logger.Fatal(err)
	}

	tarantoolConn, err := TarantoolConnect()
	if err != nil {
		e.Logger.Fatal(err)
	}
	grcpConn, err := grpc.Dial(
		"127.0.0.1:8079",
		grpc.WithInsecure(),
	)
	if err != nil {
		e.Logger.Fatal(err)
	}

	defer grcpConn.Close()

	sessManager := app.NewArticleDeliveryClient(grcpConn)

	defer DbClose(db)

	router(e, db, tarantoolConn, &sessManager)

	e.Logger.Fatal(e.Start(address))
}
