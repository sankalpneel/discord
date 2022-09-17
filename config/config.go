package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token                string
	BotPrefix            string
	config               *configStruct
	OpenWeatherMapApiKey string
)

type configStruct struct {
	Token                string `json:"Token"`
	BotPrefix            string `json:"BotPrefix"`
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

func ReadConfig() error {
	fmt.Println("Reading Config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	//fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	OpenWeatherMapApiKey = config.OpenWeatherMapApiKey

	return nil
}
