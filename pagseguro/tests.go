package pagseguro

import "testing"

func generateMockObjectBoletoOrder(t *testing.T) (*Order, *Order) {
	t.Helper()

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

	response := &Order{
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

	return req, response
}

func generateMockJsonBoletoOrder(t *testing.T) (string, string) {
	t.Helper()

	req := `{
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
	}`

	resp := `
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
	`

	return req, resp
}

func generateMockObjectCreditCardOrder(t *testing.T) (*Order, *Order) {
	t.Helper()
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
					Type:         "CREDIT_CARD",
					Installments: 1,
					Capture:      true,
					Card: Card{
						Number:       "4111111111111111",
						ExpMonth:     "12",
						ExpYear:      "2026",
						SecurityCode: "123",
						Holder: Holder{
							Name: "Jose da Silva",
						},
						Store: false,
					},
				},
			},
		},
	}

	response := &Order{
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
				ID:          "CHAR_0A58650D-4A07-4007-836A-5959A7D28158",
				ReferenceID: "referencia da cobranca",
				Description: "descricao da cobranca",
				Amount: Amount{
					Value:    500,
					Currency: "BRL",
					Summary: Summary{
						Total: 500,
						Paid:  500,
					},
				},
				Status: "PAID",
				PaymentMethod: PaymentMethod{
					Type:           "CREDIT_CARD",
					Installments:   1,
					Capture:        true,
					SoftDescriptor: "sellervirtual",
					Card: Card{
						Number:       "4111111111111111",
						ExpMonth:     "12",
						ExpYear:      "2026",
						SecurityCode: "123",
						Holder: Holder{
							Name: "Jose da Silva",
						},
						Store: false,
					},
				},
				PaymentResponse: PaymentResponse{
					Code:    "20000",
					Message: "SUCESSO",
				},
				Links: []Links{
					{
						Rel:   "SELF",
						Href:  "https://sandbox.api.pagseguro.com/charges/CHAR_0A58650D-4A07-4007-836A-5959A7D28158",
						Media: "application/json",
						Type:  "GET",
					},
					{
						Rel:   "CHARGE.CANCEL",
						Href:  "https://sandbox.api.pagseguro.com/charges/CHAR_0A58650D-4A07-4007-836A-5959A7D28158/cancel",
						Media: "application/json",
						Type:  "POST",
					},
				},
				CreatedAt: "2023-10-08T00:21:23.106-03:00",
				PaidAt:    "2023-10-08T00:21:23.000-03:00",
			},
		},
	}

	return req, response
}

func generateMockJsonCreditCardOrder(t *testing.T) (string, string) {
	t.Helper()

	req := `
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
				 "type":"CREDIT_CARD",
				 "installments":1,
				 "capture":true,
				 "card":{
					"holder":{
					   "name":"Jose da Silva",
					   "address":{
						  
					   }
					},
					"number":"4111111111111111",
					"exp_month":"12",
					"exp_year":"2026",
					"security_code":"123"
				 },
				 "boleto":{
					"instruction_lines":{
					   
					},
					"holder":{
					   "address":{
						  
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
	`

	resp := `
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
			  "id":"CHAR_0A58650D-4A07-4007-836A-5959A7D28158",
			  "reference_id":"referencia da cobranca",
			  "description":"descricao da cobranca",
			  "amount":{
				 "value":500,
				 "currency":"BRL",
				 "summary":{
					"total":500,
					"paid":500
				 }
			  },
			  "status":"PAID",
			  "payment_method":{
				 "type":"CREDIT_CARD",
				 "installments":1,
				 "capture":true,
				 "soft_descriptor":"sellervirtual",
				 "card":{
					"holder":{
					   "name":"Jose da Silva",
					   "address":{
						  
					   }
					},
					"number":"4111111111111111",
					"exp_month":"12",
					"exp_year":"2026",
					"security_code":"123"
				 },
				 "boleto":{
					"instruction_lines":{
					   
					},
					"holder":{
					   "address":{
						  
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
					"href":"https://sandbox.api.pagseguro.com/charges/CHAR_0A58650D-4A07-4007-836A-5959A7D28158",
					"media":"application/json",
					"type":"GET"
				 },
				 {
					"rel":"CHARGE.CANCEL",
					"href":"https://sandbox.api.pagseguro.com/charges/CHAR_0A58650D-4A07-4007-836A-5959A7D28158/cancel",
					"media":"application/json",
					"type":"POST"
				 }
			  ],
			  "created_at":"2023-10-08T00:21:23.106-03:00",
			  "paid_at":"2023-10-08T00:21:23.000-03:00"
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
	`

	return req, resp
}
