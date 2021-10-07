package routes

import (
	"miniproject/app/presenter/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserHandler user.UserHandler
}

func (handlerList *HandlerList) RouteRegister(e *echo.Echo) {
	usr := e.Group("users")
	usr.POST("/register", handlerList.UserHandler.Register)
	usr.POST("/login", handlerList.UserHandler.Login)
}
