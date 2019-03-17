package main

import (
	"fmt"

	"github.com/majorthorn/mthornbot/bot"
	"github.com/majorthorn/mthornbot/config"
)

//Start code from tutorial by mgerb42 on youtube
func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()
	<-make(chan struct{})

}