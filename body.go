package httpwrapper

import (
	"net/http"

	"github.com/goccy/go-json"
)

const (
	// ErrCodeParsingBody is the error code returned to the user when there is an error parsing
	// the body of the request.
	ErrCodeParsingBody = "error_parsing_body"
)

// BindBody will bind the body of the request to the given interface.
func BindBody(r *http.Request, target interface{}) *ErrorResponse {
	// nolint: gocritic
	// LATER: add more encodings
	switch r.Header.Get("Content-Type") {
	default:
		if err := json.NewDecoder(r.Body).Decode(target); err != nil {
			return &ErrorResponse{
				Error:          err,
				ErrorCode:      ErrCodeParsingBody,
				HTTPStatusCode: http.StatusBadRequest,
				ErrorMsg:       "error parsing body",
			}
		}
	}

	return nil
}
