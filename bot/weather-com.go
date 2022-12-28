package bot

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
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

func GetCurrentWeather(message string) (*discordgo.
	MessageSend, error) {
	r, _ := regexp.Compile(`\S{3,50}`)
	cep := r.FindString(message)

	if cep == "" {

		return &discordgo.MessageSend{
			Content: "Não foi possível encontrar a cidade",
		}, nil
	}

	weatherURL := fmt.Sprintf(
		"%sq=%s&units=metric&appid=%s", URL, cep,
		OpenWeatherToken)

	fmt.Println(weatherURL)

	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(weatherURL)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Não foi possível estabelecer uma conexão com o servidor",
		}, nil
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var data WeatherData
	json.Unmarshal([]byte(body), &data)

	city := data.Name
	conditions := data.Weather[0].Description
	temp := strconv.FormatFloat(data.Main.Temp,
		'f', 2, 64)
	humidity := strconv.Itoa(int(data.Main.Humidity))
	wind := strconv.FormatFloat(data.Wind.Speed, 'f', 2, 64)

	embed := &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Clima atual",
			Description: "Clima atual em " + city,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Condições",
					Value:  conditions,
					Inline: true,
				},
				{
					Name:   "Temperatura",
					Value:  temp,
					Inline: true,
				},
				{
					Name:   "Umidade",
					Value:  humidity,
					Inline: true,
				},
				{
					Name:   "Vento",
					Value:  wind + "km/h",
					Inline: true,
				},
			},
		},
	}
	return embed, nil
}
