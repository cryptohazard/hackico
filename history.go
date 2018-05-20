package hackico

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"time"

	"github.com/fatih/color"
)

type CryptocompareApiResponse struct {
	Response   string `json:"Response"`
	Type       int    `json:"Type"`
	Aggregated bool   `json:"Aggregated"`
	Data       []struct {
		Time       int64     `json:"time"`
		Close      float64 `json:"close"`
		High       float64 `json:"high"`
		Low        float64 `json:"low"`
		Open       float64 `json:"open"`
		Volumefrom float64 `json:"volumefrom"`
		Volumeto   float64 `json:"volumeto"`
	} `json:"Data"`
	TimeTo            int  `json:"TimeTo"`
	TimeFrom          int  `json:"TimeFrom"`
	FirstValueInArray bool `json:"FirstValueInArray"`
	ConversionType    struct {
		Type             string `json:"type"`
		ConversionSymbol string `json:"conversionSymbol"`
	} `json:"ConversionType"`
}

const url = "https://min-api.cryptocompare.com/data/histohour?fsym=EOS&tsym=ETH&limit=2000&aggregate=1&"

func History(history uint, period int){
	if history >80 {
		fmt.Println("You can only print up to the last 80 days!")
		history = 80
	}

	color.Cyan("Starting history")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	var api *CryptocompareApiResponse
	if errr := json.Unmarshal(body, &api); errr != nil {
		log.Fatal(errr)
	}
	stats := GetLastDailyStatistics(history, period)

	fmt.Println(" Period\t|  Percentage\t| Time")
	fmt.Println("--------------------------------------------------")
	var percentage float64
	for _,dailyICO := range stats{
		t,_:= time.Parse(time.RFC3339,dailyICO.Ends)
		tUnix := t.Unix()
		for _, histoHour := range api.Data{
			if histoHour.Time == tUnix{
				percentage = gainPercentage(dailyICO.Price, histoHour.Close)
				str := "   "+ strconv.Itoa(dailyICO.ID) +
					"\t|    " + strconv.FormatFloat(percentage, 'f', 2, 64) +
					"\t|" + t.Local().String()
				if percentage > 0.0 {
					color.Blue(str)
				} else{
					color.Red(str)
				}
			}
		}
	}
}
