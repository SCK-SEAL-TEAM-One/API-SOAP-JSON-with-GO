package model

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func Test_JsonToXML_Should_Be_XML(t *testing.T) {
	expectedXML, _ := ioutil.ReadFile("./holidayCanada.xml")
	jsonData := []byte(`{"countryCode":"Canada"}`)
	var request CountryCodeInfo
	json.Unmarshal(jsonData, &request)

	actualXML := request.JsonToXML()
	if string(expectedXML) != string(actualXML) {
		t.Errorf("Expect \n'%s' but got it \n'%s'", expectedXML, actualXML)
	}
}
