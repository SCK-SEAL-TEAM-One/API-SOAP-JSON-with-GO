package main

import (
	"fmt"
	"holiday/api"
	"holiday/log"
	"holiday/route"
	"holiday/service"

	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)

const url = "mongodb://localhost:27017"

func main() {
	DBConnection, err := mgo.Dial(url)
	if err != nil {
		fmt.Println("Cannot connect database ", err.Error())
		return
	}
	logger := log.LoggerMongo{
		Session: DBConnection,
	}
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
