package continent

import "encoding/xml"

type ResponseContinentInfo struct {
	XMLName  xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	SoapBody *SOAPContinentInfoResponse
}

type SOAPContinentInfoResponse struct {
	XMLName      xml.Name `xml:"Body"`
	Resp         *ResponseBodyContinentInfo
	FaultDetails *FaultContinent
}

type FaultContinent struct {
	XMLName     xml.Name `xml:"Fault"`
	Faultcode   string   `xml:"faultcode"`
	Faultstring string   `xml:"faultstring"`
}

type ResponseBodyContinentInfo struct {
	XMLName  xml.Name `xml:"ListOfContinentsByNameResponse"`
	Response *BodyContinentInfo
}

type BodyContinentInfo struct {
	XMLName    xml.Name    `xml:"ListOfContinentsByNameResult"`
	Continents []Continent `xml:"tContinent"`
}

type Continent struct {
	Code string `xml:"sCode"`
	Name string `xml:"sName"`
}
