package usecase

import (
	"net/http"
)

func (uc *apiPersonUsecase) GetByID(client_cd string, person_cd string) (int, error, *ApiPerson) {
	var resApiPerson ApiPerson
	status, err, prosPerson := uc.apiPersonRepo.BrosGetById(client_cd, person_cd)
	if err != nil {
		return http.StatusInternalServerError, err, &resApiPerson
	}
	if status == 404 {
		return status, nil, &resApiPerson
	}
	statusBrantect, errbrantect, brantectPerson := uc.apiPersonRepoBrantect.BrantectGetByID(client_cd, person_cd)
	if errbrantect != nil {
		return http.StatusInternalServerError, errbrantect, &resApiPerson
	}
	if statusBrantect == 404 {
		//return statusBrantect, nil, &resApiPerson
	}
	result := &ApiPerson{
		ClientCd:         prosPerson.ClientCd,
		DeptCd:           brantectPerson.DeptCd,
		BrantectPersonCd: brantectPerson.BrantectPersonCd,
		BrosPersonCd:     prosPerson.BrosPersonCd,
		PersonField:      prosPerson.PersonField,
		DivisionNm:       prosPerson.DivisionNm,
		PositionNm:       prosPerson.PositionNm,
		PersonNm:         prosPerson.PersonNm,
		PersonNmJp:       prosPerson.PersonNmJp,
		PostCd:           prosPerson.PostCd,
		Address:          prosPerson.Address,
		Tel:              prosPerson.Tel,
		Fax:              prosPerson.Fax,
		Email:            prosPerson.Email,
		Mobile:           prosPerson.Mobile,
		BrantectRemarks:  brantectPerson.BrantectRemarks,
		BrosRemarks:      prosPerson.BrosRemarks,
		InvsndDnFlg:      prosPerson.InvsndDnFlg,
		InvsndTmFlg:      prosPerson.InvsndTmFlg,
		InvsndOtFlg:      prosPerson.InvsndOtFlg,
		SignFlg:          prosPerson.SignFlg,
		SearchPersonNm:   brantectPerson.SearchPersonNm,
		Type:             brantectPerson.Type,
		Idno:             brantectPerson.Idno,
		LastVerifiedDate: brantectPerson.LastVerifiedDate,
		DeleteFlg:        prosPerson.DeleteFlg,
		BrantectUpdDate:  brantectPerson.UpdDate,
		BrantectUpdUser:  brantectPerson.UpdUser,
		BrosUpdDate:      prosPerson.UpdDate,
		BrosUpdUser:      prosPerson.UpdUser,
		BrosUpdPrgId:     prosPerson.UpdPrgId,
		BrantectInpDate:  brantectPerson.InpDate,
		BrantectInpUser:  brantectPerson.InpUser,
		BrosInpDate:      prosPerson.InpDate,
		BrosInpUser:      prosPerson.InpUser,
		BrosInpPrgId:     prosPerson.InpPrgId,
	}
	return http.StatusOK, nil, result
}
