package model

import "encoding/json"

type ApiPersonBrantect struct {
	ClientCd         string `json:"client_cd"`
	DeptCd           string `json:"dept_cd"`
	BrantectPersonCd string `json:"person_cd"`
	PositionNm       string `json:"executive"`
	PersonNm         string `json:"person_nm"`
	PersonNmJp       string `json:"person_name_kana"`
	PostCd           string `json:"zip"`
	Address          string `json:"address"`
	Tel              string `json:"tel"`
	Fax              string `json:"fax"`
	Email            string `json:"email"`
	Mobile           string `json:"mobile"`
	BrantectRemarks  string `json:"remarks"`
	SearchPersonNm   string `json:"search_person_nm"`
	Type             string `json:"type"`
	Idno             string `json:"idno"`
	LastVerifiedDate string `json:"last_verified_date"`
	DeleteFlg        string `json:"delete_flg"`
	UpdDate          string `json:"upd_date"`
	UpdUser          string `json:"upd_user"`
	InpDate          string `json:"inp_date"`
	InpUser          string `json:"inp_user"`
}

type TCPSet map[string]ApiPersonBrantect

func (item *ApiPersonBrantect) ToJSON() string {
	json, _ := json.Marshal(item)
	return string(json)
}
