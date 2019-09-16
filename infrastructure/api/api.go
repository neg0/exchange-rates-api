package api

import (
	"curve-tech-test/domain/rate"
	"net/http"
)

type ExchangeAPI interface {
	TodayRate(base string) ([]rate.Rate, error)
	LastWeekRates(currencyCode string) ([]rate.Rate, error)
}

type exchangeAPI struct {
	HTTPClient
}

func NewExchangeAPI() ExchangeAPI {
	return &exchangeAPI{
		&http.Client{},
	}
}
