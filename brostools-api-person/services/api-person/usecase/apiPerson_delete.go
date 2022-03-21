package usecase

import (
	"brostools-api-person/lib/log"
	"fmt"
	"net/http"
)

func (uc *apiPersonUsecase) Delete(client_cd string, person_cd string) (int, error) {

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

	status, errUpdate := uc.apiPersonRepo.BrosDelete(client_cd, person_cd)
	if errUpdate != nil {
		log.Infof(nil, "Error Delete Person Bros")
		uc.apiPersonRepo.Rollback()
		return status, err
	}

	// Add new Person Brantect
	statusbrantect, errbrantect := uc.apiPersonRepoBrantect.BrantectDelete(client_cd, person_cd)
	if errbrantect != nil {
		log.Infof(nil, "Error Delete person Brantect")
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
