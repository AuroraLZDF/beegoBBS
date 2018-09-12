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

func Required(str string) error {
	if err := _validate(str, "required"); err != nil {
		return err
	}

	return nil
}

func IsEmail(email string) error {
	if err := _validate(email, "required,email"); err != nil {
		return err
	}

	return nil
}

func Equal(str1 string, str2 string) error {
	if err := Required(str1); err != nil {
		return err
	}

	if err := Required(str2); err != nil {
		return err
	}

	if str1 != str2 {
		err := Error("两次输入不一致")
		return err
	}

	return nil
}

func URL(url string) error {
	if err := _validate(url, "required,url"); err != nil {
		return err
	}

	return nil
}

func Numeric(num int) error {
	errs := validate.Var(num, "numeric")
	if errs != nil {
		return errs
	}

	return nil
}
