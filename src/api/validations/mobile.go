package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/miladshalikar/golang-clean-web-api/src/common"
)

func IranianMobileNumberValidator(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	return common.IranianMobileNumberValidate(value)
}
