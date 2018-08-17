package api_test

import (
	. "api"
	"bytes"
	"flow"
	"io/ioutil"
	"net/http/httptest"
	"route"
	"testing"
)

func Test_HolidayHandler_Input_CountryCode_Canada_Should_Be_JSON(t *testing.T) {
	expected, _ := ioutil.ReadFile("./response.json")
	countryCode := []byte(`{"countryCode":"Canada"}`)
	request := httptest.NewRequest("POST", "/v1/holiday", bytes.NewBuffer(countryCode))
	writer := httptest.NewRecorder()
	api := Api{
		Flow: flow.MockGetHoliday,
	}
	testRoute := route.NewRoute(api)
	testRoute.ServeHTTP(writer, request)
	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)
	if string(expected) != string(actual) {
		t.Errorf("expect \n'%s' \nbut got it \n'%s'", expected, actual)
	}

}
