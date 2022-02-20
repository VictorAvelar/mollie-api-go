package mollie

import "fmt"

// ErrorLinks container references to common urls
// returned with errors.
type ErrorLinks struct {
	Documentation *URL `json:"documentation,omitempty"`
}

// BaseError contains the general error structure
// returned by mollie.
type BaseError struct {
	Status int         `json:"status,omitempty"`
	Title  string      `json:"title,omitempty"`
	Detail string      `json:"detail,omitempty"`
	Field  string      `json:"field,omitempty"`
	Links  *ErrorLinks `json:"_links,omitempty"`
}

// Error interface compliance.
func (be *BaseError) Error() string {
	return fmt.Sprintf("%d %s: %s", be.Status, be.Title, be.Detail)
}
