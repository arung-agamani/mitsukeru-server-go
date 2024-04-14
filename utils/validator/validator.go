package validator

import (
	"errors"
	"github.com/arung-agamani/mitsukeru-go/utils/logger"
	"github.com/go-playground/validator/v10"
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

func ValidateStruct(i interface{}) (bool, string) {
	err := validate.Struct(i)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return false, "Malformed input"
		}

		for _, err := range err.(validator.ValidationErrors) {
			logger.Errorf("Error on validation. %s : %s | %s | %s | %s", err.Field(), err.Error(), err.Param(), err.Kind(), err.Type())
		}
		return false, "Field validation error on [coming soon]"
	}
	return true, ""
}
