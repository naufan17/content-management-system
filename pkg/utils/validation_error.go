package utils

import "github.com/go-playground/validator/v10"

func ParseValidationError(err validator.ValidationErrors) map[string]string {
	errorMessage := make(map[string]string)

	for _, v := range err {
		switch v.StructField() {
		case "Name":
			if v.Tag() == "required" {
				errorMessage[v.Field()] = "name is required"
			} else if v.Tag() == "max" {
				errorMessage[v.Field()] = "name must be at most 50 characters"
			}
		case "Username":
			if v.Tag() == "required" {
				errorMessage[v.Field()] = "username is required"
			} else if v.Tag() == "max" {
				errorMessage[v.Field()] = "username must be at most 50 characters"
			}
		case "Password":
			if v.Tag() == "required" {
				errorMessage[v.Field()] = "password is required"
			} else if v.Tag() == "min" {
				errorMessage[v.Field()] = "password must be at least 10 characters"
			}
		}
	}

	return errorMessage
}
