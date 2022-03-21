package handler

import (
	"brostools-api-person/interfaces/request"
	"brostools-api-person/interfaces/util"
	"brostools-api-person/lib/log"
	"brostools-api-person/usecase"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Errors []util.Error         `json:"errors"`
	Data   []*usecase.ApiPerson `json:"data"`
}

func NewResponse() Response {
	var res Response
	res.Errors = []util.Error{}
	res.Data = []*usecase.ApiPerson{}
	return res
}

func (res *Response) AddError(code int, message string) {
	var err util.Error
	err.Code = code
	err.Message = message
	res.Errors = append(res.Errors, err)
}
func (res *Response) ToJson() string {
	json, _ := json.Marshal(res)
	return string(json)
}
func (rh *apiPersonHandler) HandleAddPerson() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		log.Infof(ctx, "HandleAddPerson")
		response := NewResponse()
		var postPerson usecase.ApiPerson
		if err := c.Bind(&postPerson); err != nil {
			response.AddError(http.StatusBadRequest, err.Error())
			log.Infof(ctx, "HandleAddPerson")
			return c.JSON(http.StatusBadRequest, response)
		}
		validateErr := request.ReqPostApiPersonValidate(&postPerson)
		if validateErr != nil {
			validateErrList := validateErr.(validator.ValidationErrors)
			for _, f := range validateErrList {
				response.AddError(800, f.Field()+": "+f.Tag())
			}
			log.Infof(ctx, "Validate error: "+response.ToJson())

			return c.JSON(http.StatusBadRequest, response)
		}
		status, error := rh.expFr.Add(&postPerson)
		if error != nil {
			response.AddError(status, error.Error())
			log.Infof(ctx, "HandleAddPerson")
			return c.JSON(status, response)
		}
		return c.JSON(status, response)
	}
}
