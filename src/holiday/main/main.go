package main

import (
	"holiday/api"
	"holiday/log"
	"holiday/route"
	"holiday/service"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	logger := log.LoggerMongo{}
	holidayService := service.HolidayService{
		Logger: &logger,
	}
	api := api.Api{
		HolidayService: holidayService,
	}
	route := route.NewRoute(api)
	route.Run(":" + port)

}
