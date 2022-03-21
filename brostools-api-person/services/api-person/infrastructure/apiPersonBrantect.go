package infrastructure

import (
	"brostools-api-person/domain/model"
	"brostools-api-person/domain/repository"
	"brostools-api-person/lib/current"
	"brostools-api-person/lib/log"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	_ "github.com/lib/pq"
)

type ApiPersonBrantectInfrastructure struct {
	User        string
	Password    string
	Connection  string
	Database    *sql.DB
	Transaction *sql.Tx
	Driver      string
	ClientCd    string
}

func NewApiPersonBrantectfrastructure(driver string) repository.ApiPersonBrantectRepository {
	return &ApiPersonBrantectInfrastructure{
		Driver: driver,
	}
}
func (db *ApiPersonBrantectInfrastructure) Connect(client_cd string) error {
	switch db.Driver {
	/*case "mysql":
	  db.ConnectToMySql()*/
	case "pgsql":
		err := db.ConnectToPgSql(client_cd)
		if err != nil {
			return err
		}
	default:

		log.Infof(nil, fmt.Sprintf("%s Driver is not support.", db.Driver))
	}
	return nil
}
func (db *ApiPersonBrantectInfrastructure) ConnectToPgSql(client_cd string) error {

	log.Infof(nil, "ConnectToPgSql")

	pgsql_instance := os.Getenv("POSTGRES_DSN_BRANTECT_API_PERSON")
	pgsql_pass := os.Getenv("POSTGRES_PASSWORD_BRANTECT_API_PERSON")
	pgsql_user := os.Getenv("POSTGRES_USER_BRANTECT_API_PERSON")
	pgsql_dbname := os.Getenv("POSTGRES_DBNAME_BRANTECT_API_PERSON")
	pgsql_dsn := ""
	runtime_env := strings.ToLower(os.Getenv("RUNTIME_ENV"))
	log.Infof(nil, fmt.Sprintf("ConnectToPgSql RUNTIME_ENV: %s", runtime_env))

	if runtime_env != "gae" && runtime_env != "gcr" { // local
		pgsql_dsn = fmt.Sprintf("host=%s dbname=%s", pgsql_instance, pgsql_dbname)
	} else {
		pgsql_dsn = pgsql_instance
	}

	dsn := fmt.Sprintf("%s port=5432 user=%s password=%s sslmode=disable", pgsql_dsn, pgsql_user, pgsql_pass)
	log.Infof(nil, fmt.Sprintf("ConnectToPgSql PgSql DSN: %s", dsn))

	pgsql_db, open_err := sql.Open("postgres", dsn)
	if open_err != nil {
		log.Infof(nil, fmt.Sprintf("ConnectToPgSql connect to PgSql open_err: %v", open_err))
		return open_err
	}

	ping_err := pgsql_db.Ping()
	if ping_err != nil {
		log.Infof(nil, fmt.Sprintf("ConnectToPgSql connect to PgSql ping_err: %v", ping_err))
		return ping_err
	}

	// set connection to general pgsql_db
	log.Infof(nil, "ConnectToPgSql set connection to general pgsql_db")
	db.Database = pgsql_db

	if client_cd != "" {

		log.Infof(nil, fmt.Sprintf("ConnectToPgSql connect to Private DB of client: %s", client_cd))

		repo := NewMstClientIdRepository(db)
		mst_client_id, find_err := repo.FindByClientCd(client_cd)

		if find_err != nil {
			log.Infof(nil, fmt.Sprintf("ConnectToPgSql find client id info find_err: %v", find_err))
			return nil
		}

		log.Infof(nil, "ConnectToPgSql close connection to general pgsql_db")
		pgsql_db.Close()

		log.Infof(nil, "ConnectToPgSql unset connection to general pgsql_db")
		db.Database = nil

		log.Infof(nil, fmt.Sprintf("ConnectToPgSql mst_client_id: %v", mst_client_id))

		db_ip := ""
		matched, _ := regexp.MatchString("/", mst_client_id.DbIp)
		if matched {
			arrSplit := strings.Split(mst_client_id.DbIp, "/")
			db_ip = arrSplit[0]
		} else {
			db_ip = mst_client_id.DbIp
		}
		db_instance := mst_client_id.DbInstance
		if db_instance == "" {
			db_instance = os.Getenv("POSTGRES_DSN_BRANTECT_API_PERSON")
		}

		api_run_type := strings.ToLower(os.Getenv("API_RUN_TYPE"))
		log.Infof(nil, fmt.Sprintf("ConnectToPgSql API_RUN_TYPE: %s", api_run_type))

		switch api_run_type {
		case "stg":
			pgsql_dbname = db_ip
			if pgsql_dbname == "" {
				pgsql_dbname = os.Getenv("POSTGRES_DBNAME_BRANTECT_API_PERSON")
			}
			pgsql_host := strings.Replace(string(pgsql_instance), "dbname=brantect", "", 1)
			pgsql_dsn = fmt.Sprintf("%s dbname=%s", pgsql_host, pgsql_dbname)
		case "dev", "prd":
			pgsql_host := ""
			if runtime_env == "gae" || runtime_env == "gcr" {
				pgsql_host = fmt.Sprintf("/cloudsql/%s", strings.Replace(string(strings.Replace(string(db_instance), "host=/cloudsql/", "", 1)), "/cloudsql/", "", 1))
			} else { // local
				pgsql_host = db_ip // runtime_env == local
				if pgsql_host == "" {
					pgsql_host = os.Getenv("POSTGRES_DSN_BRANTECT_API_PERSON")
				}
			}
			pgsql_dsn = fmt.Sprintf("host=%s dbname=%s", pgsql_host, pgsql_dbname)
		default:
		}

		dsn := fmt.Sprintf("%s port=5432 user=%s password=%s sslmode=disable", pgsql_dsn, pgsql_user, pgsql_pass)
		log.Infof(nil, fmt.Sprintf("ConnectToPgSql Private dsn: %s", dsn))

		// open private db
		pgsql_db, open_err = sql.Open("postgres", dsn)
		if open_err != nil {
			log.Infof(nil, fmt.Sprintf("ConnectToPgSql connect to private PgSql open_err: %v", open_err))

			return open_err
		}

		log.Infof(nil, "ConnectToPgSql set connection to private pgsql_db")
		db.Database = pgsql_db
	}

	return nil
}
func (infa *ApiPersonBrantectInfrastructure) ConnectToGeneralPgSql() error {
	pgsql_instance := os.Getenv("POSTGRES_DSN_BRANTECT_API_PERSON")
	pgsql_pass := os.Getenv("POSTGRES_PASSWORD_BRANTECT_API_PERSON")
	pgsql_user := os.Getenv("POSTGRES_USER_BRANTECT_API_PERSON")
	pgsql_dbname := os.Getenv("POSTGRES_DBNAME_BRANTECT_API_PERSON")
	pgsql_dsn := ""
	runtime_env := strings.ToLower(os.Getenv("RUNTIME_ENV"))
	if runtime_env != "gae" && runtime_env != "gcr" { // local
		pgsql_dsn = fmt.Sprintf("host=%s dbname=%s", pgsql_instance, pgsql_dbname)
	} else {
		pgsql_dsn = pgsql_instance
	}
	dsn := fmt.Sprintf("%s port=5432 user=%s password=%s sslmode=disable", pgsql_dsn, pgsql_user, pgsql_pass)

	pgsql_db, open_err := sql.Open("postgres", dsn)
	if open_err != nil {
		log.Infof(nil, fmt.Sprintf("ConnectToPgSql connect to PgSql open_err: %v", open_err))
		return open_err
	}

	ping_err := pgsql_db.Ping()
	if ping_err != nil {
		log.Infof(nil, fmt.Sprintf("ConnectToPgSql connect to PgSql ping_err: %v", ping_err))
		return ping_err
	}

	// set connection to general pgsql_db
	log.Infof(nil, "ConnectToPgSql set connection to general pgsql_db")
	infa.Database = pgsql_db
	return nil
}
func (infa *ApiPersonBrantectInfrastructure) ConnectToPrivatePgSql(client_id *model.MstClientId) error {
	pgsql_instance := os.Getenv("POSTGRES_DSN_BRANTECT_API_PERSON")
	pgsql_pass := os.Getenv("POSTGRES_PASSWORD_BRANTECT_API_PERSON")
	pgsql_user := os.Getenv("POSTGRES_USER_BRANTECT_API_PERSON")
	pgsql_dbname := os.Getenv("POSTGRES_DBNAME_BRANTECT_API_PERSON")
	pgsql_dsn := ""
	runtime_env := strings.ToLower(os.Getenv("RUNTIME_ENV"))
	if runtime_env != "gae" && runtime_env != "gcr" { // local
		pgsql_dsn = fmt.Sprintf("host=%s dbname=%s", pgsql_instance, pgsql_dbname)
	} else {
		pgsql_dsn = pgsql_instance
	}
	infa.Database = nil
	db_ip := ""
	matched, _ := regexp.MatchString("/", client_id.DbIp)
	if matched {
		arrSplit := strings.Split(client_id.DbIp, "/")
		db_ip = arrSplit[0]
	} else {
		db_ip = client_id.DbIp
	}
	db_instance := client_id.DbInstance
	if db_instance == "" {
		db_instance = os.Getenv("POSTGRES_DSN_BRANTECT_API_PERSON")
	}

	api_run_type := strings.ToLower(os.Getenv("API_RUN_TYPE"))
	log.Infof(nil, fmt.Sprintf("ConnectToPgSql API_RUN_TYPE: %s", api_run_type))

	switch api_run_type {
	case "stg":
		pgsql_dbname = db_ip
		if pgsql_dbname == "" {
			pgsql_dbname = os.Getenv("POSTGRES_DBNAME_BRANTECT_API_PERSON")
		}
		pgsql_host := strings.Replace(string(pgsql_instance), "dbname=brantect", "", 1)
		pgsql_dsn = fmt.Sprintf("%s dbname=%s", pgsql_host, pgsql_dbname)
	case "dev", "prd":
		pgsql_host := ""
		if runtime_env == "gae" || runtime_env == "gcr" {
			pgsql_host = fmt.Sprintf("/cloudsql/%s", strings.Replace(string(strings.Replace(string(db_instance), "host=/cloudsql/", "", 1)), "/cloudsql/", "", 1))
		} else { // local
			pgsql_host = db_ip // runtime_env == local
			if pgsql_host == "" {
				pgsql_host = os.Getenv("POSTGRES_DSN_BRANTECT_API_PERSON")
			}
		}
		pgsql_dsn = fmt.Sprintf("host=%s dbname=%s", pgsql_host, pgsql_dbname)
	default:
	}

	dsn := fmt.Sprintf("%s port=5432 user=%s password=%s sslmode=disable", pgsql_dsn, pgsql_user, pgsql_pass)
	log.Infof(nil, fmt.Sprintf("ConnectToPgSql Private dsn: %s", dsn))

	// open private db
	pgsql_db, open_err := sql.Open("postgres", dsn)
	if open_err != nil {
		log.Infof(nil, fmt.Sprintf("ConnectToPgSql connect to private PgSql open_err: %v", open_err))
		return open_err
	}
	log.Infof(nil, "ConnectToPgSql set connection to private pgsql_db")
	infa.Database = pgsql_db
	return nil
}
func (infa *ApiPersonBrantectInfrastructure) Close() {
	infa.Database.Close()
}

