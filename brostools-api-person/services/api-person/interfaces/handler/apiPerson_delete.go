package handler

import (
	"brostools-api-person/lib/log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (rh *apiPersonHandler) HandleDeletePerson() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		log.Infof(ctx, "Handle Delete Person")
		response := NewResponse()

		clientcd := c.Param("client_cd")
		personcd := c.Param("person_cd")

		status, error := rh.expFr.Delete(clientcd, personcd)
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
