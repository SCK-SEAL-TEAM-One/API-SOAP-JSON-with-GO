package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type Config struct {
	HolidayWebServiceURL string        `json:"holidayWebServiceURL"`
	TimeoutDuration      time.Duration `json:"timeoutDuration"`
	Port                 string        `json:"port"`
	MongoURL             string        `json:"mongoURL"`
}

func Configuration() (Config, error) {
	var config Config
	var environment string

	if os.Getenv("ENV") == "" {
		environment = "development"
	} else {
		environment = os.Getenv("ENV")
	}

	configFile, err := ioutil.ReadFile("./configs/" + environment + ".json")
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return config, err
	}

	if os.Getenv("TIMEOUT") != "" {
		config.TimeoutDuration, _ = time.ParseDuration(os.Getenv("TIMEOUT"))
	}

	if os.Getenv("PORT") != "" {
		config.Port = os.Getenv("PORT")
	}

	holidayWebserviceURL := os.Getenv("HOLIDAY_WEBSERVICE_URL")
	if holidayWebserviceURL != "" {
		config.HolidayWebServiceURL = holidayWebserviceURL
	}

	return config, nil
}
