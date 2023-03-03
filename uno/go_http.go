package uno

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GoHttp(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	message, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Fail ReadAll: %v", err)
	}
	return string(message)
}

/*type Message struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
}*/

func WeatherHttp() {

	var city string
	var key string = "3f5a76900002399d6164a428566fecfb"

	fmt.Println("Enter your city: ")
	var a, err = fmt.Scan(&city)
	a += 0
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(0)
	}
	var res = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%v&appid=%v", city, key)
	get := GoHttp(res)
	//fmt.Println(get)

	//weth := Message{}
	//mass := []byte(get)

	var weather interface{}

	err = json.Unmarshal([]byte(get), &weather)
	if err != nil {
		fmt.Printf("%v", err)
	}
	//fmt.Printf("%v", weather)

	lon := weather.(map[string]interface{})["coord"].(map[string]interface{})["lon"].(float64)
	lat := weather.(map[string]interface{})["coord"].(map[string]interface{})["lat"].(float64)
	fmt.Printf("\n%v, %v\n", lat, lon)

	var sityres = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&units=metric&appid=%v", lat, lon, key)
	//fmt.Println(sityres)
	resultCity := GoHttp(sityres)
	//fmt.Println(resultCity)

	var outputwether interface{}

	err = json.Unmarshal([]byte(resultCity), &outputwether)
	if err != nil {
		fmt.Printf("%v", err)
	}
	//fmt.Printf("%v", outputwether)

	result := outputwether.(map[string]interface{})["main"].(map[string]interface{})["temp"].(float64)
	fmt.Printf("\nWether in your town %v is equal to %v celsius\n", city, result)
}
