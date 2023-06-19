package apperr

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidInput         = errors.New("invalid input")
	ErrConflict             = fmt.Errorf("conflict: %w", ErrInvalidInput)
	ErrUpdatedAtDoesntMatch = errors.New("could not update secrete. Local data doesn't match with server")
	ErrSecretNotFound       = errors.New("data not found")
)
