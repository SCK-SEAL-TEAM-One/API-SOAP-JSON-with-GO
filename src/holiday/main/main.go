package main

import (
	"encoding/json"
	"fmt"
	"holiday/api"
	"holiday/log"
	"holiday/route"
	"holiday/service"
	"io/ioutil"

	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)

type Config struct {
	HolidayWebServiceURL string        `json:"holidayWebServiceURL"`
	TimeoutDuration      time.Duration `json:"timeoutDuration"`
	Port                 string        `json:"port"`
	MongoURL             string        `json:"mongoURL"`
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

	if os.Getenv("MONGO_URL") != "" {
		config.MongoURL = os.Getenv("MONGO_URL")
	}

	DBConnection, err := mgo.Dial(config.MongoURL)

	if err != nil {
		fmt.Println("Cannot connect database ", err.Error())
		return
	}

	logger := log.LoggerMongo{
		Session: DBConnection,
	}

	holidayService := service.HolidayService{
		Logger:               &logger,
		TimeoutDuration:      config.TimeoutDuration,
		HolidayWebServiceURL: config.HolidayWebServiceURL,
	}
	api := api.Api{
		HolidayService: holidayService,
	}
	route := route.NewRoute(api)
	route.Run(":" + config.Port)

}
