package pagseguro

import (
	"encoding/json"
	"errors"
	"fmt"
)

type (
	Error struct {
		message    string
		err        error
		statusCode int
	}
)

type (
	ApiError struct {
		Code          string `json:"code,omitempty"`
		Description   string `json:"description,omitempty"`
		ParameterName string `json:"parameter_name,omitempty"`
	}

	ApiErrors struct {
		ErrorMessages []ApiError `json:"error_messages,omitempty"`
	}
)

func newError(statusCode int, message string, err error) *Error {
	return &Error{
		message:    message,
		err:        err,
		statusCode: statusCode,
	}
}

func (c Error) Error() string {
	return fmt.Sprintf("status code: %d, message: %s", c.statusCode, c.message)
}

func (c Error) Unwrap() error {
	return c.err
}

func (e ApiError) Error() string {
	return fmt.Sprintf("code: %s, desc: %s, parameter: %s;", e.Code, e.Description, e.ParameterName)
}

func (e ApiErrors) Error() string {
	var errs []error
	for _, err := range e.ErrorMessages {
		errs = append(errs, err)
	}

	return errors.Join(errs...).Error()
}

func (e ApiErrors) JSON() string {
	JSON, _ := json.Marshal(e.ErrorMessages)
	return string(JSON)
}

func (e *ApiErrors) Parse(data json.RawMessage) {
	_ = json.Unmarshal(data, e)

	if len(e.ErrorMessages) == 0 {
		var err ApiError
		_ = json.Unmarshal(data, &err)
		e.ErrorMessages = append(e.ErrorMessages, err)
	}
}
