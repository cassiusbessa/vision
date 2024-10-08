package errors

import "errors"

var (
	ErrPostNotFound = errors.New("post not found")
)

type ResourceNotFound struct {
	Message string
	Cause   error
}

func (r *ResourceNotFound) Error() string {
	return r.Message
}

func NewResourceNotFound(message string) *ResourceNotFound {
	return &ResourceNotFound{
		Message: message,
	}
}

type ResourceAlreadyExists struct {
	Message string
}

func (r *ResourceAlreadyExists) Error() string {
	return r.Message
}

func NewResourceAlreadyExists(message string) *ResourceAlreadyExists {
	return &ResourceAlreadyExists{
		Message: message,
	}
}

type InvalidArgument struct {
	Message string
	Cause   error
}

func (i *InvalidArgument) Error() string {
	return i.Message
}

func NewInvalidArgument(message string) *InvalidArgument {
	return &InvalidArgument{
		Message: message,
	}
}

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		Message: message,
	}
}

type Unauthorized struct {
	Message string
}

func (u *Unauthorized) Error() string {
	return u.Message
}

func NewUnauthorized(message string) *Unauthorized {
	return &Unauthorized{
		Message: message,
	}
}
