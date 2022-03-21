package handler

import (
	"brostools-api-person/interfaces/util"
	"brostools-api-person/lib/log"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetAllResponse struct {
	Errors []util.Error `json:"errors"`
	Data   interface{}  `json:"person"`
}

func NewGetAllResponse() GetByIdResponse {
	var res GetByIdResponse
	res.Errors = []util.Error{}
	res.Data = make([]interface{}, 0)
	return res
}
func (res *GetAllResponse) AddError(code int, message string) {
	var err util.Error
	err.Code = code
	err.Message = message
	res.Errors = append(res.Errors, err)
}
func (res *GetAllResponse) ToJson() string {
	json, _ := json.Marshal(res)
	return string(json)
}
func (rh *apiPersonHandler) HandleGetAllPerson() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		log.Infof(ctx, "Handle GetAll Person")
		response := NewGetByIdResponse()
		status, getError, resApiPerson := rh.expFr.GetAll()
		if getError != nil {
			log.Infof(ctx, "Handle GetAll PersonError: "+getError.Error())
			response.AddError(http.StatusInternalServerError, getError.Error())
			return c.JSON(status, response)
		}
		if status == http.StatusNotFound {
			response.AddError(http.StatusNotFound, "Person not found")
			return c.JSON(status, response)
		}
		response.Data = resApiPerson
		return c.JSON(status, response)
	}
}
