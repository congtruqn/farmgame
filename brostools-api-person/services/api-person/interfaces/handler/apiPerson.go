package handler

import (
	"brostools-api-person/usecase"

	"github.com/labstack/echo/v4"
)

type ApiPersonHandler interface {
	HandleAddPerson() echo.HandlerFunc
	HandleUpdatePerson() echo.HandlerFunc
	HandleGePersontById() echo.HandlerFunc
	HandleDeletePerson() echo.HandlerFunc
	HandleGetAllPerson() echo.HandlerFunc
}

type apiPersonHandler struct {
	expFr usecase.ApiPersonUsecase
}

func NewApiPersonHandler(exp usecase.ApiPersonUsecase) ApiPersonHandler {
	return &apiPersonHandler{expFr: exp}
}
