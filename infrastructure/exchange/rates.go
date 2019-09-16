package exchange

import (
	DomainRate "curve-tech-test/domain/rate"
	"encoding/json"
)

type Rates map[string]float64

func (r Rates) Value(currencyCode string) float64 {
	return r[currencyCode]
}

type RatesResponse struct {
	Rates       Rates  `json:"rates"`
	Base        string `json:"base"`
	CurrentDate string `json:"date"`
}

func NewRatesResponse() *RatesResponse {
	return &RatesResponse{}
}

func (e *RatesResponse) FromJSON(data []byte) ([]DomainRate.Rate, error) {
	err := json.Unmarshal(data, &e)
	if err != nil {
		return nil, err
	}

	return FromRatesResponse(e), nil
}
