package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
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

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == discord.State.User.ID {
		return
	}

	switch {
	case strings.Contains(message.Content, "clima"):
		discord.ChannelMessageSend(message.ChannelID, "Eu posso te ajudar com isso!")
	case strings.Contains(message.Content, "bot"):
		discord.ChannelMessageSend(message.ChannelID, "Olá! Eu sou o WeatherDSBot, um bot que te ajuda a saber o clima de qualquer lugar do mundo!")

	}

}
