package usecase

import (
	"brostools-api-person/domain/model"
	"brostools-api-person/lib/log"
	"net/http"
)

func (uc *apiPersonUsecase) Add(ri *ApiPerson) (int, error) {
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
	errOpenBr := uc.apiPersonRepoBrantect.Connect(ri.ClientCd)
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
		ClientCd:     ri.ClientCd,
		BrosPersonCd: ri.BrosPersonCd,
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
		ClientCd:         ri.ClientCd,
		DeptCd:           ri.DeptCd,
		BrantectPersonCd: ri.BrantectPersonCd,
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
	status, err := uc.apiPersonRepo.Add(personBros)
	if err != nil {
		log.Infof(nil, "Error Add Person Bros")
		uc.apiPersonRepo.Rollback()
		return status, err
	}

	// Add new Person Brantect
	statusbrantect, errbrantect := uc.apiPersonRepoBrantect.AddBrantect(personBrantect)
	if errbrantect != nil {
		log.Infof(nil, "Error Add person Brantect")
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
func (uc *apiPersonUsecase) BrosUpdate(ri *ApiPerson) (int, error) {
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
	errOpenBr := uc.apiPersonRepoBrantect.Connect(ri.ClientCd)
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
		ClientCd:     ri.ClientCd,
		BrosPersonCd: ri.BrosPersonCd,
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
		ClientCd:         ri.ClientCd,
		DeptCd:           ri.DeptCd,
		BrantectPersonCd: ri.BrantectPersonCd,
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
	status, err := uc.apiPersonRepo.Add(personBros)
	if err != nil {
		log.Infof(nil, "Error Add Person Bros")
		uc.apiPersonRepo.Rollback()
		return status, err
	}

	// Add new Person Brantect
	statusbrantect, errbrantect := uc.apiPersonRepoBrantect.AddBrantect(personBrantect)
	if errbrantect != nil {
		log.Infof(nil, "Error Add person Brantect")
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
