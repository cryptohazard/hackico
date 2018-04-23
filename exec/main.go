package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cryptohazard/hackico"
	//"hackico"
)

func main() {

	//eos ico start date: "2017-07-01T13:00:00.000Z"
	startICO := time.Date(2017, time.July, 1, 13, 0, 0, 0, time.UTC)
	fmt.Println("Eos daily ico start time: ", startICO.Local())
	//TODO put all the time in local clock!!
	//only relevant for printing right? then done!
	period, endPeriod := hackico.Period(startICO)
	fmt.Println("we are on period ", period)
	fmt.Println("This period finishes at ", endPeriod.Local())
	stats := hackico.GetDailyStatistics(strconv.Itoa(period))
	fmt.Println(*stats)
	ch := make(chan float64, 20)
	go hackico.Buy(ch, period, endPeriod)

	// Infinite loop so we keep getting prices
	for {

		if time.Now().Before(endPeriod) {
			time.Sleep(77 * time.Minute)
			fmt.Println("hello we are on period ", period)
			fmt.Println("This period finishes at ", endPeriod.Local())
		} else {
			period, endPeriod = hackico.Period(startICO)
			fmt.Println("we are on period ", period)
			fmt.Println("This period finishes at ", endPeriod.Local())
			go hackico.Buy(ch, period, endPeriod)
		}

	}

}
