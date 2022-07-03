package pkg

import "errors"

var (
	ErrInvalidURL        = errors.New("URL not found.")
	ErrMethodNotAllow    = errors.New("Metode not allowed")
	ErrFormatRequestBody = errors.New("Invalid request body.")
	ErrHeaderInvalid     = errors.New("Invalid header.")
	ErrInternalServer    = errors.New("Internal Server Error.")
	ErrDataNotFound      = errors.New("Data Not Found")
)
