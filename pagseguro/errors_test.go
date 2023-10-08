package pagseguro

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPagseguro_ApiErrors(t *testing.T) {
	t.Run("order error", func(t *testing.T) {
		errs := &ApiErrors{httpStatusCode: http.StatusBadRequest}
		errs.Parse(json.RawMessage(`
			{
				"code": "40001",
				"description": "required_parameter",
				"parameter_name": "payment_methods_is_required"
		  	}
		`))

		assert.Equal(t, "error processing request(http status code: 400)", errs.Error())
		assert.Equal(t, "code: 40001, desc: required_parameter, parameter: payment_methods_is_required", errs.ErrorMessages[0].Error())
		assert.Equal(
			t,
			[]ApiError{
				{Code: "40001", Description: "required_parameter", ParameterName: "payment_methods_is_required"},
			},
			errs.ErrorMessages,
		)
	})

	t.Run("charge error", func(t *testing.T) {
		errs := &ApiErrors{httpStatusCode: http.StatusBadRequest}
		errs.Parse(json.RawMessage(`
			{
				"error_messages": [
					{
						"code": "40001",
						"description": "required_parameter",
						"parameter_name": "payment_method.capture"
					},
					{
						"code": "40002",
						"description": "invalid_parameter",
						"parameter_name": "payment_methods_is_invalid"
					}
				]
			}
		`))

		assert.Equal(t, "error processing request(http status code: 400)", errs.Error())
		assert.Equal(t, "code: 40001, desc: required_parameter, parameter: payment_method.capture", errs.ErrorMessages[0].Error())
		assert.Equal(t, "code: 40002, desc: invalid_parameter, parameter: payment_methods_is_invalid", errs.ErrorMessages[1].Error())
		assert.Equal(
			t,
			[]ApiError{
				{Code: "40001", Description: "required_parameter", ParameterName: "payment_method.capture"},
				{Code: "40002", Description: "invalid_parameter", ParameterName: "payment_methods_is_invalid"},
			},
			errs.ErrorMessages,
		)
	})
}
