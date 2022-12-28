package main

import (
	"github.com/Lutefd/weather-ds-botgo/bot"
	"log"
	"os"
)

func main() {

	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("BOT_TOKEN não encontrado")
	}
	openWeatherToken, ok := os.LookupEnv("OPEN_WEATHER_TOKEN")
	if !ok {
		log.Fatal("OPEN_WEATHER_TOKEN não encontrado")
	}

	bot.BotToken = botToken
	bot.OpenWeatherToken = openWeatherToken
	bot.Run()

}
