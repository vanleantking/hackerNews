package params

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ListItemParams struct {
	Method string            `json:"method" validate:"required,oneof=GET POST PUT DELETE"`
	Params map[string]string `json:"params" validate:"required"`
}

func ValidatorListItemsRequest(listItemsRequest ListItemParams) []error {
	validate := validator.New()
	var errs []error
	err := validate.Struct(listItemsRequest)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			errs = append(errs, err)
			return errs
		}

		for _, err := range err.(validator.ValidationErrors) {
			er := errors.New(fmt.Sprintf("Error on  validate field %s, %s", err.StructField(), err.Error()))
			errs = append(errs, er)
		}

		// from here you can create your own error messages in whatever language you wish
		return errs
	}
	return nil
}
