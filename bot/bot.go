package bot

import "fmt"

var (
	OpenWeatherToken string
	BotToken         string
)

func Run() {
	fmt.Print("Chaves fcnionando: ", OpenWeatherToken, BotToken) // ...
}
