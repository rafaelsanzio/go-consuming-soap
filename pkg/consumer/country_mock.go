package consumer

import (
	"github.com/rafaelsanzio/go-consuming-soap/pkg/country"
)

type MockCountryConsumer struct {
	country.CountryConsumer
	ListFunc func() ([]country.Country, error)
	InfoFunc func(countryCode string) (*country.CountryInfo, error)
}

func (m MockCountryConsumer) List() ([]country.Country, error) {
	if m.ListFunc != nil {
		return m.ListFunc()
	}
	return m.CountryConsumer.List()
}

func (m MockCountryConsumer) Info(countryCode string) (*country.CountryInfo, error) {
	if m.InfoFunc != nil {
		return m.InfoFunc(countryCode)
	}
	return m.CountryConsumer.Info(countryCode)
}
