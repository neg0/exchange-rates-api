package api

import (
	"curve-tech-test/domain/promise"
	"curve-tech-test/domain/rate"
	"curve-tech-test/infrastructure/exchange"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func (e *exchangeAPI) LastWeekRates(currencyCode string) ([]rate.Rate, error) {
	var hasError error
	var result []rate.Rate

	promise.NewPromise().Promise(func() (response *http.Response, err error) {
		today := time.Now()

		// ExchangesAPI is not consistent with returned result
		// for last week returns sometimes 3 and sometimes 5 that's why I addded -20 as a buffer
		overAWeekAgo := today.AddDate(0, 0, -20)

		formattedToday := today.Format(rate.DateFormat)
		formattedAWeekAgo := overAWeekAgo.Format(rate.DateFormat)
		uri := os.Getenv("EXCHANGE_API_URI")

		url := fmt.Sprintf(
			"%s/history?start_at=%s&end_at=%s&symbols=%s&base=EUR",
			uri,
			formattedAWeekAgo,
			formattedToday,
			currencyCode,
		)

		return e.HTTPClient.Get(url)
	}).Then(func(response *http.Response) {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			hasError = err
			return
		}

		newRates, err := exchange.NewHistory().FromJSON(bodyBytes)
		if err != nil {
			hasError = err
		}
		result = newRates
	}).Catch(func(e error) {
		hasError = e
	}).Await()

	return result, hasError
}
