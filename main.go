package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	//"github.com/majorthorn/mthornbot/bot"
	"github.com/majorthorn/mthornbot/config"
)

//Start code from tutorial by mgerb42 on youtube
func main() {
	err := config.ReadConfig()
	if err != nil {
		return
	}

	Start()
	<-make(chan struct{})

}

// Bot.go start
// this is only for testing purposes and reducing the amoutn of files
// in this project for now.

var (
	botID string
	goBot *discordgo.Session
)

//Start : Starts the bot
func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	botID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	message := m.Content
	authorID := m.Author.ID
	cmdPrefix := config.BotPrefix

	if strings.HasPrefix(message, cmdPrefix) {
		switch {
		case authorID == botID:
			fmt.Println("Bot said ", message)
			return
		case message == cmdPrefix+"ping":
			_, err := s.ChannelMessageSend(m.ChannelID, "pong")
			if err != nil {
				fmt.Println(err.Error())
			}
		case message == cmdPrefix+"disconnect":
			_, err := s.ChannelMessageSend(m.ChannelID, "Not Implemented yet")
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
