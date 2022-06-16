package country

import "encoding/xml"

type ResponseCountryFlag struct {
	XMLName  xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	SoapBody *SOAPCountryFlagResponse
}

type SOAPCountryFlagResponse struct {
	XMLName      xml.Name `xml:"Body"`
	Resp         *ResponseBodyCountryFlag
	FaultDetails *FaultCountryFlag
}

type FaultCountryFlag struct {
	XMLName     xml.Name `xml:"Fault"`
	Faultcode   string   `xml:"faultcode"`
	Faultstring string   `xml:"faultstring"`
}

type ResponseBodyCountryFlag struct {
	XMLName xml.Name `xml:"CountryFlagResponse"`
	Flag    string   `xml:"CountryFlagResult"`
}
