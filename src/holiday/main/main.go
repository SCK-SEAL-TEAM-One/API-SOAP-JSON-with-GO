package main

import (
	"fmt"
	"holiday/api"
	"holiday/log"
	"holiday/route"
	"holiday/service"

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
	holidayService := service.HolidayService{
		Logger: &logger,
	}
	api := api.Api{
		HolidayService: holidayService,
	}
	route := route.NewRoute(api)
	route.Run(":3000")

}
