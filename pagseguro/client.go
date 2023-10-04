package pagseguro

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/go-resty/resty/v2"
)

const (
	SandboxEnvironment    = "https://sandbox.api.pagseguro.com"
	ProductionEnvironment = "https://api.pagseguro.com"
)

type (
	Client interface {
		CreateOrder(ctx context.Context, order *Order) error
	}
)

type (
	Default struct {
		client *resty.Client
	}
)

func New(baseUrl string, token string) Default {
	client := resty.New()
	client.SetHeader("Authorization", token)
	client.SetBaseURL(baseUrl)

	return Default{client: client}
}

func (d Default) CreateOrder(ctx context.Context, order *Order) error {
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
		return &Error{
			message:    "failed to send request",
			err:        err,
			statusCode: response.StatusCode(),
		}
	}

	if response.IsSuccess() {
		return nil
	}

	httpStatusResponseBadRequest := []int{
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusConflict,
	}
	if slices.Contains(httpStatusResponseBadRequest, response.StatusCode()) {
		errs := &ApiErrors{}
		errs.Parse(*response.Error().(*json.RawMessage))
		return errs
	}

	return &Error{
		message:    "internal server error",
		statusCode: response.StatusCode(),
	}
}
