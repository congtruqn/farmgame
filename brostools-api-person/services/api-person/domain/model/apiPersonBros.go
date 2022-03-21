package model

import (
	"encoding/json"
)

type ApiPerson struct {
	ClientCd     string `json:"client_cd"`
	BrosPersonCd string `json:"person_cd"`
	PersonField  string `json:"person_field"`
	DivisionNm   string `json:"division_nm"`
	PositionNm   string `json:"position_nm"`
	PersonNm     string `json:"person_nm"`
	PersonNmJp   string `json:"person_nm_jp"`
	PostCd       string `json:"post_cd"`
	Address      string `json:"address"`
	Tel          string `json:"tel"`
	Fax          string `json:"fax"`
	Email        string `json:"email"`
	Mobile       string `json:"moby"`
	BrosRemarks  string `json:"remarks"`
	InvsndDnFlg  string `json:"invsnd_dn_flg"`
	InvsndTmFlg  string `json:"invsnd_tm_flg"`
	InvsndOtFlg  string `json:"invsnd_ot_flg"`
	SignFlg      string `json:"sign_flg"`
	DeleteFlg    string `json:"delete_flg"`
	UpdDate      string `json:"upd_date"`
	UpdUser      string `json:"upd_user"`
	UpdPrgId     string `json:"upd_prg_id"`
	InpDate      string `json:"inp_date"`
	InpUser      string `json:"inp_user"`
	InpPrgId     string `json:"inp_prg_id "`
}
type BrosMapSet map[string]ApiPerson

func (item *ApiPerson) ToJSON() string {
	json, _ := json.Marshal(item)
	return string(json)
}
