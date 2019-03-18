package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/majorthorn/mthornbot/config"
)

var botID string
var goBot *discordgo.Session

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
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		prefix := config.BotPrefix

		switch {
		case m.Author.ID == botID:
			fmt.Println("Bot said ", m.Content)
			return
		case m.Content == prefix+"ping":
			_, err := s.ChannelMessageSend(m.ChannelID, "pong")
			if err != nil {
				fmt.Println(err.Error())
			}
		case m.Content == prefix+"disconnect":
			_, err := s.ChannelMessageSend(m.ChannelID, "Not Implemented yet")
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
