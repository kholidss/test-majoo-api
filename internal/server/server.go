package server

import (
	"net/http"

	"firebase.google.com/go/auth"

	"test-majoo-api/config"
	"test-majoo-api/internal/factories"
	app_middleware "test-majoo-api/internal/middleware"
	"test-majoo-api/internal/utils"
	"test-majoo-api/internal/utils/token"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ApiServer struct {
	Router *echo.Echo
	DB     *gorm.DB
}

const defaultPort = "8080"

func NewServer(cfg *config.Config, db *gorm.DB, firebaseClient *auth.Client) *ApiServer {
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenMaker, err := token.NewJWTMaker(cfg.JwtSecret)
			if err != nil {
				utils.InternalServerError("Error when create maker")
			}
			c.Set("token-maker", tokenMaker)
			return next(c)
		}
	})

	return &ApiServer{
		Router: e,
		DB:     db,
	}
}

func (server *ApiServer) registerMiddleware() {
	server.Router.Use(middleware.Logger())
	server.Router.Use(middleware.Recover())
}

func (server *ApiServer) registerRoute() {
	loginHandler := factories.CreateLoginHandler(server.DB)
	reportHandler := factories.CreateReportHandler(server.DB)

	server.Router.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "i am alive")
	})

	//Login Route
	server.Router.POST("/login", loginHandler.UserLogin)

	//Report Route
	server.Router.GET("/report", reportHandler.MerchantReport, app_middleware.JwtAuth())
}

func (server *ApiServer) ListenAndServe(port string) {
	server.registerMiddleware()
	server.registerRoute()
	if port == "" {
		port = defaultPort
	}

	server.Router.Logger.Fatal(server.Router.Start(":" + port))
}
