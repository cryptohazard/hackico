package hackico

import (
	//"encoding/json"
	//"fmt"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"time"
)

/*dealer will do a lot of things
* send ether
* wait until it's accepted
* retrieve the eos token from contract
* initiate deal with shapeshift or other exchange API
* send eos
* get ether
* compute ROI
* but for now it just broadcast the tx
 */
func dealer(period int) {
	broadcastor(transactor())

}

//transactor is private just in case
//tx == raw transaction
func transactor() (tx string) {
	txFile, err := os.Open("tx")
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}
	defer txFile.Close()
	scanner := bufio.NewScanner(txFile)

	//tx is just a line
	scanner.Scan()
	tx = scanner.Text()
	//fmt.Println(line)

	return

}

func broadcastor(tx string) {
	url := "https://api.etherscan.io/api?module=proxy&action=eth_sendRawTransaction&hex=" + tx
	resp, err := http.Get(url)
	if err != nil {
		// just try again if not give up for now
		// better: fallback on other method... one day
		resp, err = http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer resp.Body.Close()

	respCode, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(respCode)
	fmt.Println("bought!!")

}
