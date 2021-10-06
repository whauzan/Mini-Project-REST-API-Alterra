package routes

import (
	"miniproject/app/presenter/food"

	"github.com/labstack/echo/v4"
)

type HandlerList struct {
	OpenAPIHandler food.Presenter
}

func (handler *HandlerList) RouteReg(e *echo.Echo) {
	api := e.Group("foodie/v1")
	api.GET("/food/", handler.OpenAPIHandler.GetRecipeByName)
	api.POST("/food/save", handler.OpenAPIHandler.SaveFood)
}
