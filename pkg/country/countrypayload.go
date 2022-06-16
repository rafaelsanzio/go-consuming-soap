package country

import "encoding/xml"

type ResponseCountryInfo struct {
	XMLName  xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	SoapBody *SOAPCountryInfoResponse
}

type SOAPCountryInfoResponse struct {
	XMLName      xml.Name `xml:"Body"`
	Resp         *ResponseBodyCountryInfo
	FaultDetails *FaultCountry
}

type FaultCountry struct {
	XMLName     xml.Name `xml:"Fault"`
	Faultcode   string   `xml:"faultcode"`
	Faultstring string   `xml:"faultstring"`
}

type ResponseBodyCountryInfo struct {
	XMLName  xml.Name `xml:"ListOfCountryNamesByNameResponse"`
	Response *BodyCountryInfo
}

type BodyCountryInfo struct {
	XMLName   xml.Name  `xml:"ListOfCountryNamesByNameResult"`
	Countries []Country `xml:"tCountryCodeAndName"`
}

type Country struct {
	Code string `xml:"sISOCode"`
	Name string `xml:"sName"`
}
