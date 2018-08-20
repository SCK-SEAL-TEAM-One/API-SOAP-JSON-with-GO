package service

import (
	"bytes"
	"context"
	"encoding/xml"
	"holiday/log"
	"holiday/model"
	"io/ioutil"
	"net/http"
	"time"
)

type IHolidayService interface {
	SendToHolidayWebService(countryCodeInfo model.CountryCodeInfo) (model.HolidayInfo, error)
}
type HolidayService struct {
	HolidayWebServiceURL string
	TimeoutDuration      time.Duration
	Logger               log.Logger
}

func (hs HolidayService) SendToHolidayWebService(countryCodeInfo model.CountryCodeInfo) (model.HolidayInfo, error) {
	message := countryCodeInfo.ToHolidayAvailableMessage()
	ctx, _ := context.WithTimeout(context.Background(), hs.TimeoutDuration)
	XML, err := xml.Marshal(message)
	if err != nil {
		hs.Logger.Error(err.Error())
		return model.HolidayInfo{}, err
	}
	request, _ := http.NewRequest("POST", hs.HolidayWebServiceURL, bytes.NewBuffer(XML))
	request.Header.Set("Content-Type", "text/xml")
	request = request.WithContext(ctx)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		hs.Logger.Error(err.Error())
		return model.HolidayInfo{}, err
	}
	var holidaysAvailableResult model.HolidaysAvailableResult

	data, _ := ioutil.ReadAll(response.Body)
	xml.Unmarshal(data, &holidaysAvailableResult)
	hs.Logger.Info("SUCCESS")
	return holidaysAvailableResult.ToHolidayInfo(), nil

}
