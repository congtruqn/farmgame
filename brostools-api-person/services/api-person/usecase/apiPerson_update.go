package usecase

import (
	"brostools-api-person/domain/model"
	"brostools-api-person/lib/log"
	"fmt"
	"net/http"
)

type ApiPersonUpdate struct {
	DeptCd          string `json:"dept_cd"`
	PersonCd        string `json:"person_cd" validate:"required"`
	PersonField     string `json:"person_field"`
	DivisionNm      string `json:"division_nm"`
	PositionNm      string `json:"position_nm"`
	PersonNm        string `json:"person_nm"`
	PersonNmJp      string `json:"person_nm_jp"`
	PostCd          string `json:"post_cd"`
	Address         string `json:"address"`
	Tel             string `json:"tel"`
	Fax             string `json:"fax"`
	Email           string `json:"email"`
	Mobile          string `json:"mobile"`
	BrantectRemarks string `json:"brantect_remarks"`
	BrosRemarks     string `json:"bros_remarks"`
	InvsndDnFlg     string `json:"invsnd_dn_flg"`
	InvsndTmFlg     string `json:"invsnd_tm_flg"`
	InvsndOtFlg     string `json:"invsnd_ot_flg"`
	SignFlg         string `json:"sign_flg"`
	SearchPersonNm  string `json:"search_person_nm"`
	Type            string `json:"type"`
	Idno            string `json:"idno"`
}

func (uc *apiPersonUsecase) Update(client_cd string, person_cd string, ri *ApiPersonUpdate) (int, error) {

	status, err, prosPerson := uc.apiPersonRepo.BrosGetById(client_cd, person_cd)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if status == 404 {
		return status, nil
	}
	statusBrantect, errbrantect, brantectPerson := uc.apiPersonRepoBrantect.BrantectGetByID(client_cd, person_cd)
	if errbrantect != nil {
		return http.StatusInternalServerError, errbrantect
	}
	if statusBrantect == 404 {
		return statusBrantect, nil
	}
	fmt.Println(brantectPerson)
	fmt.Println(prosPerson)
	// Connect to Brostools
	errOpen := uc.apiPersonRepo.Open()
	if errOpen != nil {
		return http.StatusInternalServerError, errOpen
	}
	errBegin := uc.apiPersonRepo.Begin()
	if errBegin != nil {
		log.Infof(nil, "Error errBegin")
		return http.StatusInternalServerError, errBegin
	}
	//End connect to Brostools

	// Connect to Brantect
	errOpenBr := uc.apiPersonRepoBrantect.Connect(client_cd)
	if errOpen != nil {
		log.Infof(nil, "Error Connect Bantect")
		return http.StatusInternalServerError, errOpenBr
	}

	errBeginBran := uc.apiPersonRepoBrantect.Begin()
	if errBegin != nil {
		log.Infof(nil, "Error errBegin Bantect")
		return http.StatusInternalServerError, errBeginBran
	}
	//End connect to Brantect

	personBros := &model.ApiPerson{
		BrosPersonCd: ri.PersonCd,
		PersonField:  ri.PersonField,
		DivisionNm:   ri.DivisionNm,
		PositionNm:   ri.PositionNm,
		PersonNm:     ri.PersonNm,
		PersonNmJp:   ri.PersonNmJp,
		PostCd:       ri.PostCd,
		Address:      ri.Address,
		Tel:          ri.Tel,
		Fax:          ri.Fax,
		Email:        ri.Email,
		Mobile:       ri.Mobile,
		BrosRemarks:  ri.BrosRemarks,
		InvsndDnFlg:  ri.InvsndDnFlg,
		InvsndTmFlg:  ri.InvsndTmFlg,
		InvsndOtFlg:  ri.InvsndOtFlg,
		SignFlg:      ri.SignFlg,
	}

	personBrantect := &model.ApiPersonBrantect{
		DeptCd:           ri.DeptCd,
		BrantectPersonCd: ri.PersonCd,
		PositionNm:       ri.PositionNm,
		PersonNm:         ri.PersonNm,
		PersonNmJp:       ri.PersonNmJp,
		PostCd:           ri.PostCd,
		Address:          ri.Address,
		Tel:              ri.Tel,
		Fax:              ri.Fax,
		Email:            ri.Email,
		Mobile:           ri.Mobile,
		BrantectRemarks:  ri.BrantectRemarks,
		SearchPersonNm:   ri.SearchPersonNm,
		Type:             ri.Type,
		Idno:             ri.Idno,
	}
	// Add new Person Bros
	status, errUpdate := uc.apiPersonRepo.BrosUpdate(client_cd, person_cd, personBros)
	if errUpdate != nil {
		log.Infof(nil, "Error Update Person Bros")
		uc.apiPersonRepo.Rollback()
		return status, err
	}

	// Add new Person Brantect
	statusbrantect, errbrantect := uc.apiPersonRepoBrantect.BrantectUpdate(client_cd, person_cd, personBrantect)
	if errbrantect != nil {
		log.Infof(nil, "Error Update person Brantect")
		uc.apiPersonRepo.Rollback()
		uc.apiPersonRepoBrantect.Rollback()
		return statusbrantect, errbrantect
	}

	errCommitBros := uc.apiPersonRepo.Commit()
	if errCommitBros != nil {
		log.Infof(nil, "Error Commit ")
		uc.apiPersonRepo.Rollback()
		uc.apiPersonRepoBrantect.Rollback()
		return http.StatusInternalServerError, errCommitBros
	}
	errCommitBrantect := uc.apiPersonRepoBrantect.Commit()
	if errCommitBrantect != nil {
		log.Infof(nil, "Error Commit")
		uc.apiPersonRepo.Rollback()
		uc.apiPersonRepoBrantect.Rollback()
		return http.StatusInternalServerError, errCommitBrantect
	}
	return 0, err
}
