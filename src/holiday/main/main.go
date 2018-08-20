package main

import (
	"holiday/api"
	"holiday/log"
	"holiday/route"
	"holiday/service"
)

func main() {
	logger := log.LoggerMongo{}
	holidayService := service.HolidayService{
		Logger: &logger,
	}
	api := api.Api{
		HolidayService: holidayService,
	}
	route := route.NewRoute(api)
	route.Run(":3000")

}
