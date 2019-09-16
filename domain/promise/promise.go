package promise

import (
	"net/http"
	"sync"
)

type resolve func(*http.Response)
type reject func(error)
type execution func() (*http.Response, error)

type Promise interface {
	Promise(f execution) Promise
	Then(f resolve) Promise
	Catch(f reject) Promise
	Await()
}

type promise struct {
	resolve
	reject
	sync.WaitGroup
}

func NewPromise() Promise {
	return &promise{}
}

func (p *promise) Promise(f execution) Promise {
	p.Add(1)
	go func(p *promise, w *sync.WaitGroup) {
		resp, err := f()
		if err != nil {
			p.reject(err)
		} else {
			p.resolve(resp)
		}

		p.Done()
	}(p, &p.WaitGroup)

	return p
}

func (p *promise) Then(f resolve) Promise {
	p.Add(1)
	p.resolve = f
	p.Done()

	return p
}

func (p *promise) Catch(f reject) Promise {
	p.Add(1)
	p.reject = f
	p.Done()

	return p
}

func (p *promise) Await() {
	p.Wait()
}
