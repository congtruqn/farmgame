package repository

import (
	"brostools-api-person/domain/model"
)

type ApiPersonBrantectRepository interface {
	Connect(client_cd string) error
	ConnectToGeneralPgSql() error
	ConnectToPrivatePgSql(client_id *model.MstClientId) error
	Close()
	Begin() error
	Rollback()
	Commit() error
	AddBrantect(ri *model.ApiPersonBrantect) (int, error)
	BrantectUpdate(client_cd string, person_cd string, ri *model.ApiPersonBrantect) (int, error)
	BrantectGetByID(client_cd string, person_cd string) (int, error, *model.ApiPersonBrantect)
	BrantectDelete(client_cd string, person_cd string) (int, error)
	BrantectGetAll() (int, error, model.TCPSet)
}
