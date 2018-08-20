package main

import (
	"encoding/json"
	"holiday/api"
	"holiday/log"
	"holiday/route"
	"holiday/service"
	"io/ioutil"
	"os"
	"time"
)

type Config struct {
	HolidayWebServiceURL string        `json:"holidayWebServiceURL"`
	TimeoutDuration      time.Duration `json:"timeoutDuration"`
	Port                 string        `json:"port"`
}

func main() {
	environment := os.Getenv("ENV")
	if environment == "" {
		environment = "development"
	}
	file, _ := ioutil.ReadFile("./configs/" + environment + ".json")
	var config Config
	json.Unmarshal(file, &config)

	if os.Getenv("PORT") != "" {
		config.Port = os.Getenv("PORT")
	}
	if os.Getenv("TIMEOUT") != "" {
		config.TimeoutDuration, _ = time.ParseDuration(os.Getenv("TIMEOUT"))
	}
	config.TimeoutDuration = config.TimeoutDuration * time.Second

	if os.Getenv("HOLIDAY_WEBSERVICE_URL") != "" {
		config.HolidayWebServiceURL = os.Getenv("HOLIDAY_WEBSERVICE_URL")
	}

	logger := log.LoggerMongo{}
	holidayService := service.HolidayService{
		Logger:               &logger,
		TimeoutDuration:      config.TimeoutDuration,
		HolidayWebServiceURL: config.HolidayWebServiceURL,
	}
	api := api.Api{
		HolidayService: holidayService,
	}
	route := route.NewRoute(api)
	route.Run(":" + port)

}
