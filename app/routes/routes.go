package routes

import (
	"miniproject/controller/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController
}

func (ctrlList *ControllerList) RouteRegister(e *echo.Echo) {
	usr := e.Group("users")
	usr.POST("/register", ctrlList.UserController.Register)
	usr.POST("/login", ctrlList.UserController.Login)
}
