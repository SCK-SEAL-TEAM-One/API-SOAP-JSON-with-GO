package model

type CountryCodeInfo struct {
	CountryCode string `json:"countryCode"`
}

type HolidayInfo struct {
	Holiday []Holiday `json:"holidays"`
}

type Holiday struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
