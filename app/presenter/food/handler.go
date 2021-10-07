package food

import (
	_handler "miniproject/app/presenter"
	"miniproject/app/presenter/food/request"
	"miniproject/business/food"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FoodHandler struct {
	FoodService food.Service
}

func NewFoodHandler(service food.Service) *FoodHandler {
	return &FoodHandler{
		FoodService: service,
	}
}

func (handler *FoodHandler) SaveFood(ctx echo.Context) error {
	req := request.FoodSave{}
	if err := ctx.Bind(&req); err != nil {
		return _handler.NewErrorResponse(ctx, http.StatusBadRequest, err)
	}
	data, err := handler.FoodService.SaveFood(req.Food_Name, req.Food_ID)
	if err != nil {
		return _handler.NewErrorResponse(ctx, http.StatusBadRequest, err)
	}
	return _handler.NewSuccessResponse(ctx, data)
}
