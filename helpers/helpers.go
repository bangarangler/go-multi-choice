package helpers

import (
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"strings"
)

func ValidateInputs(dataSet interface{}) (bool, string) {

	var validate *validator.Validate

	validate = validator.New()

	err := validate.Struct(dataSet)

	if err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		var errString string

		reflected := reflect.ValueOf(dataSet)

		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflected.Type().FieldByName(err.StructField())
			var name string
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				errString = "The " + name + " is required"
				break
			case "email":
				errString = "The " + name + " should be a valid email"
				break
			case "eqfield":
				errString = "The " + name + " should be equal to the " + err.Param()
				break
			default:
				errString = "The " + name + " is invalid"
				break
			}
		}
		return false, errString
	}
	return true, ""
}
