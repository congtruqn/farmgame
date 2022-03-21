package infrastructure

import (
	"brostools-api-person/domain/model"
	"brostools-api-person/domain/repository"
	"brostools-api-person/lib/current"
	"brostools-api-person/lib/log"
	"database/sql"
	"fmt"
	"io/ioutil"
)

type MstClientIdRepository struct {
	Conn *ApiPersonBrantectInfrastructure
}

func NewMstClientIdRepository(conn *ApiPersonBrantectInfrastructure) repository.MstClientIdRepository {
	return &MstClientIdRepository{Conn: conn}
}

func (r *MstClientIdRepository) FindByClientCd(client_cd string) (*model.MstClientId, error) {
	sqlSelectTMNoSeqNo, errLoadFile := ioutil.ReadFile(current.GetCurrentDir() + "/infrastructure/sql/apiperson/brantect/getMstClientIdByClientCd.sql")
	if errLoadFile != nil {
		log.Infof(nil, fmt.Sprintf("FindByClientCd ioutil.ReadFile err: %v", errLoadFile))
		return nil, errLoadFile
	}

	sql_str := string(sqlSelectTMNoSeqNo)
	log.Infof(nil, fmt.Sprintf("FindByClientCd sql: %s", sql_str))

	mst_client_id := &model.MstClientId{}

	log.Infof(nil, fmt.Sprintf("FindByClientCd r.Conn %v", r.Conn))

	row := r.Conn.Database.QueryRow(sql_str, client_cd)
	if scan_err := row.Scan(
		&mst_client_id.ClientId,
		&mst_client_id.ClientCd,
		&mst_client_id.AuditFlg,
		&mst_client_id.Pwd,
		&mst_client_id.NrFlg,
		&mst_client_id.DbIp,
		&mst_client_id.FileDirPath,
		&mst_client_id.DbIpGlobal,
		&mst_client_id.DeFlg,
		&mst_client_id.DeAutoNoFlg,
		&mst_client_id.DeCsvDlFlg,
		&mst_client_id.DeCondSaveFlg,
		&mst_client_id.DeDataSetFlg,
		&mst_client_id.TmAutoNoFlg,
		&mst_client_id.TmCsvDlFlg,
		&mst_client_id.TmCondSaveFlg,
		&mst_client_id.TmDataSetFlg,
		&mst_client_id.TmAreaSearchFlg,
		&mst_client_id.TmDateCstFlg,
		&mst_client_id.TmImgAtFlg,
		&mst_client_id.TmTaskActionFlg,
		&mst_client_id.TmDispMstFlg,
		&mst_client_id.TmMadridProtSetFlg,
		&mst_client_id.LmFlg,
		&mst_client_id.LmApiStatus,
		&mst_client_id.LmApiIp,
		&mst_client_id.MoUrlNumFlg,
		&mst_client_id.MoHisSavePeriodFlg,
		&mst_client_id.MoAutoGettingFlg,
		&mst_client_id.TimeStampUseFlg,
		&mst_client_id.TimeStampLicenceFlg,
		&mst_client_id.TmElAppFlg,
		&mst_client_id.TmIpdlSyncFlg,
		&mst_client_id.DeHagueSetFlg,
		&mst_client_id.DnDateCstFlg,
		&mst_client_id.DbName,
		&mst_client_id.JsonIpArrow,
		&mst_client_id.IpArrowFlg,
		&mst_client_id.LastPasswordChangeDays,
		&mst_client_id.LastPasswordChangeDaysFlg,
		&mst_client_id.LastPasswordChangeUnit,
		&mst_client_id.SignUpFlg,
		&mst_client_id.SignUpAuthKey,
		&mst_client_id.JsonDomainMail,
		&mst_client_id.NmMngUseFlag,
		&mst_client_id.TmMngUseFlag,
		&mst_client_id.WmMngUseFlag,
		&mst_client_id.DbInstance,
	); scan_err != nil {

		log.Infof(nil, fmt.Sprintf("FindByClientCd QueryRow scan_err: %v", scan_err))
		if scan_err == sql.ErrNoRows {

		} else {
		}
		return nil, scan_err
	}

	return mst_client_id, nil
}
func (r *MstClientIdRepository) FindAllClient() ([]*model.MstClientId, error) {
	sqlSelect, errLoadFile := ioutil.ReadFile(current.GetCurrentDir() + "/infrastructure/sql/apiperson/brantect/getAllMstClientId.sql")
	if errLoadFile != nil {
		log.Infof(nil, fmt.Sprintf("FindAllClient ioutil.ReadFile err: %v", errLoadFile))
		return nil, errLoadFile
	}

	sql_str := string(sqlSelect)
	log.Infof(nil, fmt.Sprintf("FindAllClient sql: %s", sql_str))

	mst_client := []*model.MstClientId{}

	stmt, errPrepare := r.Conn.Database.Prepare(string(sqlSelect))
	if errPrepare != nil {
		return mst_client, errPrepare
	}
	rows, queryErr := stmt.Query()
	switch {
	case queryErr == sql.ErrNoRows:
		return mst_client, nil
	case queryErr != nil:
		return mst_client, nil
	}

	log.Infof(nil, fmt.Sprintf("FindAllClient r.Conn %v", r.Conn))

	for rows.Next() {
		mst_client_id := &model.MstClientId{}
		if scanErr := rows.Scan(
			&mst_client_id.ClientId,
			&mst_client_id.ClientCd,
			&mst_client_id.AuditFlg,
			&mst_client_id.Pwd,
			&mst_client_id.NrFlg,
			&mst_client_id.DbIp,
			&mst_client_id.FileDirPath,
			&mst_client_id.DbIpGlobal,
			&mst_client_id.DeFlg,
			&mst_client_id.DeAutoNoFlg,
			&mst_client_id.DeCsvDlFlg,
			&mst_client_id.DeCondSaveFlg,
			&mst_client_id.DeDataSetFlg,
			&mst_client_id.TmAutoNoFlg,
			&mst_client_id.TmCsvDlFlg,
			&mst_client_id.TmCondSaveFlg,
			&mst_client_id.TmDataSetFlg,
			&mst_client_id.TmAreaSearchFlg,
			&mst_client_id.TmDateCstFlg,
			&mst_client_id.TmImgAtFlg,
			&mst_client_id.TmTaskActionFlg,
			&mst_client_id.TmDispMstFlg,
			&mst_client_id.TmMadridProtSetFlg,
			&mst_client_id.LmFlg,
			&mst_client_id.LmApiStatus,
			&mst_client_id.LmApiIp,
			&mst_client_id.MoUrlNumFlg,
			&mst_client_id.MoHisSavePeriodFlg,
			&mst_client_id.MoAutoGettingFlg,
			&mst_client_id.TimeStampUseFlg,
			&mst_client_id.TimeStampLicenceFlg,
			&mst_client_id.TmElAppFlg,
			&mst_client_id.TmIpdlSyncFlg,
			&mst_client_id.DeHagueSetFlg,
			&mst_client_id.DnDateCstFlg,
			&mst_client_id.DbName,
			&mst_client_id.JsonIpArrow,
			&mst_client_id.IpArrowFlg,
			&mst_client_id.LastPasswordChangeDays,
			&mst_client_id.LastPasswordChangeDaysFlg,
			&mst_client_id.LastPasswordChangeUnit,
			&mst_client_id.SignUpFlg,
			&mst_client_id.SignUpAuthKey,
			&mst_client_id.JsonDomainMail,
			&mst_client_id.NmMngUseFlag,
			&mst_client_id.TmMngUseFlag,
			&mst_client_id.WmMngUseFlag,
			&mst_client_id.DbInstance,
		); scanErr != nil {
			return mst_client, scanErr
		}
		mst_client = append(mst_client, mst_client_id)
	}
	return mst_client, nil
}
