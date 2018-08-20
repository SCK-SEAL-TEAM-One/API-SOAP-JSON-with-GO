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
	logger := log.LoggerMongo{}
	holidayService := service.HolidayService{
		Logger:          &logger,
		TimeoutDuration: 30 * time.Second,
	}
	api := api.Api{
		HolidayService: holidayService,
	}
	route := route.NewRoute(api)
	route.Run(":" + port)

}
