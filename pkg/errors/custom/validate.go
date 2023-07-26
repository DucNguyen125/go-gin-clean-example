package custom

import (
	"bytes"
	"context"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var errorLogicInvalidParam = NewLogicError(http.StatusBadRequest, "invalid parameter")

type FieldError struct {
	Code  string `json:"code"`
	Field string `json:"field"`
}

type ValidateError struct {
	HTTPCode    int          `json:"-"`
	FieldErrors []FieldError `json:"fieldErrors,omitempty"`
	Code        string       `json:"code,omitempty"`
}

func (ve *ValidateError) Error() string {
	return ve.Code
}

func (ve *ValidateError) GetHTTPCode() int {
	return ve.HTTPCode
}

func NewValidateError(ctx context.Context, input interface{}, err error) CustomError {
	validateError := ValidateError{
		HTTPCode:    http.StatusBadRequest,
		FieldErrors: make([]FieldError, 0),
	}
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return errorLogicInvalidParam
	}
	for _, validationError := range validationErrors {
		field := lowerCaseFieldName(validationError.Namespace())
		validateError.FieldErrors = append(validateError.FieldErrors, FieldError{
			Code:  validationError.ActualTag(),
			Field: field,
		})
	}
	return &validateError
}

func lowerCaseFieldName(field string) string {
	arrayString := strings.Split(field, ".")
	var newArrayString []string
	for i, v := range arrayString {
		if i == 0 {
			continue
		}
		bts := []byte(v)
		lc := bytes.ToLower([]byte{bts[0]})
		rest := bts[1:]
		newArrayString = append(newArrayString, string(bytes.Join([][]byte{lc, rest}, nil)))
	}
	return strings.Join(newArrayString, ".")
}
