package models

type BaseResponse struct {
	Err       int `json:"err,omitempty"`
	Errdesc   string `json:"errdesc,omitempty"`
}
func CreateResponse(err int , errdesc string) BaseResponse{
	return BaseResponse{
		Err:err,
		Errdesc:errdesc,
	}
}