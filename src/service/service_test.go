package service_test

import (
	"model"
	. "service"
	"testing"
)

func Test_SendToHolidayWebService_Input_Canada_Should_Be_48_Days(t *testing.T) {
	message := model.HolidayAvailableMessage{
		Namespace:   "http://schemas.xmlsoap.org/soap/envelope/",
		NamespaceHs: "http://www.holidaywebservice.com/HolidayService_v2/",
		CountryCode: "Canada",
	}
	expected := 48
	actual, _ := SendToHolidayWebService(message)
	if expected != len(actual.Holidays) {
		t.Errorf("expected %d but got it %d", expected, len(actual.Holidays))
	}

}
