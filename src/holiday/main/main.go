package main

import (
	"holiday/api"
	"holiday/log"
	"holiday/route"
	"holiday/service"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	timeoutSecond, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		timeoutSecond = 30
	}
	timeoutDuration := timeoutSecond * time.Second
	holidayWebServiceURL := os.Getenv("HOLIDAY_WEBSERVICE_URL")
	if holidayWebServiceURL == "" {
		holidayWebServiceURL = "http://www.holidaywebservice.com/HolidayService_v2/HolidayService2.asmx?wsdl"
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
