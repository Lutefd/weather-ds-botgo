package bot

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	OpenWeatherToken string
	BotToken         string
)

func Run() {

	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(newMessage)

	discord.Open()

	defer discord.Close()
}
