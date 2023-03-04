package api

import (
	"github.com/git-adithyanair/cs130-group-project/util"
	"github.com/go-playground/validator/v10"
)

var validItemQuantityType validator.Func = func(fieldLevel validator.FieldLevel) bool {
	currency, ok := fieldLevel.Field().Interface().(string)
	if !ok {
		return false
	}
	return util.IsValidItemQuantityType(currency)
}
