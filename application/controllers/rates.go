package controllers

import (
	"curve-tech-test/domain/rate"
	"curve-tech-test/infrastructure/api"
	"encoding/json"
	"net/http"
	"time"
)

func Rates(w http.ResponseWriter, r *http.Request) {
	baseCurrency := r.URL.Query().Get("base")
	if baseCurrency == "" {
		baseCurrency = "EUR"
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	result, err := api.NewExchangeAPI().TodayRate(baseCurrency)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	rates := make(map[string]float64)
	for _, res := range result {
		rates[res.Currency()] = res.Value()
	}

	ratesResp := &rate.Response{
		Rates: rates,
		Base:  baseCurrency,
		Date:  time.Now().Format(rate.DateFormat),
	}

	ratesJSON, _ := json.Marshal(ratesResp)

	_, _ = w.Write(ratesJSON)
}
