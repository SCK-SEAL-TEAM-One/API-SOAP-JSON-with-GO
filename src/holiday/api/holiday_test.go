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
	"time"
)

type mockHolidayService struct {
}

type mockLogger struct {
	ErrorCount int
	Success    int
}

func (mocklog *mockLogger) Error(s string) bool {
	mocklog.ErrorCount++
	return true
}

func (mocklog *mockLogger) Info(s string) bool {
	mocklog.Success++
	return true
}

func (mhs mockHolidayService) SendToHolidayWebService(countryCodeInfo model.CountryCodeInfo) (model.HolidayInfo, error) {
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
	mockGetHoliday := mockHolidayService{}
	api := Api{
		HolidayService: &mockGetHoliday,
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
	holiday := service.HolidayService{
		Logger:               &mockLogger{},
		TimeoutDuration:      30 * time.Second,
		HolidayWebServiceURL: "http://www.holidaywebservice.com/HolidayService_v2/HolidayService2.asmx?wsdl",
	}
	api := Api{
		HolidayService: &holiday,
	}
	testRoute := route.NewRoute(api)
	testRoute.ServeHTTP(writer, request)
	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)
	if string(expected) != string(actual) {
		t.Errorf("expect \n'%s' \nbut got it \n'%s'", expected, actual)
	}

}
