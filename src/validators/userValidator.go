package validators

import (
	"github.com/go-playground/validator/v10"
)

type UserName string

func init() {
	registerValidation("username", UserName("required,min=4,max=8").toFunc())
}

// var VUserName validator.Func = func(f1 validator.FieldLevel) bool {
// 	username, ok := f1.Field().Interface().(string)
// 	if ok && len(username) >= 4 {
// 		return true
// 	}
// 	return false
// }

func (s UserName) toFunc() validator.Func {
	validatorError["username"] = "用户名为必填项，且必须在4-8位之间"
	return func(f1 validator.FieldLevel) bool {
		v, ok := f1.Field().Interface().(string)
		if ok {
			return s.validate(v)
		}
		return false
	}
}

func (s UserName) validate(v string) bool {
	if err := jstockValidator.Var(v, string(s)); err != nil {
		return false
	}
	return true
}
