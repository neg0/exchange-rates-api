package api

import (
	"curve-tech-test/domain/promise"
	"curve-tech-test/domain/rate"
	"curve-tech-test/infrastructure/exchange"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func (e *exchangeAPI) TodayRate(base string) ([]rate.Rate, error) {
	var hasError error
	var result []rate.Rate

	promise.NewPromise().Promise(func() (response *http.Response, err error) {
		url := fmt.Sprintf(
			"%s/latest?symbols=USD,GBP&base=%s",
			os.Getenv("EXCHANGE_API_URI"),
			base,
		)

		return e.HTTPClient.Get(url)
	}).Then(func(response *http.Response) {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			hasError = err
			return
		}

		newRates, err := exchange.NewRatesResponse().FromJSON(bodyBytes)
		if err != nil {
			return
		}

		if len(newRates) < 1 {
			hasError = errors.New("error creating rates object")
			return

		}
		result = newRates
	}).Catch(func(e error) {
		hasError = e
	}).Await()

	return result, hasError
}
