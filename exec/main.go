package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/cryptohazard/hackico"
	"github.com/fatih/color"
	//"hackico"
)

func main() {
	history := flag.Uint("h", 0, "How many days of history do you want to print? Max: 80!")
	threshold := flag.Float64("t", 9.37, "threshold to decide when we get in. 5 % is conservative, 20% is greedy ;-)")
	flag.Parse()

	//eos ico start date: "2017-07-01T13:00:00.000Z"
	startICO := time.Date(2017, time.July, 1, 13, 0, 0, 0, time.UTC)
	color.Cyan("Eos daily ico start time: %s", startICO.Local())
	period, endPeriod := hackico.Period(startICO)
	fmt.Println("we are on period ", period)

	if *history != 0 {
		hackico.History(*history, period)
		return
	}

	fmt.Println("This period finishes at ", endPeriod.Local())
	fmt.Println("Countdown: ", endPeriod.Sub(time.Now()).Round(time.Minute), "left")
	ch := make(chan float64, 20)
	go hackico.Buy(ch, period, endPeriod, *threshold)

	// Infinite loop so we keep getting prices
	for {

		if time.Now().Before(endPeriod) {
			time.Sleep(7 * time.Minute)
			fmt.Println("Hello we are on period ", period)
			fmt.Println("This period finishes at ", endPeriod.Local())
			fmt.Println("Countdown: ", endPeriod.Sub(time.Now()).Round(time.Minute), "left")
		} else {
			period, endPeriod = hackico.Period(startICO)
			fmt.Println("we are on period ", period)
			fmt.Println("This period finishes at ", endPeriod.Local())
			fmt.Println("Countdown: ", endPeriod.Sub(time.Now()).Round(time.Minute), "left")
			go hackico.Buy(ch, period, endPeriod, *threshold)
		}

	}

}
