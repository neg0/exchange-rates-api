package exchange

import (
	DomainRate "curve-tech-test/domain/rate"
	"encoding/json"
)

type History struct {
	Rates   map[string]Rates `json:"rates"`
	StartAt string           `json:"start_at"`
	Base    string           `json:"base"`
	EndAt   string           `json:"end_at"`
}

// Creates instance of History
func NewHistory() *History {
	return &History{}
}

// Factory method for History from API payload
func (erh *History) FromJSON(data []byte) ([]DomainRate.Rate, error) {
	err := json.Unmarshal(data, &erh)
	if err != nil {
		return nil, err
	}

	var rates []DomainRate.Rate
	for date := range erh.Rates {
		for currencyCode, currencyValue := range erh.Rates[date] {
			newRate := NewRate(
				date,
				currencyCode,
				erh.Base,
				currencyValue,
			)
			rates = append(rates, newRate)
		}
	}

	return rates, nil
}
