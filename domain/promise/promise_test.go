package promise

import (
	"net/http"
	"testing"
)

func TestPromise(t *testing.T) {
	sut := NewPromise()

	sut.Promise(func() (response *http.Response, err error) {
		rs := &http.Response{}
		rs.StatusCode = 400
		return rs, nil
	}).Then(func(response *http.Response) {
		t.Log(response.StatusCode)
	}).Catch(func(e error) {
		t.Log(e)
	}).Await()
}
