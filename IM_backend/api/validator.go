package api

import (
	"github.com/Awadabang/Quasar-IM/util"
	"github.com/go-playground/validator/v10"
)

var validPassword validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if password, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedPassword(password)
	}
	return false
}
