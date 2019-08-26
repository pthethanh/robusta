package validator

import (
	"sync"

	validate "github.com/go-playground/validator"
)

var (
	once      sync.Once
	validator *validate.Validate
)

// New return instance of validator
func New() *validate.Validate {
	once.Do(func() {
		validator = validate.New()
	})
	return validator
}

// Validate the given struct base on the definition of 'validate' tag of the struct.
func Validate(v interface{}) error {
	return New().Struct(v)
}
