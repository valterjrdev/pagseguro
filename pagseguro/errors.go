package pagseguro

import (
	"encoding/json"
	"fmt"
)

type (
	ApiError struct {
		Code          string `json:"code,omitempty"`
		Description   string `json:"description,omitempty"`
		ParameterName string `json:"parameter_name,omitempty"`
	}

	ApiErrors struct {
		ErrorMessages []ApiError `json:"error_messages,omitempty"`

		err            error
		httpStatusCode int
	}
)

func (e ApiError) Error() string {
	return fmt.Sprintf("code: %s, desc: %s, parameter: %s", e.Code, e.Description, e.ParameterName)
}

func (e ApiErrors) Unwrap() error {
	return e.err
}

func (e ApiErrors) Error() string {
	if len(e.ErrorMessages) == 0 {
		return fmt.Sprintf("error processing request(http status code: %d): non-standard error response, contact pagseguro support", e.httpStatusCode)
	}

	return fmt.Sprintf("error processing request(http status code: %d)", e.httpStatusCode)
}

func (e *ApiErrors) Parse(data json.RawMessage) {
	apiError := ApiError{}
	e.ErrorMessages = make([]ApiError, 0)

	_ = json.Unmarshal(data, e)
	_ = json.Unmarshal(data, &apiError)

	if apiError.Code != "" {
		e.ErrorMessages = append(e.ErrorMessages, apiError)
	}
}
