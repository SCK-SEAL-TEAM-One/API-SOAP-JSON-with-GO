package main

import (
	"fmt"
	"holiday/api"
	configuration "holiday/config"
	"holiday/log"
	"holiday/route"
	"holiday/service"
	"time"

	"gopkg.in/mgo.v2"
)

func main() {
	config, err := configuration.Configuration()
	if err != nil {
		fmt.Println(err)
		return
	}

	DBConnection, err := mgo.Dial(config.MongoURL)
	if err != nil {
		fmt.Printf("Can not connect database")
		return
	}
	defer DBConnection.Close()

	logger := log.LoggerMongo{
		Session: DBConnection,
	}
	holidayService := service.HolidayService{
		Logger:               &logger,
		TimeoutDuration:      config.TimeoutDuration * time.Second,
		HolidayWebServiceURL: config.HolidayWebServiceURL,
	}
	api := api.Api{
		HolidayService: holidayService,
	}
	route := route.NewRoute(api)
	route.Run(":" + config.Port)

}
