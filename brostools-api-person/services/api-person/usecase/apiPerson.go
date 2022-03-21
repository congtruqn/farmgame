package usecase

import (
	"brostools-api-person/domain/repository"
)

type ApiPerson struct {
	ClientCd         string `json:"client_cd" validate:"required"`
	DeptCd           string `json:"dept_cd"`
	BrantectPersonCd string `json:"brantect_person_cd" validate:"required"`
	BrosPersonCd     string `json:"bros_person_cd" validate:"required"`
	PersonField      string `json:"person_field"`
	DivisionNm       string `json:"division_nm"`
	PositionNm       string `json:"position_nm"`
	PersonNm         string `json:"person_nm"`
	PersonNmJp       string `json:"person_nm_jp"`
	PostCd           string `json:"post_cd"`
	Address          string `json:"address"`
	Tel              string `json:"tel"`
	Fax              string `json:"fax"`
	Email            string `json:"email"`
	Mobile           string `json:"mobile"`
	BrantectRemarks  string `json:"brantect_remarks"`
	BrosRemarks      string `json:"bros_remarks"`
	InvsndDnFlg      string `json:"invsnd_dn_flg"`
	InvsndTmFlg      string `json:"invsnd_tm_flg"`
	InvsndOtFlg      string `json:"invsnd_ot_flg"`
	SignFlg          string `json:"sign_flg"`
	SearchPersonNm   string `json:"search_person_nm"`
	Type             string `json:"type"`
	Idno             string `json:"idno"`
	LastVerifiedDate string `json:"last_verified_date"`
	DeleteFlg        string `json:"delete_flg"`
	BrantectUpdDate  string `json:"brantect_upd_date"`
	BrantectUpdUser  string `json:"brantect_upd_user"`
	BrosUpdDate      string `json:"bros_upd_date"`
	BrosUpdUser      string `json:"bros_upd_user"`
	BrosUpdPrgId     string `json:"bros_upd_prg_id"`
	BrantectInpDate  string `json:"brantect_inp_date"`
	BrantectInpUser  string `json:"brantect_inp_user"`
	BrosInpDate      string `json:"bros_inp_date"`
	BrosInpUser      string `json:"bros_inp_user"`
	BrosInpPrgId     string `json:"bros_inp_prg_id"`
}

type ApiPersonUsecase interface {
	Add(ri *ApiPerson) (int, error)
	Update(client_cd string, person_cd string, ri *ApiPersonUpdate) (int, error)
	GetByID(client_cd string, person_cd string) (int, error, *ApiPerson)
	Delete(client_cd string, person_cd string) (int, error)
	GetAll() (int, error, []ApiPerson)
}

type apiPersonUsecase struct {
	apiPersonRepo         repository.ApiPersonRepository
	apiPersonRepoBrantect repository.ApiPersonBrantectRepository
}

func NewApiPersonUsecase(apiPerson repository.ApiPersonRepository, apiPersonBantect repository.ApiPersonBrantectRepository) ApiPersonUsecase {
	return &apiPersonUsecase{
		apiPersonRepo:         apiPerson,
		apiPersonRepoBrantect: apiPersonBantect,
	}
}
