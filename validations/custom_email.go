package validations

import (
	"github.com/dlclark/regexp2"
	"github.com/go-playground/validator/v10"
)

var CustomEmail validator.Func = func(fl validator.FieldLevel) bool {
	matched, _ := regexp2.MustCompile("[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}$", 0).
		MatchString(fl.Field().String())
	return matched
}
