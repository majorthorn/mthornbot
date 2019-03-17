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
		fmt.Printf("Error creating session: %d", err)
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
		fmt.Printf("Error opening websocket: %d", err)
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
			_, err := s.ChannelMessageSend(m.ChannelID, "pong")
			if err != nil {
				fmt.Printf("Error sending message: %d", err)
			}

		}
	}
	// End Code from Tutorial from mgerb42 on Youtube
}
