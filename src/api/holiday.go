package api

import (
	"encoding/json"
	"model"
	"net/http"
)

type Api struct {
	Flow func(model.CountryCodeInfo) model.HolidayInfo
}

func (api Api) HolidayHandler(writer http.ResponseWriter, request *http.Request) {
	var countryCodeInfo model.CountryCodeInfo
	err := json.NewDecoder(request.Body).Decode(&countryCodeInfo)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	holidays := api.Flow(countryCodeInfo)

	err = json.NewEncoder(writer).Encode(holidays)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
