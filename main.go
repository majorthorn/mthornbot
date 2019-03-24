package main

import (
	"github.com/majorthorn/mthornbot/bot"
	"github.com/majorthorn/mthornbot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		return
	}

	bot.Start()
	<-make(chan struct{})

}
