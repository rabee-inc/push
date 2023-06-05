package handler

import "gopkg.in/go-playground/validator.v9"

var v = validator.New()

func Validate(s interface{}) error {
	return v.Struct(s)
}
