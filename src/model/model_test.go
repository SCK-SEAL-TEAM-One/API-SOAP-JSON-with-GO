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
	holidayInfo := holidaysAvailableResult.ToJSON()
	actual, _ := json.Marshal(holidayInfo)
	if string(expected) != string(actual) {
		t.Errorf("expected \n'%s' but got it \n'%s'", expected, actual)
	}

}
