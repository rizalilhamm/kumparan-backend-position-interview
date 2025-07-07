package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ValidationUtil struct {
	validator *validator.Validate
}

func NewValidationUtil() echo.Validator {
	return &ValidationUtil{validator: validator.New()}
}

func (v *ValidationUtil) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func BindValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	if err := c.Validate(i); err != nil {
		return err
	}
	return nil
}

func ValidateImageSizeUnder1MB(c echo.Context, fieldName string) error {
	// Get the uploaded file from the request
	file, err := c.FormFile(fieldName)
	if err != nil {
		return err
	}
	// Get the size of the file in bytes
	size := file.Size

	// Check if the file is less than 1 MB in size
	if size > 1000000 {
		return errors.New("File size exceeds 1 MB limit")
	}

	return nil
}

func GenerateErrorMessage(field string, tag string, param string) string {
	switch {
	case tag == "required":
		return fmt.Sprintf("The %s field is required", field)
	case tag == "gt":
		return fmt.Sprintf("The %s field must be greater than %s characters", field, param)
	case tag == "gte":
		return fmt.Sprintf("The %s field must be greater than or equal to %s characters", field, param)
	case tag == "lt":
		return fmt.Sprintf("The %s field must be less than %s characters", field, param)
	case tag == "lte":
		return fmt.Sprintf("The %s field must be less than or equal to %s characters", field, param)
	case tag == "len":
		return fmt.Sprintf("The %s field must be equal to %s characters", field, param)
	case tag == "min":
		return fmt.Sprintf("The %s field must have a minimum %s characters", field, param)
	case tag == "max":
		return fmt.Sprintf("The %s field must have a maximum %s characters", field, param)
	case tag == "eq":
		return fmt.Sprintf("The %s field must be equal to %s", field, param)
	case strings.HasPrefix(tag, "eq="):
		values := strings.Split(tag, "|")
		equalValues := make([]string, len(values))
		for i, value := range values {
			equalValues[i] = strings.TrimPrefix(value, "eq=")
		}
		return fmt.Sprintf("The %s field must be equal to %s", field, strings.Join(equalValues, " or "))
	}
	return "Error when generating error message"
}
