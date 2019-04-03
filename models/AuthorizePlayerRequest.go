package models

type AuthorizePlayerRequest struct {
	Ipaddress string `json:"ipaddress" validate:"required"`
	Username string `json:"username"`
	Userid string `json:"userid"`
	Lang string `json:"lang"`
	Cur string `json:"cur"`
	Betlimitid int `json:"betlimitid"`
	Platformtype int `json:"platformtype"`
	Istestplayer bool `json:"istestplayer"`
	Playertype int `json:"playertype"`
}
