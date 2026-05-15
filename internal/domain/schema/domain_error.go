// Package schema defines player domain models and validation.
package schema

import (
	"errors"
	"fmt"
)

var (
	ErrEmpty          = errors.New("cannot be empty")
	ErrInvalidFormat  = errors.New("invalid format")
	ErrMustBePositive = errors.New("must be a positive value")
	ErrTooLong        = errors.New("too long")
	ErrMustBePast     = errors.New("must be a past date")
)

type DomainError struct {
	Operation string
	Entity    string
	Field     string
	EntityID  string
	Reason    error
}

func (e DomainError) Error() string {
	return fmt.Sprintf("%s %s failed on %s: %s", e.Operation, e.Entity, e.Field, e.Reason)
}

func (e DomainError) Unwrap() error {
	return e.Reason
}
