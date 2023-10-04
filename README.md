# Pagseguro SDK


## Installation

```bash
go get github.com/valterjrdev/pagseguro
```

## Usage

#### 1° - Import
```go
import (
	"github.com/valterjrdev/pagseguro"
	"github.com/valterjrdev/pagseguro/orders"
)
```

#### 2° - initialize the client
```go
client := pagseguro.New(pagseguro.SandboxEnvironment, "{YOUR TOKEN HERE}")
```

#### 3° - Create your request (Boleto)
```go
request := &orders.Order{
		ReferenceID: "ex-00001",
		Customer: orders.Customer{
			Name:  "Jose da Silva",
			Email: "email@gmail.com",
			TaxID: "12345678909",
			Phones: []orders.Phone{
				{
					Country: "55",
					Area:    "11",
					Number:  "999999999",
					Type:    "MOBILE",
				},
			},
		},
		Items: []orders.Item{
			{
				ReferenceID: "referencia do item",
				Name:        "nome do item",
				Quantity:    1,
				UnitAmount:  500,
			},
		},
		Shipping: orders.Shipping{
			Address: orders.Address{
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
		Charges: []orders.Charge{
			{
				ReferenceID: "referencia da cobranca",
				Description: "descricao da cobranca",
				Amount: orders.Amount{
					Value:    500,
					Currency: "BRL",
				},
				PaymentMethod: orders.PaymentMethod{
					Type: "BOLETO",
					Boleto: orders.Boleto{
						DueDate: "2024-12-31",
						InstructionLines: orders.InstructionLines{
							Line1: "Pagamento processado para DESC Fatura",
							Line2: "Via PagSeguro",
						},
						Holder: orders.Holder{
							Name:  "Jose da Silva",
							TaxID: "22222222222",
							Email: "jose@email.com",
							Address: orders.Address{
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

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
    defer cancel()

	err := c.Send(ctx, request)
	if err != nil {
		log.Println(err)
		return
	}
```

#### 4° - Check the response in the request object
```go
log.Println(request.Charges)
```


## Getting help

If you have questions, concerns, bug reports, etc, please file an issue in this repository's Issue Tracker.