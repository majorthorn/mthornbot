package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	//Token : is the Token the bot uses to connect to the server
	Token string
	//BotPrefix : is the prefix of a command that the bot is looking for
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

//ReadConfig : Reads the config file of the bot
func ReadConfig() error {
	fmt.Println("Reading from config file...")

	file, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	fmt.Println("Token Set")

	BotPrefix = config.BotPrefix
	fmt.Printf("BotPrefix set to %v \n", BotPrefix)

	return nil
}
