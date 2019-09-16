package exchange

import DomainRate "curve-tech-test/domain/rate"

type rate struct {
	date         string
	currency     string
	bassCurrency string
	value        float64
}

// Forcing the business logic that Rates initial rate should have these properties set
func NewRate(date string, currency string, bassCurrency string, value float64) DomainRate.Rate {
	return &rate{
		date,
		currency,
		bassCurrency,
		value,
	}
}

// Factory method to create a rate consistent with Domain rate from Rates Response
func FromRatesResponse(rr *RatesResponse) []DomainRate.Rate {
	var collection []DomainRate.Rate
	for currencyCode := range rr.Rates {
		newRate := NewRate(
			rr.CurrentDate,
			currencyCode,
			rr.Base,
			rr.Rates[currencyCode],
		)
		collection = append(collection, newRate)
	}

	return collection
}

func (r *rate) Date() string {
	return r.date
}

func (r *rate) Currency() string {
	return r.currency
}

func (r *rate) BassCurrency() string {
	return r.bassCurrency
}

func (r *rate) Value() float64 {
	return r.value
}
