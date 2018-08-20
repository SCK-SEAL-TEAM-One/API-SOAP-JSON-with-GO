package api_test

import (
	"bytes"
	"encoding/json"
	. "holiday/api"
	"holiday/model"
	"holiday/route"
	"holiday/service"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func mockGetHoliday(countryCodeInfo model.CountryCodeInfo) (model.HolidayInfo, error) {
	var holidayInfo model.HolidayInfo
	holidays, _ := ioutil.ReadFile("./response.json")
	json.Unmarshal(holidays, &holidayInfo)
	return holidayInfo, nil
}

func Test_MockHolidayHandler_Input_CountryCode_Canada_Should_Be_JSON(t *testing.T) {
	expected, _ := ioutil.ReadFile("./response.json")
	countryCode := []byte(`{"countryCode":"Canada"}`)
	request := httptest.NewRequest("POST", "/v1/holiday", bytes.NewBuffer(countryCode))
	request.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()
	api := Api{
		HolidayService: mockGetHoliday,
	}
	testRoute := route.NewRoute(api)
	testRoute.ServeHTTP(writer, request)
	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)
	if string(expected) != string(actual) {
		t.Errorf("expect \n'%s' \nbut got it \n'%s'", expected, actual)
	}

}

func Test_HolidayHandler_Input_CountryCode_GreatBritain_Should_Be_JSON(t *testing.T) {
	expected, _ := ioutil.ReadFile("./response.json")
	countryCode := []byte(`{"countryCode":"GreatBritain"}`)
	request := httptest.NewRequest("POST", "/v1/holiday", bytes.NewBuffer(countryCode))
	request.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()
	api := Api{
		HolidayService: service.SendToHolidayWebService,
	}
	testRoute := route.NewRoute(api)
	testRoute.ServeHTTP(writer, request)
	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)
	if string(expected) != string(actual) {
		t.Errorf("expect \n'%s' \nbut got it \n'%s'", expected, actual)
	}

}
