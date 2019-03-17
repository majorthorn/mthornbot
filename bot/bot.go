package bot

import (
	"fmt"
	"strings"

	"github.com/majorthorn/Go-Random/testbed1/config"

	"github.com/bwmarrin/discordgo"
)

var botID string
var goBot *discordgo.Session

//Start : Starts the bot
func Start() {
	fmt.Println("test", discordgo.APIVersion)
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
