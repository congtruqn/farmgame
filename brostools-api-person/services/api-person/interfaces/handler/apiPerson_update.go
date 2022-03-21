package handler

import (
	"brostools-api-person/interfaces/request"
	"brostools-api-person/lib/log"
	"brostools-api-person/usecase"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (rh *apiPersonHandler) HandleUpdatePerson() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		log.Infof(ctx, "HandleAddPerson")
		response := NewResponse()
		var updatePerson usecase.ApiPersonUpdate

		clientcd := c.Param("client_cd")
		personcd := c.Param("person_cd")

		if err := c.Bind(&updatePerson); err != nil {
			response.AddError(http.StatusBadRequest, err.Error())
			log.Infof(ctx, "HandleAddPerson")
			return c.JSON(http.StatusBadRequest, response)
		}

		validateErr := request.ReqUpdateApiPersonValidate(&updatePerson)
		if validateErr != nil {
			validateErrList := validateErr.(validator.ValidationErrors)
			for _, f := range validateErrList {
				response.AddError(800, f.Field()+": "+f.Tag())
			}
			log.Infof(ctx, "Validate error: "+response.ToJson())

			return c.JSON(http.StatusBadRequest, response)
		}
		status, error := rh.expFr.Update(clientcd, personcd, &updatePerson)
		if error != nil {
			response.AddError(status, error.Error())
			return c.JSON(status, response)
		}
		if status == http.StatusNotFound {
			response.AddError(status, "Person not found")
			return c.JSON(status, response)
		}
		return c.JSON(status, response)
	}
}
