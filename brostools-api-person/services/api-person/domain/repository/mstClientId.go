package repository

import (
	"brostools-api-person/domain/model"
)

type MstClientIdRepository interface {
	FindByClientCd(client_cd string) (*model.MstClientId, error)
	FindAllClient() ([]*model.MstClientId, error)
}
