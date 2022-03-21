package handler

import (
	"brostools-api-person/interfaces/util"
	"brostools-api-person/lib/log"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetByIdResponse struct {
	Errors []util.Error `json:"errors"`
	Data   interface{}  `json:"person"`
}

func NewGetByIdResponse() GetByIdResponse {
	var res GetByIdResponse
	res.Errors = []util.Error{}
	res.Data = make([]interface{}, 0)
	return res
}
func (res *GetByIdResponse) AddError(code int, message string) {
	var err util.Error
	err.Code = code
	err.Message = message
	res.Errors = append(res.Errors, err)
}
func (res *GetByIdResponse) ToJson() string {
	json, _ := json.Marshal(res)
	return string(json)
}
func (rh *apiPersonHandler) HandleGePersontById() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		log.Infof(ctx, "Handle Get Person By ID")
		response := NewGetByIdResponse()
		clientCd := c.Param("client_cd")
		personcd := c.Param("person_cd")
		status, getError, resApiPerson := rh.expFr.GetByID(clientCd, personcd)
		if getError != nil {
			log.Infof(ctx, "Handle Get Person By ID Error: "+getError.Error())
			response.AddError(http.StatusBadRequest, getError.Error())
			return c.JSON(status, response)
		}
		if status == http.StatusNotFound {
			response.AddError(http.StatusNotFound, "Person not found")
			return c.JSON(status, response)
		}
		response.Data = *resApiPerson
		return c.JSON(status, response)
	}
}
