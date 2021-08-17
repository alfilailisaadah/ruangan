package routes

import (
	"rentRoom/controllers/buildings"
	"rentRoom/controllers/rents"
	"rentRoom/controllers/rooms"
	"rentRoom/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	UserController     users.UserController
	BuildingsController     buildings.BuildingsController
	RoomsController rooms.RoomsController
	RentsController rents.RentsController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", cl.UserController.Store)
	users.GET("/token", cl.UserController.CreateToken)

	rooms := e.Group("rooms")
	rooms.GET("/list", cl.RoomsController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))
	rooms.POST("/store", cl.RoomsController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))

	buildings := e.Group("buildings")
	buildings.GET("/list", cl.BuildingsController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))
	buildings.POST("/store", cl.BuildingsController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))

	
	rents := e.Group("rents")
	rents.POST("/store", cl.RentsController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
}
