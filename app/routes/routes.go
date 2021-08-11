package routes

import (
	"rentRoom/controllers/category"
	"rentRoom/controllers/rooms"
	"rentRoom/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	UserController     users.UserController
	RoomsController     rooms.RoomsController
	CategoryController category.CategoryController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", cl.UserController.Store)
	users.GET("/token", cl.UserController.CreateToken)

	category := e.Group("category")
	category.GET("/list", cl.CategoryController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))

	news := e.Group("rooms")
	news.POST("/store", cl.RoomsController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
}
