package country

import "encoding/xml"

type ResponseCountryCurrency struct {
	XMLName  xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	SoapBody *SOAPCountryCurrencyResponse
}

type SOAPCountryCurrencyResponse struct {
	XMLName      xml.Name `xml:"Body"`
	Resp         *ResponseBodyCountryCurrency
	FaultDetails *FaultCountryCurrency
}

type FaultCountryCurrency struct {
	XMLName     xml.Name `xml:"Fault"`
	Faultcode   string   `xml:"faultcode"`
	Faultstring string   `xml:"faultstring"`
}

type ResponseBodyCountryCurrency struct {
	XMLName  xml.Name `xml:"CountryCurrencyResponse"`
	Response *BodyCountryCurrency
}

type BodyCountryCurrency struct {
	XMLName xml.Name `xml:"CountryCurrencyResult"`
	Code    string   `xml:"sISOCode"`
	Name    string   `xml:"sName"`
}

type Currency struct {
	Code string
	Name string
}
