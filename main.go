package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "http://www.weather.com.cn/data/sk/101190101.html"
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("get weather.com error: %v", err)
	}

	data, _ := ioutil.ReadAll(resp.Body)

	response := Response{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Printf("unmarshal json error: %v", err)
	}

	//	city := response.WeatherInfo.City
	temp := response.WeatherInfo.Temp
	sd := response.WeatherInfo.SD
	rain := response.WeatherInfo.Rain
	time := response.WeatherInfo.Time

	fmt.Printf("%s: %sC %s %s %s\n", "Nanjing", temp, sd, rain, time)
}
