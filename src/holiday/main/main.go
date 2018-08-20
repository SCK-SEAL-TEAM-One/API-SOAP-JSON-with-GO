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
	environment := os.Getenv("ENV") //development, production
	if environment == "" {
		environment = "development"
	}
	file, _ := ioutil.ReadFile("./configs/" + environment + ".json")
	var config Config
	json.Unmarshal(file, &config)

	port := os.Getenv("PORT")
	if port == "" {
		port = config.Port
	}

	timeoutSecond, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		timeoutSecond = config.TimeoutDuration
	}
	timeoutDuration := timeoutSecond * time.Second
	holidayWebServiceURL := os.Getenv("HOLIDAY_WEBSERVICE_URL")
	if holidayWebServiceURL == "" {
		holidayWebServiceURL = config.HolidayWebServiceURL
	}

	logger := log.LoggerMongo{}
	holidayService := service.HolidayService{
		Logger:               &logger,
		TimeoutDuration:      timeoutDuration,
		HolidayWebServiceURL: holidayWebServiceURL,
	}
	api := api.Api{
		HolidayService: holidayService,
	}
	route := route.NewRoute(api)
	route.Run(":" + port)

}
