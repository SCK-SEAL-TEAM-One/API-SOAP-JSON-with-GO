package model

import (
	"encoding/xml"
)

type HolidaysAvailableResult struct {
	XMLName     xml.Name      `xml:"Envelope"`
	HolidayCode []HolidayCode `xml:"Body>GetHolidaysAvailableResponse>GetHolidaysAvailableResult>HolidayCode"`
}

type HolidayCode struct {
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
}

func (har HolidaysAvailableResult) ToJSON() HolidayInfo {
	holiday := make([]Holiday, len(har.HolidayCode))
	for index := range har.HolidayCode {
		holiday[index] = Holiday{
			Code:        har.HolidayCode[index].Code,
			Description: har.HolidayCode[index].Description,
		}
	}
	return HolidayInfo{
		Holidays: holiday,
	}
}
