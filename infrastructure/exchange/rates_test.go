package exchange

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	mockPayload, err := ioutil.ReadFile("./.mock-rates.json")
	if err != nil {
		t.Error(err)
	}

	sut := NewRatesResponse()

	_, err = sut.FromJSON(mockPayload)
	if err != nil {
		t.Error(err)
	}

	if len(sut.Rates) != 32 {
		t.Error(len(sut.Rates))
	}

	if reflect.TypeOf(sut.Rates["GBP"]).String() != "float64" {
		t.Fail()
	}
}
