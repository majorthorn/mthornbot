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
		if m.Author.ID == botID {
			return
		}

		if m.Content == "!ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong")

		}
	}
	// End Code from Tutorial from mgerb42 on Youtube
}
