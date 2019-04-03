package services

import (
	"MyWeb/enums"
	"MyWeb/errors"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate
func Validate(request interface{}) {
	validate = validator.New()
	if err := validate.Struct(request); err!=nil{
		panic(errors.ValidationError{enums.InvalidArguments,err.Error()})
	}
}
