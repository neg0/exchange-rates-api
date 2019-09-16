package controllers

import (
	"curve-tech-test/domain/prediction"
	"curve-tech-test/infrastructure/api"
	"encoding/json"
	"net/http"
)

func Predict(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	currencyCode := r.URL.Query().Get("currency")
	if currencyCode == "" {
		currencyCode = "GBP"
	}

	result, err := api.NewExchangeAPI().LastWeekRates(currencyCode)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	predict, err := prediction.NewPredict(result, currencyCode)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	var respJSON struct {
		Sell bool `json:"sell"`
	}
	respJSON.Sell = predict.Predict()

	respBytes, _ := json.Marshal(respJSON)

	_, _ = w.Write(respBytes)
}
