package prediction

import (
	"curve-tech-test/domain/rate"
	"testing"
)

func TestPredict_WithValidPayload(t *testing.T) {
	mockPayload := []rate.Rate{
		&mockRate{},
		&mockRate{},
		&mockRate{},
		&mockRate{},
		&mockRate{},
		&mockRate{},
		&mockRate{},
	}

	sut, err := NewPredict(mockPayload, "GBP")
	if err != nil {
		t.Error(err)
	}

	if sut.Predict() == false {
		t.Fail()
	}
}

func TestPredict_WithInvalidPayload(t *testing.T) {
	mockPayload := []rate.Rate{
		&mockRate{},
		&mockRate{},
		&mockRate{},
	}

	_, err := NewPredict(mockPayload, "GBP")
	if err == nil {
		t.Fail()
	}
}

type mockRate struct{}

func (mr *mockRate) Date() string {
	return "2019-09-11"
}

func (mr *mockRate) Currency() string {
	return "GBP"
}

func (mr *mockRate) BassCurrency() string {
	return "USD"
}

func (mr *mockRate) Value() float64 {
	return 0.8
}
