package usecase

import (
	"brostools-api-person/domain/model"
	"encoding/json"
	"net/http"
)

func (uc *apiPersonUsecase) GetAll() (int, error, []ApiPerson) {
	var resApiPerson []ApiPerson
	status, err, prosPerson := uc.apiPersonRepo.BrosGetAll()
	if err != nil {
		return http.StatusInternalServerError, err, resApiPerson
	}
	if status == 404 {
		//return status, nil, resApiPerson
	}
	statusBrantect, errbrantect, brantectPerson := uc.apiPersonRepoBrantect.BrantectGetAll()

	if errbrantect != nil {
		return http.StatusInternalServerError, errbrantect, resApiPerson
	}
	if statusBrantect == 404 {
		//return statusBrantect, nil, &resApiPerson
	}
	for k := range prosPerson {
		jsonString, _ := json.Marshal(prosPerson[k])
		var s model.ApiPerson
		json.Unmarshal([]byte(jsonString), &s)

		brantectjsonString, _ := json.Marshal(brantectPerson[k])
		var brante model.ApiPersonBrantect
		json.Unmarshal([]byte(brantectjsonString), &brante)
		result := &ApiPerson{}
		if brante.ClientCd == "" {
			result = &ApiPerson{
				ClientCd:         s.ClientCd,
				DeptCd:           "",
				BrantectPersonCd: "",
				BrosPersonCd:     s.BrosPersonCd,
				PersonField:      s.PersonField,
				DivisionNm:       s.DivisionNm,
				PositionNm:       s.PositionNm,
				PersonNm:         s.PersonNm,
				PersonNmJp:       s.PersonNmJp,
				PostCd:           s.PostCd,
				Address:          s.Address,
				Tel:              s.Tel,
				Fax:              s.Fax,
				Email:            s.Email,
				Mobile:           s.Mobile,
				BrantectRemarks:  "",
				BrosRemarks:      s.BrosRemarks,
				InvsndDnFlg:      s.InvsndDnFlg,
				InvsndTmFlg:      s.InvsndTmFlg,
				InvsndOtFlg:      s.InvsndOtFlg,
				SignFlg:          s.SignFlg,
				SearchPersonNm:   "",
				Type:             "",
				Idno:             "",
				LastVerifiedDate: "",
				DeleteFlg:        s.DeleteFlg,
				BrantectUpdDate:  "",
				BrantectUpdUser:  "",
				BrosUpdDate:      s.UpdDate,
				BrosUpdUser:      s.UpdUser,
				BrosUpdPrgId:     s.UpdPrgId,
				BrantectInpDate:  "",
				BrantectInpUser:  "",
				BrosInpDate:      s.InpDate,
				BrosInpUser:      s.InpUser,
				BrosInpPrgId:     s.InpPrgId,
			}
		} else {
			result = &ApiPerson{
				ClientCd:         s.ClientCd,
				DeptCd:           brante.DeptCd,
				BrantectPersonCd: brante.BrantectPersonCd,
				BrosPersonCd:     s.BrosPersonCd,
				PersonField:      s.PersonField,
				DivisionNm:       s.DivisionNm,
				PositionNm:       s.PositionNm,
				PersonNm:         s.PersonNm,
				PersonNmJp:       s.PersonNmJp,
				PostCd:           s.PostCd,
				Address:          s.Address,
				Tel:              s.Tel,
				Fax:              s.Fax,
				Email:            s.Email,
				Mobile:           s.Mobile,
				BrantectRemarks:  brante.BrantectRemarks,
				BrosRemarks:      s.BrosRemarks,
				InvsndDnFlg:      s.InvsndDnFlg,
				InvsndTmFlg:      s.InvsndTmFlg,
				InvsndOtFlg:      s.InvsndOtFlg,
				SignFlg:          s.SignFlg,
				SearchPersonNm:   brante.SearchPersonNm,
				Type:             brante.Type,
				Idno:             brante.Idno,
				LastVerifiedDate: brante.LastVerifiedDate,
				DeleteFlg:        s.DeleteFlg,
				BrantectUpdDate:  brante.UpdDate,
				BrantectUpdUser:  brante.UpdUser,
				BrosUpdDate:      s.UpdDate,
				BrosUpdUser:      s.UpdUser,
				BrosUpdPrgId:     s.UpdPrgId,
				BrantectInpDate:  brante.InpDate,
				BrantectInpUser:  brante.InpUser,
				BrosInpDate:      s.InpDate,
				BrosInpUser:      s.InpUser,
				BrosInpPrgId:     s.InpPrgId,
			}

		}
		resApiPerson = append(resApiPerson, *result)
		delete(brantectPerson, k)
		//fmt.Printf("key[%s] value[%s]\n", k, brantectPerson[k])

	}
	for x := range brantectPerson {
		brantectjsonString, _ := json.Marshal(brantectPerson[x])
		var brantec model.ApiPersonBrantect
		json.Unmarshal([]byte(brantectjsonString), &brantec)
		result := &ApiPerson{}
		result = &ApiPerson{
			ClientCd:         brantec.ClientCd,
			DeptCd:           brantec.DeptCd,
			BrantectPersonCd: brantec.BrantectPersonCd,
			BrosPersonCd:     "",
			PersonField:      "",
			DivisionNm:       "",
			PositionNm:       brantec.PositionNm,
			PersonNm:         brantec.PersonNm,
			PersonNmJp:       brantec.PersonNmJp,
			PostCd:           brantec.PostCd,
			Address:          brantec.Address,
			Tel:              brantec.Tel,
			Fax:              brantec.Fax,
			Email:            brantec.Email,
			Mobile:           brantec.Mobile,
			BrantectRemarks:  brantec.BrantectRemarks,
			BrosRemarks:      "",
			InvsndDnFlg:      "",
			InvsndTmFlg:      "",
			InvsndOtFlg:      "",
			SignFlg:          "",
			SearchPersonNm:   brantec.SearchPersonNm,
			Type:             brantec.Type,
			Idno:             brantec.Idno,
			LastVerifiedDate: brantec.LastVerifiedDate,
			DeleteFlg:        brantec.DeleteFlg,
			BrantectUpdDate:  brantec.UpdDate,
			BrantectUpdUser:  brantec.UpdUser,
			BrosUpdDate:      "",
			BrosUpdUser:      "",
			BrosUpdPrgId:     "",
			BrantectInpDate:  brantec.InpDate,
			BrantectInpUser:  brantec.InpUser,
			BrosInpDate:      "",
			BrosInpUser:      "",
			BrosInpPrgId:     "",
		}
		resApiPerson = append(resApiPerson, *result)

	}
	return http.StatusOK, nil, resApiPerson
}
