package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate *validator.Validate

type validationError struct {
	Namespace       string `json:"namespace"` // can differ when a custom TagNameFunc is registered or
	Field           string `json:"field"`     // by passing alt name to ReportError like below
	StructNamespace string `json:"structNamespace"`
	StructField     string `json:"structField"`
	Tag             string `json:"tag"`
	ActualTag       string `json:"actualTag"`
	Kind            string `json:"kind"`
	Type            string `json:"type"`
	Value           string `json:"value"`
	Param           string `json:"param"`
	Message         string `json:"message"`
}

func Init() {
	validate = validator.New()
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStruct(i interface{}) (bool, *[]FieldError) {
	err := validate.Struct(i)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return false, nil
		}
		var ef []FieldError
		for _, err := range err.(validator.ValidationErrors) {
			//fmt.Printf("Error on validation. %s | %s | %s | %v | %v | %s | %s | %s | %s | %s | %s\n", err.Field(),
			//	err.Error(), err.Param(), err.Kind(), err.Type(),
			//	err.StructNamespace(), err.StructField(),
			//	err.Tag(), err.ActualTag(), err.Value(), err.Namespace(),
			//)
			var msg string
			if strings.TrimSpace(err.Param()) != "" {
				msg = fmt.Sprintf("Input failed on tag `%s` with param `%s`", err.Tag(), err.Param())
			} else {
				msg = fmt.Sprintf("Input failed on condition `%s`", err.Tag())
			}
			ef = append(ef, FieldError{
				Field:   err.Field(),
				Message: msg,
			})
		}
		return false, &ef
	}
	return true, nil
}

func ValidateVariable(i interface{}, rule string) (bool, *[]FieldError) {
	err := validate.Var(i, rule)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return false, nil
		}
		var ef []FieldError
		for _, err := range err.(validator.ValidationErrors) {
			var msg string
			if strings.TrimSpace(err.Param()) != "" {
				msg = fmt.Sprintf("Input failed on tag `%s` with param `%s`", err.Tag(), err.Param())
			} else {
				msg = fmt.Sprintf("Input failed on condition `%s`", err.Tag())
			}
			ef = append(ef, FieldError{
				Field:   err.Field(),
				Message: msg,
			})
		}
		return false, &ef
	}
	return true, nil
}
