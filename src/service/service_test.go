package service_test

import (
	"model"
	. "service"
	"testing"
)

func Test_SendToHolidayWebService_Input_Canada_Should_Be_48_Days(t *testing.T) {
	countryCode := model.CountryCodeInfo{
		CountryCode: "Canada",
	}
	expected := 48
	actual, _ := SendToHolidayWebService(countryCode)
	if expected != len(actual.Holidays) {
		t.Errorf("expected %d but got it %d", expected, len(actual.Holidays))
	}

}