func (infa *ApiPersonBrantectInfrastructure) Begin() error {
	tx, errBegin := infa.Database.Begin()
	if errBegin != nil {
		return errBegin
	}
	infa.Transaction = tx
	return nil
}

func (infa *ApiPersonBrantectInfrastructure) Rollback() {
	infa.Transaction.Rollback()
}
func (infa *ApiPersonBrantectInfrastructure) Commit() error {
	return infa.Transaction.Commit()
}
func (repo *ApiPersonBrantectInfrastructure) AddBrantect(ri *model.ApiPersonBrantect) (int, error) {
	fmt.Println("Bantect Api Person Repository Add")

	sqlSelectTMNoSeqNo, errLoadFile := ioutil.ReadFile(current.GetCurrentDir() + "/infrastructure/sql/apiperson/brantect/AddApiPersonBrantect.sql")
	defer repo.Close()
	if errLoadFile != nil {
		return http.StatusInternalServerError, errLoadFile
	}

	stmt, errPrepare := repo.Transaction.Prepare(string(sqlSelectTMNoSeqNo))
	if errPrepare != nil {
		return http.StatusInternalServerError, errPrepare
	}

	_, errExec := stmt.Exec(
		&ri.ClientCd,
		&ri.DeptCd,
		&ri.BrantectPersonCd,
		&ri.PositionNm,
		&ri.PersonNm,
		&ri.PersonNmJp,
		&ri.PostCd,
		&ri.Address,
		&ri.Tel,
		&ri.Fax,
		&ri.Email,
		&ri.Mobile,
		&ri.BrantectRemarks,
		&ri.SearchPersonNm,
		&ri.Type,
		&ri.Idno,
		"/brostools/api/person/",
		"0",
	)
	if errExec != nil {
		return http.StatusInternalServerError, errExec
	}
	return http.StatusOK, nil
}
func (repo *ApiPersonBrantectInfrastructure) BrantectGetByID(client_cd string, person_cd string) (int, error, *model.ApiPersonBrantect) {

	fmt.Println("Bantect Api Person Get By ID")
	var branTectPerson model.ApiPersonBrantect

	errOpenBr := repo.Connect(client_cd)
	if errOpenBr != nil {
		log.Infof(nil, "Error Connect Bantect")
		return http.StatusInternalServerError, errOpenBr, &branTectPerson
	}
	sqlSelect, errLoadFile := ioutil.ReadFile(current.GetCurrentDir() + "/infrastructure/sql/apiperson/brantect/ApiPersonBrantectGetByID.sql")
	defer repo.Close()
	if errLoadFile != nil {
		return http.StatusInternalServerError, errLoadFile, &branTectPerson
	}

	stmt, errPrepare := repo.Database.Prepare(string(sqlSelect))
	if errPrepare != nil {
		return http.StatusInternalServerError, errPrepare, &branTectPerson
	}
	queryErr := stmt.QueryRow(client_cd, person_cd).Scan(
		&branTectPerson.ClientCd,
		&branTectPerson.DeptCd,
		&branTectPerson.BrantectPersonCd,
		&branTectPerson.PositionNm,
		&branTectPerson.PersonNm,
		&branTectPerson.PersonNmJp,
		&branTectPerson.PostCd,
		&branTectPerson.Address,
		&branTectPerson.Tel,
		&branTectPerson.Fax,
		&branTectPerson.Email,
		&branTectPerson.Mobile,
		&branTectPerson.BrantectRemarks,
		&branTectPerson.SearchPersonNm,
		&branTectPerson.Type,
		&branTectPerson.Idno,
		&branTectPerson.LastVerifiedDate,
		&branTectPerson.DeleteFlg,
		&branTectPerson.UpdDate,
		&branTectPerson.UpdUser,
		&branTectPerson.InpDate,
		&branTectPerson.InpUser,
	)
	switch {
	case queryErr == sql.ErrNoRows:
		return http.StatusNotFound, nil, &branTectPerson
	case queryErr != nil:
		return http.StatusInternalServerError, nil, &branTectPerson
	default:
		return http.StatusOK, nil, &branTectPerson
	}
}
func (repo *ApiPersonBrantectInfrastructure) BrantectUpdate(client_cd string, person_cd string, ri *model.ApiPersonBrantect) (int, error) {
	fmt.Println("Bantect Api Person Repository Update")

	sqlSelectTMNoSeqNo, errLoadFile := ioutil.ReadFile(current.GetCurrentDir() + "/infrastructure/sql/apiperson/brantect/UpdateApiPersonBrantect.sql")
	defer repo.Close()
	if errLoadFile != nil {
		return http.StatusInternalServerError, errLoadFile
	}

	stmt, errPrepare := repo.Transaction.Prepare(string(sqlSelectTMNoSeqNo))
	if errPrepare != nil {
		return http.StatusInternalServerError, errPrepare
	}
	_, errExec := stmt.Exec(
		&ri.DeptCd,
		&ri.BrantectPersonCd,
		&ri.PositionNm,
		&ri.PersonNm,
		&ri.PersonNmJp,
		&ri.PostCd,
		&ri.Address,
		&ri.Tel,
		&ri.Fax,
		&ri.Email,
		&ri.Mobile,
		&ri.BrantectRemarks,
		&ri.SearchPersonNm,
		&ri.Type,
		&ri.Idno,
		"/brostools/api/person/",
		client_cd,
		person_cd,
	)
	if errExec != nil {
		return http.StatusInternalServerError, errExec
	}
	return http.StatusOK, nil
}
func (repo *ApiPersonBrantectInfrastructure) BrantectDelete(client_cd string, person_cd string) (int, error) {
	sqlDelete, errLoadFile := ioutil.ReadFile(current.GetCurrentDir() + "/infrastructure/sql/apiperson/brantect/DeleteApiPersonBrantect.sql")
	defer repo.Close()
	if errLoadFile != nil {
		return http.StatusInternalServerError, errLoadFile
	}
	stmt, errPrepare := repo.Transaction.Prepare(string(sqlDelete))
	if errPrepare != nil {
		return http.StatusInternalServerError, errPrepare
	}
	_, errExec := stmt.Exec(
		"/brostools/api/person/",
		client_cd,
		person_cd,
	)
	if errExec != nil {
		return http.StatusInternalServerError, errExec
	}
	return http.StatusOK, nil
}
func (repo *ApiPersonBrantectInfrastructure) BrantectGetAll() (int, error, model.TCPSet) {
	tmp := make(model.TCPSet)
	fmt.Println("Bantect Api Person Get All")

	conect_error := repo.ConnectToGeneralPgSql()

	if conect_error != nil {
		log.Infof(nil, fmt.Sprintf("ConnectToPgSql connect to PgSql error: %v", conect_error))
		return http.StatusInternalServerError, conect_error, tmp
	}

	// Get data from General DB

	sqlSelect, errLoadFile := ioutil.ReadFile(current.GetCurrentDir() + "/infrastructure/sql/apiperson/brantect/ApiPersonBrantectGetAll.sql")
	if errLoadFile != nil {
		return http.StatusInternalServerError, errLoadFile, tmp
	}

	stmt, errPrepare := repo.Database.Prepare(string(sqlSelect))
	if errPrepare != nil {
		return http.StatusInternalServerError, errPrepare, tmp
	}
	rows, queryErr := stmt.Query()

	switch {
	case queryErr != nil:
		return http.StatusInternalServerError, nil, tmp
	}
	for rows.Next() {
		var branTectPerson model.ApiPersonBrantect
		if scanErr := rows.Scan(
			&branTectPerson.ClientCd,
			&branTectPerson.DeptCd,
			&branTectPerson.BrantectPersonCd,
			&branTectPerson.PositionNm,
			&branTectPerson.PersonNm,
			&branTectPerson.PersonNmJp,
			&branTectPerson.PostCd,
			&branTectPerson.Address,
			&branTectPerson.Tel,
			&branTectPerson.Fax,
			&branTectPerson.Email,
			&branTectPerson.Mobile,
			&branTectPerson.BrantectRemarks,
			&branTectPerson.SearchPersonNm,
			&branTectPerson.Type,
			&branTectPerson.Idno,
			&branTectPerson.LastVerifiedDate,
			&branTectPerson.DeleteFlg,
			&branTectPerson.UpdDate,
			&branTectPerson.UpdUser,
			&branTectPerson.InpDate,
			&branTectPerson.InpUser,
		); scanErr != nil {
			return http.StatusInternalServerError, scanErr, tmp
		}
		tmp[branTectPerson.ClientCd+branTectPerson.BrantectPersonCd] = branTectPerson
	}

	mstclientrepo := NewMstClientIdRepository(repo)
	mst_client, find_err := mstclientrepo.FindAllClient()
	defer repo.Close()
	if find_err != nil {
		log.Infof(nil, fmt.Sprintf("ConnectToPgSql connect to PgSql err: %v", find_err))
		return http.StatusInternalServerError, find_err, tmp
	}
	// Get data from Private DB

	for k := range mst_client {

		privateconect_error := repo.ConnectToPrivatePgSql(mst_client[k])
		if privateconect_error != nil {
			log.Infof(nil, fmt.Sprintf("ConnectToPgSql connect to PgSql err: %v", privateconect_error))
			return http.StatusInternalServerError, privateconect_error, tmp
		}
		sqlSelect, errLoadFile := ioutil.ReadFile(current.GetCurrentDir() + "/infrastructure/sql/apiperson/brantect/ApiPersonBrantectGetAll.sql")

		if errLoadFile != nil {
			return http.StatusInternalServerError, errLoadFile, tmp
		}

		stmt, errPrepare := repo.Database.Prepare(string(sqlSelect))
		if errPrepare != nil {
			return http.StatusInternalServerError, errPrepare, tmp
		}
		rows, queryErr := stmt.Query()

		switch {
		case queryErr != nil:
			return http.StatusInternalServerError, nil, tmp
		}
		for rows.Next() {
			var branTectPerson model.ApiPersonBrantect
			if scanErr := rows.Scan(
				&branTectPerson.ClientCd,
				&branTectPerson.DeptCd,
				&branTectPerson.BrantectPersonCd,
				&branTectPerson.PositionNm,
				&branTectPerson.PersonNm,
				&branTectPerson.PersonNmJp,
				&branTectPerson.PostCd,
				&branTectPerson.Address,
				&branTectPerson.Tel,
				&branTectPerson.Fax,
				&branTectPerson.Email,
				&branTectPerson.Mobile,
				&branTectPerson.BrantectRemarks,
				&branTectPerson.SearchPersonNm,
				&branTectPerson.Type,
				&branTectPerson.Idno,
				&branTectPerson.LastVerifiedDate,
				&branTectPerson.DeleteFlg,
				&branTectPerson.UpdDate,
				&branTectPerson.UpdUser,
				&branTectPerson.InpDate,
				&branTectPerson.InpUser,
			); scanErr != nil {
				return http.StatusInternalServerError, scanErr, tmp
			}
			tmp[branTectPerson.ClientCd+branTectPerson.BrantectPersonCd] = branTectPerson
		}
		repo.Close()
	}

	return http.StatusOK, nil, tmp
}
