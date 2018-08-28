package utils

import (
	"gopkg.in/go-playground/validator.v9"
)

//var validate *validator.Validate
var validate = validator.New()

func _validate(str string, validator string) error {
	errs := validate.Var(str, validator)

	if errs != nil {
		return errs
	}
	return nil
}

func Required(str string) (bool, error) {
	if err := _validate(str, "required"); err != nil {
		return false, err
	}

	return true, nil
}

func IsEmail(email string) (bool, error) {
	if err := _validate(email, "required,email"); err != nil {
		return false, err
	}

	return true, nil
}

func Equal(str1 string, str2 string) (bool, string) {

	if str1 != str2 {
		return false, "两次输入不一致"
	}

	return true, ""
}
