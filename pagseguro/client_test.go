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
	t.Run("create order (Boleto)", func(t *testing.T) {
		token := gofakeit.LetterN(60)

		req := &Order{
			ReferenceID: "ex-00001",
			Customer: Customer{
				Name:  "Jose da Silva",
				Email: "email@gmail.com",
				TaxID: "12345678909",
				Phones: []Phone{
					{
						Country: "55",
						Area:    "11",
						Number:  "999999999",
						Type:    "MOBILE",
					},
				},
			},
			Items: []Item{
				{
					ReferenceID: "referencia do item",
					Name:        "nome do item",
					Quantity:    1,
					UnitAmount:  500,
				},
			},
			Shipping: Shipping{
				Address: Address{
					Street:     "Avenida Brigadeiro Faria Lima",
					Number:     "1384",
					Locality:   "Pinheiros",
					City:       "São Paulo",
					Region:     "São Paulo",
					RegionCode: "SP",
					Country:    "BRA",
					PostalCode: "01452002",
				},
			},
			NotificationUrls: []string{"https://meusite.com/notificacoes"},
			Charges: []Charge{
				{
					ReferenceID: "referencia da cobranca",
					Description: "descricao da cobranca",
					Amount: Amount{
						Value:    500,
						Currency: "BRL",
					},
					PaymentMethod: PaymentMethod{
						Type: "BOLETO",
						Boleto: Boleto{
							DueDate: "2024-12-31",
							InstructionLines: InstructionLines{
								Line1: "Pagamento processado para DESC Fatura",
								Line2: "Via PagSeguro",
							},
							Holder: Holder{
								Name:  "Jose da Silva",
								TaxID: "22222222222",
								Email: "jose@email.com",
								Address: Address{
									Country:    "Brasil",
									Region:     "São Paulo",
									RegionCode: "SP",
									City:       "Sao Paulo",
									PostalCode: "01452002",
									Street:     "Avenida Brigadeiro Faria Lima",
									Number:     "1384",
									Locality:   "Pinheiros",
								},
							},
						},
					},
				},
			},
		}

		expected := &Order{
			ReferenceID: "ex-00001",
			Customer: Customer{
				Name:  "Jose da Silva",
				Email: "email@gmail.com",
				TaxID: "12345678909",
				Phones: []Phone{
					{
						Country: "55",
						Area:    "11",
						Number:  "999999999",
						Type:    "MOBILE",
					},
				},
			},
			Items: []Item{
				{
					ReferenceID: "referencia do item",
					Name:        "nome do item",
					Quantity:    1,
					UnitAmount:  500,
				},
			},
			Shipping: Shipping{
				Address: Address{
					Street:     "Avenida Brigadeiro Faria Lima",
					Number:     "1384",
					Locality:   "Pinheiros",
					City:       "São Paulo",
					Region:     "São Paulo",
					RegionCode: "SP",
					Country:    "BRA",
					PostalCode: "01452002",
				},
			},
			NotificationUrls: []string{"https://meusite.com/notificacoes"},
			Charges: []Charge{
				{
					ID:          "CHAR_869FFC81-1418-48F9-B6C7-9549317D9852",
					ReferenceID: "referencia da cobranca",
					Description: "descricao da cobranca",
					Amount: Amount{
						Value:    500,
						Currency: "BRL",
						Summary: Summary{
							Total: 500,
						},
					},
					Status: "WAITING",
					PaymentMethod: PaymentMethod{
						Type: "BOLETO",
						Boleto: Boleto{
							DueDate: "2024-12-31",
							InstructionLines: InstructionLines{
								Line1: "Pagamento processado para DESC Fatura",
								Line2: "Via PagSeguro",
							},
							Holder: Holder{
								Name:  "Jose da Silva",
								TaxID: "22222222222",
								Email: "jose@email.com",
								Address: Address{
									Country:    "Brasil",
									Region:     "São Paulo",
									RegionCode: "SP",
									City:       "Sao Paulo",
									PostalCode: "01452002",
									Street:     "Avenida Brigadeiro Faria Lima",
									Number:     "1384",
									Locality:   "Pinheiros",
								},
							},
						},
					},
					PaymentResponse: PaymentResponse{
						Code:    "20000",
						Message: "SUCESSO",
					},
					Links: []Links{
						{
							Rel:   "SELF",
							Href:  "https://boleto.sandbox.pagseguro.com.br/0bc2c68a-1d40-4323-8591-86ee92ba594d.pdf",
							Media: "application/pdf",
							Type:  "GET",
						},
						{
							Rel:   "SELF",
							Href:  "https://boleto.sandbox.pagseguro.com.br/0bc2c68a-1d40-4323-8591-86ee92ba594d.png",
							Media: "image/png",
							Type:  "GET",
						},
						{
							Rel:   "SELF",
							Href:  "https://sandbox.api.pagseguro.com/charges/CHAR_869FFC81-1418-48F9-B6C7-9549317D9852",
							Media: "application/json",
							Type:  "GET",
						},
					},
					CreatedAt: "2023-10-07T23:22:10.893-03:00",
					PaidAt:    "0001-01-01T00:00:00Z",
				},
			},
		}

		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, token, r.Header.Get("Authorization"))
			assert.Equal(t, createOrderEndpoint, r.URL.Path)
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
			assert.Equal(t, "application/json", r.Header.Get("Accept"))

			requestBody, err := io.ReadAll(r.Body)
			assert.NoError(t, err)

			assert.JSONEq(t, `
				{
					"reference_id":"ex-00001",
					"customer":{
					"name":"Jose da Silva",
					"email":"email@gmail.com",
					"tax_id":"12345678909",
					"phones":[
						{
							"country":"55",
							"area":"11",
							"number":"999999999",
							"type":"MOBILE"
						}
					]
					},
					"shipping":{
					"address":{
						"street":"Avenida Brigadeiro Faria Lima",
						"number":"1384",
						"locality":"Pinheiros",
						"city":"São Paulo",
						"region":"São Paulo",
						"region_code":"SP",
						"country":"BRA",
						"postal_code":"01452002"
					}
					},
					"billing":{
					"address":{
						
					}
					},
					"charges":[
						{
							"reference_id":"referencia da cobranca",
							"description":"descricao da cobranca",
							"amount":{
								"value":500,
								"currency":"BRL",
								"summary":{
									
								}
							},
							"payment_method":{
								"type":"BOLETO",
								"card":{
									"holder":{
									"address":{
										
									}
									}
								},
								"boleto":{
									"due_date":"2024-12-31",
									"instruction_lines":{
									"line_1":"Pagamento processado para DESC Fatura",
									"line_2":"Via PagSeguro"
									},
									"holder":{
									"name":"Jose da Silva",
									"tax_id":"22222222222",
									"email":"jose@email.com",
									"address":{
										"street":"Avenida Brigadeiro Faria Lima",
										"number":"1384",
										"locality":"Pinheiros",
										"city":"Sao Paulo",
										"region":"São Paulo",
										"region_code":"SP",
										"country":"Brasil",
										"postal_code":"01452002"
									}
									}
								}
							},
							"payment_response":{
								
							}
						}
					],
					"items":[
					{
						"reference_id":"referencia do item",
						"name":"nome do item",
						"quantity":1,
						"unit_amount":500
					}
					],
					"notification_urls":[
					"https://meusite.com/notificacoes"
					]
				}
			`, string(requestBody))

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`
				{
					"reference_id":"ex-00001",
					"customer":{
					"name":"Jose da Silva",
					"email":"email@gmail.com",
					"tax_id":"12345678909",
					"phones":[
						{
							"country":"55",
							"area":"11",
							"number":"999999999",
							"type":"MOBILE"
						}
					]
					},
					"shipping":{
					"address":{
						"street":"Avenida Brigadeiro Faria Lima",
						"number":"1384",
						"locality":"Pinheiros",
						"city":"São Paulo",
						"region":"São Paulo",
						"region_code":"SP",
						"country":"BRA",
						"postal_code":"01452002"
					}
					},
					"charges":[
						{
							"id":"CHAR_869FFC81-1418-48F9-B6C7-9549317D9852",
							"reference_id":"referencia da cobranca",
							"description":"descricao da cobranca",
							"amount":{
								"value":500,
								"currency":"BRL",
								"summary":{
									"total":500
								}
							},
							"status":"WAITING",
							"payment_method":{
								"type":"BOLETO",
								"card":{
									"holder":{
									"address":{
										
									}
									}
								},
								"boleto":{
									"due_date":"2024-12-31",
									"instruction_lines":{
									"line_1":"Pagamento processado para DESC Fatura",
									"line_2":"Via PagSeguro"
									},
									"holder":{
									"name":"Jose da Silva",
									"tax_id":"22222222222",
									"email":"jose@email.com",
									"address":{
										"street":"Avenida Brigadeiro Faria Lima",
										"number":"1384",
										"locality":"Pinheiros",
										"city":"Sao Paulo",
										"region":"São Paulo",
										"region_code":"SP",
										"country":"Brasil",
										"postal_code":"01452002"
									}
									}
								}
							},
							"payment_response":{
								"code":"20000",
								"message":"SUCESSO"
							},
							"links":[
								{
									"rel":"SELF",
									"href":"https://boleto.sandbox.pagseguro.com.br/0bc2c68a-1d40-4323-8591-86ee92ba594d.pdf",
									"media":"application/pdf",
									"type":"GET"
								},
								{
									"rel":"SELF",
									"href":"https://boleto.sandbox.pagseguro.com.br/0bc2c68a-1d40-4323-8591-86ee92ba594d.png",
									"media":"image/png",
									"type":"GET"
								},
								{
									"rel":"SELF",
									"href":"https://sandbox.api.pagseguro.com/charges/CHAR_869FFC81-1418-48F9-B6C7-9549317D9852",
									"media":"application/json",
									"type":"GET"
								}
							],
							"created_at":"2023-10-07T23:22:10.893-03:00",
							"paid_at":"0001-01-01T00:00:00Z"
						}
					],
					"items":[
					{
						"reference_id":"referencia do item",
						"name":"nome do item",
						"quantity":1,
						"unit_amount":500
					}
					],
					"notification_urls":[
					"https://meusite.com/notificacoes"
					]
				}
			`))
		}))

		defer srv.Close()

		client := New(srv.URL, token)
		err := client.CreateOrder(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, expected, req)
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
