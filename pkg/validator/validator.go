package validator

import (
	"hackerNewsApi/pkg/config"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validate *validator.Validate
}

func NewValidator(config *config.Config) *Validator {
	return &Validator{Validate: validator.New()}
}
