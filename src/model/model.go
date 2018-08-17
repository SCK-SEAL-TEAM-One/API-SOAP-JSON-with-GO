package model

import (
	"encoding/xml"
)

type HolidayAvailableMessage struct {
	XMLName     xml.Name `xml:"soapenv:Envelope"`
	CountryCode string   `xml:"soapenv:Body>hs:GetHolidaysAvailable>hs:countryCode"`
	Namespace   string   `xml:"xmlns:soapenv,attr"`
	NamespaceHs string   `xml:"xmlns:hs,attr"`
}
type CountryCodeInfo struct {
	CountryCode string `json:"countryCode"`
}

func (countryCode CountryCodeInfo) JsonToXML() []byte {
	xmlHolidayAvailableMessage := HolidayAvailableMessage{
		Namespace:   "http://schemas.xmlsoap.org/soap/envelope/",
		NamespaceHs: "http://www.holidaywebservice.com/HolidayService_v2/",
		CountryCode: countryCode.CountryCode,
	}
	xml, _ := xml.MarshalIndent(xmlHolidayAvailableMessage, "", "\t")
	return xml
}
