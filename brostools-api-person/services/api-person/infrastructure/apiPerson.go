package infrastructure

import (
	"brostools-api-person/domain/model"
	"brostools-api-person/domain/repository"
	"brostools-api-person/lib/current"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type ApiPersonInfrastructure struct {
	User        string
	Password    string
	Connection  string
	Database    *sql.DB
	Transaction *sql.Tx
}

func NewApiPersonfrastructure() repository.ApiPersonRepository {
	return &ApiPersonInfrastructure{
		User:       os.Getenv("MYSQL_USER_BROSTOOLS_API_PERSON"),
		Password:   os.Getenv("MYSQL_PASSWORD_BROSTOOLS_API_PERSON"),
		Connection: os.Getenv("MYSQL_DSN_BROSTOOLS_API_PERSON"),
	}
}

func (infa *ApiPersonInfrastructure) Open() error {
	dsn := infa.User + ":" + infa.Password + "@" + infa.Connection
	db, errOpen := sql.Open("mysql", dsn)
	if errOpen != nil {
		return errOpen
	}
	infa.Database = db
	return nil
}

func (infa *ApiPersonInfrastructure) Close() {
	infa.Database.Close()
}

func (infa *ApiPersonInfrastructure) Begin() error {
	tx, errBegin := infa.Database.Begin()
	if errBegin != nil {
		return errBegin
	}
	infa.Transaction = tx
	return nil
}

func (infa *ApiPersonInfrastructure) Rollback() {
	infa.Transaction.Rollback()
}
func (infa *ApiPersonInfrastructure) Commit() error {
	return infa.Transaction.Commit()
}
func (repo *ApiPersonInfrastructure) Add(ri *model.ApiPerson) (int, error) {
	fmt.Println("Api Person Repository Add")

	sqlSelectTMNoSeqNo, errLoadFile := ioutil.ReadFile(current.GetCurrentDir() + "/infrastructure/sql/apiperson/brostools/AddApiPersonBros.sql")
	if errLoadFile != nil {
		return http.StatusInternalServerError, errLoadFile
	}
	defer repo.Close()
	stmt, errPrepare := repo.Transaction.Prepare(string(sqlSelectTMNoSeqNo))
	if errPrepare != nil {
		return http.StatusInternalServerError, errPrepare
	}

	_, errExec := stmt.Exec(
		&ri.ClientCd,
		&ri.BrosPersonCd,
		&ri.PersonField,
		&ri.DivisionNm,
		&ri.PositionNm,
		&ri.PersonNm,
		&ri.PersonNmJp,
		&ri.PostCd,
		&ri.Address,
		&ri.Tel,
		&ri.Fax,
		&ri.Email,
		&ri.Mobile,
		&ri.BrosRemarks,
		&ri.InvsndDnFlg,
		&ri.InvsndTmFlg,
		&ri.InvsndOtFlg,
		&ri.SignFlg,
		"0",
		"/brostools/api/person/",
		"/brostools/api/person/",
	)
	if errExec != nil {
		return http.StatusInternalServerError, errExec
	}
	return http.StatusOK, nil
}
func (repo *ApiPersonInfrastructure) BrosGetById(client_cd string, person_cd string) (int, error, *model.ApiPerson) {
	fmt.Println("Api Person Repository Get By ID")
	var brosPerson model.ApiPerson
	errOpen := repo.Open()
	if errOpen != nil {
		return http.StatusInternalServerError, errOpen, &brosPerson
	}
	defer repo.Close()
	sqlSelectTMNoSeqNo, errLoadFile := ioutil.ReadFile(current.GetCurrentDir() + "/infrastructure/sql/apiperson/brostools/ApiPersonBrosGetByID.sql")
	if errLoadFile != nil {
		return http.StatusInternalServerError, errLoadFile, &brosPerson
	}

	stmt, errPrepare := repo.Database.Prepare(string(sqlSelectTMNoSeqNo))
	if errPrepare != nil {
		return http.StatusInternalServerError, errPrepare, &brosPerson
	}

	queryErr := stmt.QueryRow(client_cd, person_cd).Scan(
		&brosPerson.ClientCd,
		&brosPerson.BrosPersonCd,
		&brosPerson.PersonField,
		&brosPerson.DivisionNm,
		&brosPerson.PositionNm,
		&brosPerson.PersonNm,
		&brosPerson.PersonNmJp,
		&brosPerson.PostCd,
		&brosPerson.Address,
		&brosPerson.Tel,
		&brosPerson.Fax,
		&brosPerson.Email,
		&brosPerson.Mobile,
		&brosPerson.BrosRemarks,
		&brosPerson.InvsndDnFlg,
		&brosPerson.InvsndTmFlg,
		&brosPerson.InvsndOtFlg,
		&brosPerson.SignFlg,
		&brosPerson.DeleteFlg,
		&brosPerson.UpdDate,
		&brosPerson.UpdUser,
		&brosPerson.UpdPrgId,
		&brosPerson.InpDate,
		&brosPerson.InpUser,
		&brosPerson.InpPrgId,
	)
	switch {
	case queryErr == sql.ErrNoRows:
		return http.StatusNotFound, nil, &brosPerson
	case queryErr != nil:
		return http.StatusInternalServerError, nil, &brosPerson
	default:
		return http.StatusOK, nil, &brosPerson
	}
}
func (repo *ApiPersonInfrastructure) BrosUpdate(client_cd string, person_cd string, ri *model.ApiPerson) (int, error) {
	fmt.Println("Api Person Repository Update")

	sqlSelect, errLoadFile := ioutil.ReadFile(current.GetCurrentDir() + "/infrastructure/sql/apiperson/brostools/UpdateApiPersonBros.sql")
	if errLoadFile != nil {
		return http.StatusInternalServerError, errLoadFile
	}
	defer repo.Close()
	stmt, errPrepare := repo.Transaction.Prepare(string(sqlSelect))
	if errPrepare != nil {
		return http.StatusInternalServerError, errPrepare
	}

	_, errExec := stmt.Exec(
		&ri.BrosPersonCd,
		&ri.PersonField,
		&ri.DivisionNm,
		&ri.PositionNm,
		&ri.PersonNm,
		&ri.PersonNmJp,
		&ri.PostCd,
		&ri.Address,
		&ri.Tel,
		&ri.Fax,
		&ri.Email,
		&ri.Mobile,
		&ri.BrosRemarks,
		&ri.InvsndDnFlg,
		&ri.InvsndTmFlg,
		&ri.InvsndOtFlg,
		&ri.SignFlg,
		"/brostools/api/person/",
		"/brostools/api/person/",
		client_cd,
		person_cd,
	)
	if errExec != nil {
		return http.StatusInternalServerError, errExec
	}
	return http.StatusOK, nil
}
func (repo *ApiPersonInfrastructure) BrosDelete(client_cd string, person_cd string) (int, error) {
	sqlSelect, errLoadFile := ioutil.ReadFile(current.GetCurrentDir() + "/infrastructure/sql/apiperson/brostools/DeleteApiPersonBros.sql")
	if errLoadFile != nil {
		return http.StatusInternalServerError, errLoadFile
	}
	defer repo.Close()
	stmt, errPrepare := repo.Transaction.Prepare(string(sqlSelect))
	if errPrepare != nil {
		return http.StatusInternalServerError, errPrepare
	}

	_, errExec := stmt.Exec(
		"/brostools/api/person/",
		"/brostools/api/person/",
		client_cd,
		person_cd,
	)
	if errExec != nil {
		return http.StatusInternalServerError, errExec
	}
	return http.StatusOK, nil
}
func (repo *ApiPersonInfrastructure) BrosGetAll() (int, error, model.BrosMapSet) {
	tmp := make(model.BrosMapSet)
	fmt.Println("Api Person Repository Get By ID")
	errOpen := repo.Open()
	if errOpen != nil {
		return http.StatusInternalServerError, errOpen, tmp
	}
	defer repo.Close()
	sqlSelect, errLoadFile := ioutil.ReadFile(current.GetCurrentDir() + "/infrastructure/sql/apiperson/brostools/ApiPersonBrosGetAll.sql")
	if errLoadFile != nil {
		return http.StatusInternalServerError, errLoadFile, tmp
	}

	stmt, errPrepare := repo.Database.Prepare(string(sqlSelect))
	if errPrepare != nil {
		return http.StatusInternalServerError, errPrepare, tmp
	}

	rows, queryErr := stmt.Query()
	switch {
	case queryErr == sql.ErrNoRows:
		return http.StatusNotFound, nil, tmp
	case queryErr != nil:
		return http.StatusInternalServerError, nil, tmp
	}
	for rows.Next() {
		var brosPerson model.ApiPerson
		if scanErr := rows.Scan(
			&brosPerson.ClientCd,
			&brosPerson.BrosPersonCd,
			&brosPerson.PersonField,
			&brosPerson.DivisionNm,
			&brosPerson.PositionNm,
			&brosPerson.PersonNm,
			&brosPerson.PersonNmJp,
			&brosPerson.PostCd,
			&brosPerson.Address,
			&brosPerson.Tel,
			&brosPerson.Fax,
			&brosPerson.Email,
			&brosPerson.Mobile,
			&brosPerson.BrosRemarks,
			&brosPerson.InvsndDnFlg,
			&brosPerson.InvsndTmFlg,
			&brosPerson.InvsndOtFlg,
			&brosPerson.SignFlg,
			&brosPerson.DeleteFlg,
			&brosPerson.UpdDate,
			&brosPerson.UpdUser,
			&brosPerson.UpdPrgId,
			&brosPerson.InpDate,
			&brosPerson.InpUser,
			&brosPerson.InpPrgId,
		); scanErr != nil {
			return http.StatusInternalServerError, scanErr, tmp
		}
		tmp[brosPerson.ClientCd+brosPerson.BrosPersonCd] = brosPerson
	}
	return http.StatusOK, nil, tmp
}
