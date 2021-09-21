package routes

import (
	"log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/myrachanto/power/controllers"
	"github.com/spf13/viper"
	// m "github.com/myrachanto/power/middlewares"
)
//StoreAPI =>entry point to routes
type Open struct {
	Port string `mapstructure:"PORT"`
	Key  string `mapstructure:"EncryptionKey"`
}

// func LoadConfig(path string) (open Open, err error) {
func LoadConfig() (open Open, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&open)
	return
}
func StoreApi() {

	open, err := LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	e := echo.New()

	e.Static("/", "public")
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	JWTgroup := e.Group("/api/")
	JWTgroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey: []byte(open.Key),
	}))
	front := e.Group("/front")
	front.POST("/register", controllers.UserController.Create)
	front.POST("/login", controllers.UserController.Login)
	JWTgroup.GET("logout/:token", controllers.UserController.Logout)
	// JWTgroup.GET("users", controllers.UserController.GetAll)
	// JWTgroup.GET("users/:id", controllers.UserController.GetOne)
	JWTgroup.PUT("users/:id", controllers.UserController.Update)
	JWTgroup.DELETE("users/:id", controllers.UserController.Delete)
	//e.DELETE("loggoutall/:id", controllers.UserController.DeleteALL) logout all accounts
	///////////workspace/////////////////////////////
	JWTgroup.POST("workspaces", controllers.WorkspaceController.Create)
	JWTgroup.GET("workspaces", controllers.WorkspaceController.GetAll)
	JWTgroup.GET("workspaces/list/:code", controllers.WorkspaceController.Getlist)
	JWTgroup.GET("workspaces/:id", controllers.WorkspaceController.GetOne)
	JWTgroup.PUT("workspaces/:id", controllers.WorkspaceController.Update)
	JWTgroup.DELETE("workspaces/:code", controllers.WorkspaceController.Delete)
	///////////project/////////////////////////////
	JWTgroup.POST("projects", controllers.ProjectController.Create)
	JWTgroup.GET("projects", controllers.ProjectController.GetAll)
	JWTgroup.POST("projects/upload", controllers.ProjectController.Upload)
	JWTgroup.GET("projects/list", controllers.ProjectController.Getlist)
	JWTgroup.GET("projects/info", controllers.ProjectController.Getinfo)
	JWTgroup.GET("projects/results", controllers.ProjectController.Results)
	JWTgroup.GET("projects/:id", controllers.ProjectController.GetOne)
	JWTgroup.PUT("projects/:id", controllers.ProjectController.Update)
	JWTgroup.DELETE("projects/:code", controllers.ProjectController.Delete)

	// Start server
	e.Logger.Fatal(e.Start(open.Port))
}