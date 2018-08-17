package flow

import (
	"encoding/json"
	"io/ioutil"
	"model"
)

func MockGetHoliday(countryCodeInfo model.CountryCodeInfo) model.HolidayInfo {
	var holidayInfo model.HolidayInfo
	holidays, _ := ioutil.ReadFile("./response.json")
	json.Unmarshal(holidays, &holidayInfo)
	return holidayInfo
}

func GetHoliday(countryCodeInfo model.CountryCodeInfo) model.HolidayInfo {
	var holidayInfo model.HolidayInfo

	return holidayInfo
}
