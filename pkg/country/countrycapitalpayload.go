package country

import "encoding/xml"

type ResponseCountryCapital struct {
	XMLName  xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	SoapBody *SOAPCountryCapitalResponse
}

type SOAPCountryCapitalResponse struct {
	XMLName      xml.Name `xml:"Body"`
	Resp         *ResponseBodyCountryCapital
	FaultDetails *FaultCountryCapital
}

type FaultCountryCapital struct {
	XMLName     xml.Name `xml:"Fault"`
	Faultcode   string   `xml:"faultcode"`
	Faultstring string   `xml:"faultstring"`
}

type ResponseBodyCountryCapital struct {
	XMLName xml.Name `xml:"CapitalCityResponse"`
	Capital string   `xml:"CapitalCityResult"`
}
