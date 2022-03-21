package request

import (
	"brostools-api-person/usecase"

	"github.com/go-playground/validator"
)

type ReqSign struct {
	ClientId string            `json:"client_id" validate:"required"`
	UserId   string            `json:"user_id" validate:"required"`
	Db       string            `json:"db" validate:"required"`
	Role     string            `json:"role" validate:"required"`
	Payloads map[string]string `json:"payloads"`
}

type ReqVerify struct {
	Jwt string `json:"jwt" validate:"required"`
}

func ReqSignValidate(req ReqSign) error {
	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}

func ReqVerifyValidate(req ReqVerify) error {
	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}
func ReqPostApiPersonValidate(rcvinv *usecase.ApiPerson) error {
	validate := validator.New()
	err := validate.Struct(rcvinv)
	if err != nil {
		return err
	}
	return nil
}
func ReqUpdateApiPersonValidate(rcvinv *usecase.ApiPersonUpdate) error {
	validate := validator.New()
	err := validate.Struct(rcvinv)
	if err != nil {
		return err
	}
	return nil
}
