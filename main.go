package main

import "log"

func main() {

	botToken, ok := os.LookupEbv("BOT_TOKEN")
	if !ok {
		log.Fatal("BOT_TOKEN não encontrado")
	}
	openWeatherToken, ok := os.LookupEbv("OPEN_WEATHER_TOKEN")
	if !ok {
		log.Fatal("OPEN_WEATHER_TOKEN não encontrado")
	}

	bot.BotToken = botToken
	bot.OpenWeatherToken = openWeatherToken
	bot.Run()

}
