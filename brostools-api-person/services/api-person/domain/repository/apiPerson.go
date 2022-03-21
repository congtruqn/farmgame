package repository

import "brostools-api-person/domain/model"

type ApiPersonRepository interface {
	Open() error
	Close()
	Begin() error
	Rollback()
	Commit() error
	Add(ri *model.ApiPerson) (int, error)
	BrosUpdate(client_cd string, person_cd string, ri *model.ApiPerson) (int, error)
	BrosGetById(client_cd string, person_cd string) (int, error, *model.ApiPerson)
	BrosDelete(client_cd string, person_cd string) (int, error)
	BrosGetAll() (int, error, model.BrosMapSet)
}
