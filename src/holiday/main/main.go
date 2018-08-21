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

	"gopkg.in/mgo.v2"
)

type Config struct {
	HolidayWebServiceURL string        `json:"holidayWebServiceURL"`
	TimeoutDuration      time.Duration `json:"timeoutDuration"`
	Port                 string        `json:"port"`
	MongoURL             string        `json:"mongoURL"`
}

func main() {
	config, err := Configuration()
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

func Configuration() (Config, error) {
	var config Config
	environment := os.Getenv("ENV")
	if environment == "" {
		environment = "development"
	}
	configFile, err := ioutil.ReadFile("./configs/" + environment + ".json")
	if err != nil {
		return config, err
	}

	json.Unmarshal(configFile, &config)

	timeOut := os.Getenv("TIMEOUT")
	if timeOut != "" {
		config.TimeoutDuration, _ = time.ParseDuration(timeOut)
	}

	port := os.Getenv("PORT")
	if port != "" {
		config.Port = port
	}

	holidayWebserviceURL := os.Getenv("HOLIDAY_WEBSERVICE_URL")
	if holidayWebserviceURL != "" {
		config.HolidayWebServiceURL = holidayWebserviceURL
	}
	return config, nil
}
