package hackico

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/cryptohazard/coinmarketcap"
)

//Period determines what is the current period
// t is the start of the ICO (startICO)
func Period(t time.Time) (period int, endPeriod time.Time) {
	now := time.Now()
	fmt.Println("Today: ", now)
	delta := now.Sub(t).Hours()
	period = int(math.Ceil(delta / 23))
	endPeriod = t.Add(time.Duration(period*23) * time.Hour)
	//fmt.Println(period)
	//fmt.Println(endPeriod)

	return

}

//TrackPrice keeps asking the current price
func TrackPrice(ch chan float64) {
	var ticker coinmarketcap.Ticker
	var err error

	// Which coins do you want to watch?
	coin := make([]string, 0)
	coin = append(coin, "eos")
	coin = append(coin, "ethereum")

	// Endpoints are updated every 5 minutes, se we use that here
	//for testing
	//period := 30
	//for production
	period := 60 * 5
	clock := time.NewTicker(time.Second * time.Duration(period))

	// Because we are impatient, call it now
	ticker, err = coinmarketcap.GetData(coin)

	// If this is not nil then we encountered a problem, use this to determine
	// what to do next.
	// LastUpdate can be used to determine how stale the data is
	if err != nil {
		fmt.Printf("Error! %s, Last Updated: %s\n", err, ticker.LastUpdate)
	}

	eos := ticker.Coins["eos"]
	ether := ticker.Coins["ethereum"]
	//fmt.Println(eos)
	//fmt.Println(ether)

	price := eos.PriceEUR / ether.PriceEUR
	fmt.Println("eos/eth ", price)
	ch <- price
	//add date and timestamp
	fmt.Println("eos/eth ", price)

	// Infinite loop so we keep getting prices
	for _ = range clock.C {
		// Get off the channel
		ticker, err = coinmarketcap.GetData(coin)

		// If this is not nil then we encountered a problem, use this to determine
		// what to do next.
		// LastUpdate can be used to determine how stale the data is
		if err != nil {
			fmt.Printf("Error! %s, Last Updated: %s\n", err, ticker.LastUpdate)
		}

		eos = ticker.Coins["eos"]
		ether = ticker.Coins["ethereum"]

		price = eos.PriceEUR / ether.PriceEUR
		fmt.Println("eos/eth ", price)
		ch <- price
		coinmarketcap.PrintData(ticker)

	}

	return
}

//Buy : let's buy some tokens
func Buy(ch chan float64, period int, endPeriod time.Time) {

	//production
	//moneyTime = (blocks average(15) * 5) rounded for margin = 87 sec = 1m27s
	moneyTime := 87
	delta := endPeriod.Sub(time.Now()).Seconds()
	d := int(math.Floor(delta)) - moneyTime + 2 // add 2 just in case
	// easy testing
	//d := 5
	// fancy testing
	//moneyTime := 137
	//delta := endPeriod.Sub(time.Now()).Seconds()
	//d := int(math.Floor(delta)) - moneyTime + 2 - 25200 // - 7 heures
	time.AfterFunc(time.Duration(d)*time.Second, func() {
		buy(ch, period, endPeriod)
	})
	go TrackPrice(ch)
}

// actual buying function after we enter moneyTime
func buy(ch chan float64, period int, endPeriod time.Time) {
	// threshold decide when we get in.
	// 20 % is conservative, 50% is greedy ;-)
	threshold := 25.37
	marketPrice := <-ch
	stats := GetDailyStatistics(strconv.Itoa(period))
	saleCurrentPrice := stats.Price
	percentage := gainPercentage(saleCurrentPrice, marketPrice)
	//TODO put those on logs
	fmt.Println("Current expected percentage is: ", percentage, "%")
	if percentage > threshold {
		go dealer(period)
		fmt.Println("Get rich ... today perhaps")

	} else {
		fmt.Println("Get rich ... another day")

	}

}

// compute the expected gain percentage
func gainPercentage(saleCurrentPrice float64, marketPrice float64) (percentage float64) {
	percentage = 100 * (marketPrice - saleCurrentPrice) / saleCurrentPrice
	return
}
