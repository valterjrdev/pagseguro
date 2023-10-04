package pagseguro

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/go-resty/resty/v2"
)

type (
	Request interface {
		Endpoint() string
	}

	Client interface {
		Send(ctx context.Context, req Request) error
	}
)

type (
	Default struct {
		client *resty.Client
	}
)

func New(baseUrl string, token string, timeout time.Duration) Default {
	client := resty.New()
	client.SetHeader("Authorization", token)
	client.SetBaseURL(baseUrl)
	client.SetTimeout(timeout)

	return Default{client: client}
}

func (d Default) Send(ctx context.Context, req Request) error {
	response, err := d.client.R().
		SetBody(req).
		SetResult(req).
		SetError(&json.RawMessage{}).
		Post(req.Endpoint())
	if err != nil {
		return newError(response.StatusCode(), "failed to send request", err)
	}

	if response.IsSuccess() {
		return nil
	}

	httpStatusResponse := []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict}
	if slices.Contains(httpStatusResponse, response.StatusCode()) {
		errs := &ApiErrors{}
		errs.Parse(*response.Error().(*json.RawMessage))
		return errs
	}

	return newError(response.StatusCode(), "internal server error", nil)
}
