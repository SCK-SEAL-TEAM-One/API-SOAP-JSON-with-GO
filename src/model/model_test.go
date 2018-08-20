package model_test

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"model"
	"testing"
)

func Test_XMLToJSON_Should_Be_JSON(t *testing.T) {
	expected, _ := ioutil.ReadFile("./countrycodeinfo.json")
	xmlfile, _ := ioutil.ReadFile("./result.xml")
	var holidaysAvailableResult model.HolidaysAvailableResult

	xml.Unmarshal(xmlfile, &holidaysAvailableResult)
	holidayInfo := holidaysAvailableResult.ToHolidayInfo()
	actual, _ := json.Marshal(holidayInfo)
	if string(expected) != string(actual) {
		t.Errorf("expected \n'%s' but got it \n'%s'", expected, actual)
	}
}

func Test_ToHolidayAvailableMessage_Should_Be_XML(t *testing.T) {
	expectedXML, _ := ioutil.ReadFile("./holidayCanada.xml")
	jsonData := []byte(`{"countryCode":"Canada"}`)
	var request model.CountryCodeInfo
	json.Unmarshal(jsonData, &request)

	actualHolidayAvailableMessage := request.ToHolidayAvailableMessage()
	actualXML, _ := xml.MarshalIndent(actualHolidayAvailableMessage, "", "\t")
	if string(expectedXML) != string(actualXML) {
		t.Errorf("Expect \n'%s' but got it \n'%s'", expectedXML, actualXML)
	}
}
