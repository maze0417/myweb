package models

type AuthorizePlayerResponse struct {
	Err       int `json:"err,omitempty"`
	Errdesc   string `json:"errdesc,omitempty"`
	Isnew     bool `json:"isnew,omitempty"`
	Authtoken string `json:"authtoken,omitempty"`
}
