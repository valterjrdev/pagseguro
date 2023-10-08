package pagseguro

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestPagseguro_Client(t *testing.T) {
	t.Run("create order: boleto", func(t *testing.T) {
		token := gofakeit.LetterN(60)
		req, resp := generateMockObjectBoletoOrder(t)

		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, token, r.Header.Get("Authorization"))
			assert.Equal(t, createOrderEndpoint, r.URL.Path)
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
			assert.Equal(t, "application/json", r.Header.Get("Accept"))

			reqBody, err := io.ReadAll(r.Body)
			assert.NoError(t, err)

			reqMock, respMock := generateMockJsonBoletoOrder(t)
			assert.JSONEq(t, reqMock, string(reqBody))

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(respMock))
		}))

		defer srv.Close()

		client := New(srv.URL, token)
		err := client.CreateOrder(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, resp, req)
	})

	t.Run("create order: credit card", func(t *testing.T) {
		token := gofakeit.LetterN(60)
		req, resp := generateMockObjectCreditCardOrder(t)

		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, token, r.Header.Get("Authorization"))
			assert.Equal(t, createOrderEndpoint, r.URL.Path)
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
			assert.Equal(t, "application/json", r.Header.Get("Accept"))

			reqBody, err := io.ReadAll(r.Body)
			assert.NoError(t, err)

			reqMock, respMock := generateMockJsonCreditCardOrder(t)
			assert.JSONEq(t, reqMock, string(reqBody))

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(respMock))
		}))

		defer srv.Close()

		client := New(srv.URL, token)
		err := client.CreateOrder(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, resp, req)
	})

	t.Run("error responses", func(t *testing.T) {
		t.Run("order standard", func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`
					{
						"code": "40001",
						"description": "required_parameter",
						"parameter_name": "payment_methods_is_required"
					}
				`))
			}))

			defer srv.Close()

			client := New(srv.URL, "")
			err := client.CreateOrder(context.Background(), &Order{})
			assert.EqualError(t, err, "error processing request(http status code: 400)")
			assert.Equal(
				t,
				[]ApiError{
					{Code: "40001", Description: "required_parameter", ParameterName: "payment_methods_is_required"},
				},
				err.(*ApiErrors).ErrorMessages,
			)
		})

		t.Run("charge standard", func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`
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
			}))

			defer srv.Close()

			client := New(srv.URL, "")
			err := client.CreateOrder(context.Background(), &Order{})
			assert.EqualError(t, err, "error processing request(http status code: 400)")
			assert.Equal(
				t,
				[]ApiError{
					{Code: "40001", Description: "required_parameter", ParameterName: "payment_method.capture"},
					{Code: "40002", Description: "invalid_parameter", ParameterName: "payment_methods_is_invalid"},
				},
				err.(*ApiErrors).ErrorMessages,
			)
		})

		t.Run("internal server error", func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(``))
			}))

			defer srv.Close()

			client := New(srv.URL, "")
			err := client.CreateOrder(context.Background(), &Order{})
			assert.EqualError(t, err, "error processing request(http status code: 500): non-standard error response, contact pagseguro support")
		})

		t.Run("non-standard error response", func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`
					{
						"teste_message": "test"
					}
				`))
			}))

			defer srv.Close()

			client := New(srv.URL, "")
			err := client.CreateOrder(context.Background(), &Order{})
			assert.EqualError(t, err, "error processing request(http status code: 400): non-standard error response, contact pagseguro support")
		})
	})
}
