package pagseguro

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPagseguro_Error(t *testing.T) {
	err := newError(http.StatusInternalServerError, "err custom", errors.New("err internal"))
	assert.Equal(t, "status code: 500, message: err custom", err.Error())
	assert.EqualError(t, err.Unwrap(), "err internal")
}

func TestPagseguro_ApiErrors(t *testing.T) {
	t.Run("order error", func(t *testing.T) {
		errs := &ApiErrors{}
		errs.Parse(json.RawMessage(`
			{
				"code": "40001",
				"description": "required_parameter",
				"parameter_name": "name"
		  	}
		`))

		assert.EqualError(t, errs, "code: 40001, desc: required_parameter, parameter: name;")
		assert.JSONEq(t, `[{"code":"40001","description":"required_parameter","parameter_name":"name"}]`, errs.JSON())
	})

	t.Run("charge error", func(t *testing.T) {
		errs := &ApiErrors{}
		errs.Parse(json.RawMessage(`
			{
				"error_messages": [
					{
						"code": "40001",
						"description": "required_parameter",
						"parameter_name": "name"
					},
					{
						"code": "40002",
						"description": "invalid_parameter",
						"parameter_name": "number"
					}
				]
			}
		`))

		assert.EqualError(t, errs, "code: 40001, desc: required_parameter, parameter: name;\ncode: 40002, desc: invalid_parameter, parameter: number;")
		assert.JSONEq(t, `[{"code":"40001","description":"required_parameter","parameter_name":"name"}, {"code":"40002","description":"invalid_parameter","parameter_name":"number"}]`, errs.JSON())
	})
}
