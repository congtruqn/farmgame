package model

type BrosToken struct {
	Exp       int    `json:"exp"`
	Iat       int    `json:"iat"`
	Nbf       int    `json:"nbf"`
	SessionId string `json:"session_id"`
	UserId    string `json:"user_id"`
}
