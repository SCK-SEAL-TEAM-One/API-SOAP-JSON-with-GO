package service_test

import (
	"holiday/model"
	. "holiday/service"
	"testing"
	"time"
)

func Test_SendToHolidayWebService_Input_Canada_Should_Be_48_Days(t *testing.T) {
	countryCode := model.CountryCodeInfo{
		CountryCode: "Canada",
	}
	expected := 48
	logger := mockLogger{}
	holidayService := HolidayService{
		Logger:          &logger,
		TimeoutDuration: 30 * time.Second,
	}
	actual, _ := holidayService.SendToHolidayWebService(countryCode)

	if expected != len(actual.Holidays) {
		t.Errorf("expected %d but got it %d", expected, len(actual.Holidays))
	}
	if logger.Success != 1 {
		t.Errorf("expected logger call %d but got it %d", 1, logger.Success)
	}

}
