package hackico

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//DailyStatistics ...
type DailyStatistics struct {
	ID           int     `json:"id"`
	CreateOnDay  int     `json:"createOnDay"`
	DailyTotal   float64 `json:"dailyTotal"`
	Price        float64 `json:"price"`
	Ends         string  `json:"ends"`
	Period       int     `json:"today"`
	NumberOfDays int     `json:"numberOfDays"`
	Begins       string  `json:"begins"`
}

//FullStatistics ...
type FullStatistics struct {
	Daily map[string]*DailyStatistics
}

const eosStatsURL = "https://api.eos.io/eos-sales-statistic"

// GetStatistics makes the actual http request, parses the JSON and returns
// the data in a struct
func GetStatistics() (FullStatistics, error) {

	resp, err := http.Get(eosStatsURL)
	if err != nil {
		return FullStatistics{}, err
	}

	defer resp.Body.Close()

	ret, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return FullStatistics{}, err
	}

	var array []*DailyStatistics

	if errr := json.Unmarshal(ret, &array); errr != nil {
		log.Fatal(errr)
		return FullStatistics{}, errr
	}

	var res = make(map[string]*DailyStatistics, len(array))
	for index, d := range array {
		res[strconv.Itoa(index)] = d
	}

	full := FullStatistics{Daily: res}
	return full, err
}

//GetDailyStatistics ...
func GetDailyStatistics(day string) (dailyStat *DailyStatistics) {
	full, err := GetStatistics()
	if err != nil {
		log.Fatal(err)
		return &DailyStatistics{}
	}
	return full.Daily[day]

}
