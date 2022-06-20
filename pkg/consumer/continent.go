package consumer

import (
	"net/http"
	"strings"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/continent"
)

type continentConsumer struct{}

var continentConsumerSingleton continent.ContinentConsumer

func GetContinentConsumer() continent.ContinentConsumer {
	if continentConsumerSingleton == nil {
		return getContinentConsumer()
	}
	return continentConsumerSingleton
}

func getContinentConsumer() *continentConsumer {
	return &continentConsumer{}
}

func SetContinentConsumer(consumer continent.ContinentConsumer) {
	continentConsumerSingleton = consumer
}

func (consumer continentConsumer) List() ([]continent.Continent, error) {

	payload := strings.NewReader(`<?xml version="1.0" encoding="utf-8"?>
	<soap12:Envelope xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
	<soap12:Body>
		<ListOfContinentsByName xmlns="http://www.oorsprong.org/websamples.countryinfo">
		</ListOfContinentsByName>
	</soap12:Body>
	</soap12:Envelope>`)

	rContinent, err := SoapRequest(http.MethodPost, URL, payload, &continent.ResponseContinentInfo{})
	if err != nil {
		return nil, err
	}

	r := rContinent.(*continent.ResponseContinentInfo)

	return r.SoapBody.Resp.Response.Continents, nil
}
