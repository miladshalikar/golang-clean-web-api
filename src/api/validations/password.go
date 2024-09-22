package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/miladshalikar/golang-clean-web-api/src/common"
)

func PasswordValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		fld.Param()
		return false
	}

	return common.CheckPassword(value)
}
