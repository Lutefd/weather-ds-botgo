package bot

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

const URL string = "https://api.openweathermap.org/data/2.5/weather?"

type WeatherData struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Name string `json:"name"`
}

func getCurrentWeather(message string) (*discordgo.MessageSend, error) {
	r, _ := regexp.Compile(`\d{5}`)
	cep := r.FindString(message)

	if cep == "" {
		return &discordgo.MessageSend{
			Content: "Não foi possível encontrar a cidade",
		}, nil
	}

	weatherURL := fmt.Sprintf("%scep=%s&units=metric&appid=%s", URL, cep, OpenWeatherToken)

	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(weatherURL)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Desculpe, ocorreu um erro ao buscar o clima",
		}, nil
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var data WeatherData
	json.Unmarshal([]byte(body), &data)
}
