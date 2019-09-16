package exchange

import (
	"io/ioutil"
	"testing"
)

func TestExchangeRatesCreationFromJSONBytes(t *testing.T) {
	mockPayload, err := ioutil.ReadFile("./.mock-history.json")
	if err != nil {
		t.Error(err)
	}

	sut, err := NewHistory().FromJSON(mockPayload)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(sut) != 7 {
		t.Error(len(sut))
	}

	for _, s := range sut {
		t.Logf("%s %f", s.Currency(), s.Value())
	}
}
