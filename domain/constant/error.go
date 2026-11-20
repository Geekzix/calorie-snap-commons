package constant

import "errors"

var (
	ErrNotFound     = errors.New("resource not found")
	ErrAlreadyExit  = errors.New("resource already exists")
	ErrInvalidInput = errors.New("invalid input")
)

type DomainError struct {
	Sentinel error
	Message  string
}

func (e *DomainError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	if e.Sentinel != nil {
		return e.Sentinel.Error()
	}
	return "internal server error"
}

func (e *DomainError) Unwrap() error {
	return e.Sentinel
}

func AlreadyExits(msg string) error {
	return &DomainError{Sentinel: ErrAlreadyExit, Message: msg}
}

func NotFound(msg string) error {
	return &DomainError{Sentinel: ErrNotFound, Message: msg}
}

func InvalidInput(msg string) error {
	return &DomainError{Sentinel: ErrInvalidInput, Message: msg}
}
