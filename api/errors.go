package api

const (
	// ErrServer should be returned on 500 errors
	ErrServer = "there was an issue processing this request"

	// ErrInvalidJSON should be returned when body is invalid
	ErrInvalidJSON = "that was not a valid json body"
)
