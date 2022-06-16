package country

import "encoding/xml"

type ResponseCountryPhoneCode struct {
	XMLName  xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	SoapBody *SOAPCountryPhoneCodeResponse
}

type SOAPCountryPhoneCodeResponse struct {
	XMLName      xml.Name `xml:"Body"`
	Resp         *ResponseBodyCountryPhoneCode
	FaultDetails *FaultCountryPhoneCode
}

type FaultCountryPhoneCode struct {
	XMLName     xml.Name `xml:"Fault"`
	Faultcode   string   `xml:"faultcode"`
	Faultstring string   `xml:"faultstring"`
}

type ResponseBodyCountryPhoneCode struct {
	XMLName   xml.Name `xml:"CountryIntPhoneCodeResponse"`
	PhoneCode int      `xml:"CountryIntPhoneCodeResult"`
}
