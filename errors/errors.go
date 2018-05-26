package errors

import "fmt"

// E is an error struct that contains meta
type E struct {
	Code    int
	Message string
	Details string
}

func (e *E) Error() string {
	return fmt.Sprintf("%d: %s - %s", e.Code, e.Message, e.Details)
}

// New returns a new error
func New(err string, code int, details string) *E {
	return &E{Code: code, Message: err}
}
