package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
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

	fmt.Print("Bot foi iniciado com sucesso")

	channel := make(chan os.Signal, 1)

	signal.Notify(channel, os.Interrupt)
	<-channel
}
