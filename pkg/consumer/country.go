package consumer

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/country"
)

type countryConsumer struct{}

var countryConsumerSingleton country.CountryConsumer

func GetCountryConsumer() country.CountryConsumer {
	if countryConsumerSingleton == nil {
		return getCountryConsumer()
	}
	return countryConsumerSingleton
}

func getCountryConsumer() *countryConsumer {
	return &countryConsumer{}
}

func SetCountryConsumer(consumer country.CountryConsumer) {
	countryConsumerSingleton = consumer
}

func (consumer countryConsumer) List() ([]country.Country, error) {
	payload := strings.NewReader(`<?xml version="1.0" encoding="utf-8"?>
  <soap12:Envelope xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
    <soap12:Body>
      <ListOfCountryNamesByName xmlns="http://www.oorsprong.org/websamples.countryinfo">
      </ListOfCountryNamesByName>
    </soap12:Body>
  </soap12:Envelope>`)

	rListCountries, err := SoapRequest(http.MethodPost, URL, payload, &country.ResponseCountryInfo{})
	if err != nil {
		return nil, err
	}

	r := rListCountries.(*country.ResponseCountryInfo)

	return r.SoapBody.Resp.Response.Countries, nil
}

func (consumer countryConsumer) Info(countryCode string) (*country.CountryInfo, error) {

	payload := strings.NewReader(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
  <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
    <soap:Body>
      <CapitalCity xmlns="http://www.oorsprong.org/websamples.countryinfo">
        <sCountryISOCode>%s</sCountryISOCode>
      </CapitalCity>
    </soap:Body>
  </soap:Envelope>`, countryCode))

	rCapital, err := SoapRequest(http.MethodPost, URL, payload, &country.ResponseCountryCapital{})
	if err != nil {
		return nil, err
	}

	rCap := rCapital.(*country.ResponseCountryCapital)

	countryCapital := rCap.SoapBody.Resp.Capital

	payload = strings.NewReader(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
  <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
    <soap:Body>
      <CountryCurrency xmlns="http://www.oorsprong.org/websamples.countryinfo">
        <sCountryISOCode>%s</sCountryISOCode>
      </CountryCurrency>
    </soap:Body>
  </soap:Envelope>`, countryCode))

	rCurrency, err := SoapRequest(http.MethodPost, URL, payload, &country.ResponseCountryCurrency{})
	if err != nil {
		return nil, err
	}

	rCurr := rCurrency.(*country.ResponseCountryCurrency)

	countryCurrencyCode := rCurr.SoapBody.Resp.Response.Code
	countryCurrencyName := rCurr.SoapBody.Resp.Response.Name

	countryCurrency := country.Currency{Code: countryCurrencyCode, Name: countryCurrencyName}

	payload = strings.NewReader(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
  <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
    <soap:Body>
      <CountryFlag xmlns="http://www.oorsprong.org/websamples.countryinfo">
        <sCountryISOCode>%s</sCountryISOCode>
      </CountryFlag>
    </soap:Body>
  </soap:Envelope>`, countryCode))

	rFlag, err := SoapRequest(http.MethodPost, URL, payload, &country.ResponseCountryFlag{})
	if err != nil {
		return nil, err
	}

	rF := rFlag.(*country.ResponseCountryFlag)

	countryFlag := rF.SoapBody.Resp.Flag

	payload = strings.NewReader(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
  <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
    <soap:Body>
      <CountryIntPhoneCode xmlns="http://www.oorsprong.org/websamples.countryinfo">
        <sCountryISOCode>%s</sCountryISOCode>
      </CountryIntPhoneCode>
    </soap:Body>
  </soap:Envelope>`, countryCode))

	rPhoneCode, err := SoapRequest(http.MethodPost, URL, payload, &country.ResponseCountryPhoneCode{})
	if err != nil {
		return nil, err
	}

	rPC := rPhoneCode.(*country.ResponseCountryPhoneCode)

	countryPhoneCode := rPC.SoapBody.Resp.PhoneCode

	countryInfo := country.CountryInfo{
		Capital:   countryCapital,
		Currency:  countryCurrency,
		Flag:      countryFlag,
		PhoneCode: countryPhoneCode,
	}

	return &countryInfo, nil
}
