package presenter

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"rc"`
		Message  string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(ctx echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = param

	return ctx.JSON(http.StatusOK, response)
}

func NewErrorResponse(ctx echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "Error occured"
	response.Meta.Messages = []string{err.Error()}

	return ctx.JSON(status, response)
}
