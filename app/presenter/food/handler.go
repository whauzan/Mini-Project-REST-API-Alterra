package food

import (
	// _request "miniproject/app/presenter/food/request"
	// _response "miniproject/app/presenter/food/response"
	"miniproject/business/food"
	// "miniproject/helper"
	// "net/http"

	// "github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceFood food.Service
}

func NewHandler(foodService food.Service) *Presenter {
	return &Presenter{
		serviceFood: foodService,
	}
}

// func (handler *Presenter) GetRecipeByName(echoContext echo.Context) error {
// 	name := echoContext.QueryParam("name")
// 	resp, err := handler.serviceFood.GetRecipeByName(name)
// 	if err != nil {
// 		response := helper.APIResponse("Failed Get Food", http.StatusBadRequest, "Error", err)
// 		return echoContext.JSON(http.StatusBadRequest, response)
// 	}
// 	response := helper.APIResponse("Succes", http.StatusOK, "Succes", _response.FromDomain(resp))
// 	return echoContext.JSON(http.StatusBadRequest, response)
// }

// func (handler *Presenter) SaveFood(echoContext echo.Context) error {
// 	var req _request.Food
// 	if err := echoContext.Bind(&req); err != nil {
// 		response := helper.APIResponse("Failed Save Food", http.StatusBadRequest, "Error", err)
// 		return echoContext.JSON(http.StatusBadRequest, response)
// 	}
// 	domain := _request.ToDomain(req)
// 	resp, err := handler.serviceFood.SaveFood(_response.FromDomainAPI(domain))
// 	if err != nil {
// 		response := helper.APIResponse("Failed Save Food", http.StatusBadRequest, "Error", err)
// 		return echoContext.JSON(http.StatusBadRequest, response)
// 	}
// 	response := helper.APIResponse("Success Save Food", http.StatusOK, "Success", _response.FromDomain(resp))
// 	return echoContext.JSON(http.StatusBadRequest, response)
// }

// func (handler *Presenter) GetRecipeByName(c echo.Context) error {
// 	data, err := handler.serviceFood.GetRecipeByName()
// }