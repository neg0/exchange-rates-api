package prediction

import (
	"curve-tech-test/domain/rate"
	"errors"
	"time"
)

type Predict interface {
	Predict() bool
}

type predict struct {
	rates        []rate.Rate
	currencyCode string
}

func NewPredict(rates []rate.Rate, currencyCode string) (Predict, error) {
	if len(rates) < 7 {
		return nil, errors.New("can not predict with less than seven rates")
	}

	return &predict{
		rates[0:7],
		currencyCode,
	}, nil
}

// Simple logic determine to sell or not
// It calculates number of days value has been above today's value and compares it to average of rates
func (p *predict) Predict() bool {
	var comparableRate float64

	// there is no specification on their API Timezone I extract one day for safety
	today := time.Now().AddDate(0, 0, -1).Format(rate.DateFormat)
	for _, singleRate := range p.rates[0:] {
		if singleRate.Date() == today {
			comparableRate = singleRate.Value()
		}
	}

	counter := 0
	for _, singleRate := range p.rates {
		if singleRate.Value() > comparableRate {
			counter++
		}
	}

	average := int(float64(len(p.rates) / 2))

	return counter > average
}
