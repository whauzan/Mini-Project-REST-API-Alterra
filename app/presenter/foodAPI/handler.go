package foodAPI

import (
	// "encoding/json"
	// "log"
	"miniproject/app/presenter/foodAPI/request"
	"miniproject/business/recipesAPI"
	_handler "miniproject/app/presenter"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FoodAPIHandler struct {
	FoodAPIRepo recipesAPI.Repository
}

func NewFoodAPIHandler(repo recipesAPI.Repository) *FoodAPIHandler {
	return &FoodAPIHandler{
		FoodAPIRepo: repo,
	}
}

func (handler *FoodAPIHandler) GetRecipeAPI(ctx echo.Context) error {
	req := request.FoodAPISearch{}

	if err := ctx.Bind(&req); err != nil {
		return _handler.NewErrorResponse(ctx, http.StatusBadRequest, err)
	}
	data, err := handler.FoodAPIRepo.GetRecipeAPI(req.Name)
	if err != nil {
		return _handler.NewErrorResponse(ctx, http.StatusInternalServerError, err)
	}
	// dataJSON, err := json.Marshal(data)
	// if err != nil {
	// 	log.Fatal("Cannot encode to JSON ", err)
	// }
	return _handler.NewSuccessResponse(ctx, data)
}
