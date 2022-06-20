package consumer

import (
	"github.com/rafaelsanzio/go-consuming-soap/pkg/continent"
)

type MockContinentConsumer struct {
	continent.ContinentConsumer
	ListFunc func() ([]continent.Continent, error)
}

func (m MockContinentConsumer) List() ([]continent.Continent, error) {
	if m.ListFunc != nil {
		return m.ListFunc()
	}
	return m.ContinentConsumer.List()
}
