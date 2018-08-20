package service

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"model"
	"net/http"
	"time"
)

const timeoutDuration = 30 * time.Second
const holidayWebServiceURL = "http://www.holidaywebservice.com/HolidayService_v2/HolidayService2.asmx?wsdl"

func GetHoliday(countryCodeInfo model.CountryCodeInfo) model.HolidayInfo {
	var holidayInfo model.HolidayInfo

	return holidayInfo
}

func SendToHolidayWebService(message model.HolidayAvailableMessage) (model.HolidayInfo, error) {
	ctx, _ := context.WithTimeout(context.Background(), timeoutDuration)
	XML, err := xml.Marshal(message)
	if err != nil {
		return model.HolidayInfo{}, err
	}
	request, _ := http.NewRequest("POST", holidayWebServiceURL, bytes.NewBuffer(XML))
	request.Header.Set("Content-Type", "text/xml")
	request = request.WithContext(ctx)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return model.HolidayInfo{}, err
	}
	var holidaysAvailableResult model.HolidaysAvailableResult

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("%s", data)
	xml.Unmarshal(data, &holidaysAvailableResult)

	return holidaysAvailableResult.ToHolidayInfo(), nil

}
