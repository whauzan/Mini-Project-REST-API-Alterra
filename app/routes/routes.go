package routes

import (
	"miniproject/business/foods"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	OpenAPIController foods.Repository
}

func (c *ControllerList) RouteReg(e *echo.Echo) {
	api := e.Group("openAPI")
	api.GET("/", c.OpenAPIController.GetRecipeByName)
}
