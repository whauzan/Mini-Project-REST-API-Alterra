package routes

import (
	"miniproject/app/presenter/food"
	"miniproject/app/presenter/foodAPI"

	"github.com/labstack/echo/v4"
)

type HandlerList struct {
	FoodAPIHandler foodAPI.FoodAPIHandler
	FoodHandler food.FoodHandler
}

func (handler *HandlerList) RouteReg(e *echo.Echo) {
	api := e.Group("foodie/v1")
	api.GET("/foodSearch", handler.FoodAPIHandler.GetRecipeAPI)
	api.POST("/SaveFood", handler.FoodHandler.SaveFood)
}
