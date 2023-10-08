package pagseguro

import (
	"context"
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/valterjrdev/pagseguro-sdk-go/pagseguro/models"
)

const (
	SandboxEnvironment    = "https://sandbox.api.pagseguro.com"
	ProductionEnvironment = "https://api.pagseguro.com"
)

type (
	Order interface {
		CreateOrder(ctx context.Context, order *models.Order) error
	}
)

type (
	Default struct {
		client *resty.Client
	}
)

func New(baseUrl string, token string) Default {
	client := resty.New()
	client.SetBaseURL(baseUrl)
	client.SetHeader("Authorization", token)
	client.SetHeader("Content-Type", "application/json")
	client.SetHeader("Accept", "application/json")

	return Default{client: client}
}

func (d Default) CreateOrder(ctx context.Context, order *models.Order) error {
	response, err := d.client.R().
		SetContext(ctx).
		SetBody(order).
		SetResult(order).
		SetError(&json.RawMessage{}).
		Post("/orders")

	return d.handler(response, err)
}

func (d Default) handler(response *resty.Response, err error) error {
	if err != nil {
		return &ApiErrors{
			err:            err,
			httpStatusCode: response.StatusCode(),
		}
	}

	if response.IsError() {
		errsResponse := &ApiErrors{httpStatusCode: response.StatusCode()}
		errsResponse.Parse(*response.Error().(*json.RawMessage))
		return errsResponse
	}

	return nil
}
