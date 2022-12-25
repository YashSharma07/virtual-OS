package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func showWeatherApp(w fyne.Window) {


	// accessing response from the api (api calling and storing the response)
	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=delhi&APPID=9b4a2664f855098cbc38f4378740a885")
	if err != nil {
		fmt.Println(err)
	}


	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	weather, err := UnmarshalWelcome(body)
	if err != nil {
		fmt.Println(err)
	}

	heading := canvas.NewText("WEATHER UPDATES LIVE", color.White)
	heading.TextStyle = fyne.TextStyle{Bold: true}

	img := canvas.NewImageFromFile("weather2.jfif")
	img.FillMode = canvas.ImageFillContain

	label1 := canvas.NewText(fmt.Sprintf("Country : %s", weather.Sys.Country), color.White)
	label2 := canvas.NewText(fmt.Sprintf("Wind Speed : %.2f knot", weather.Wind.Speed), color.White)
	label3 := canvas.NewText(fmt.Sprintf("Temprature : %.2f Kelvin", weather.Main.Temp), color.White)
	label4 := canvas.NewText(fmt.Sprintf("Humidity : %d", weather.Main.Humidity), color.White)
	weatherContainer := container.NewCenter(

		container.NewGridWithColumns(2,
			container.NewGridWithColumns(1,
				img,
			),
			container.NewGridWithColumns(1,
				heading,
				label1,
				label2,
				label3,
				label4,
			),
		),
	)

	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, weatherContainer))

	w.Show()
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:


func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int64     `json:"timezone"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Cod        int64     `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64 `json:"deg"`
}
