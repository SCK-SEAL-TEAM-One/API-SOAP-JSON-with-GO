package api

import (
	"encoding/json"
	"flow"
	"model"
	"net/http"
)

func HolidayHandler(writer http.ResponseWriter, request *http.Request) {
	var countryCodeInfo model.CountryCodeInfo
	err := json.NewDecoder(request.Body).Decode(&countryCodeInfo)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	holidays := flow.GetHoliday(countryCodeInfo)

	err = json.NewEncoder(writer).Encode(holidays)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
