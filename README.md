# Hackico
This (*Not so quick and not so dirty*) project help you profit from market discrepancy in ICOs like EOS. EOS has a new batch of tokens sold everyday, for 350 period of 23 Hours. The token can already be traded, while the daily ICO is ongoing.

This create an opportunity to profit: the daily ICO price won't always be aligned with the market price due to various reason (FOMO, finishing ICO time compared to investors timezone, loss of popularity, spam attacks...).
We use two strategies:
* we enter the ICO *before* the last minute. The price is decided at the end so most of people join at the end. You can see nice chart here: http://eosscan.io/
* we choose an expected return before we get in. For instance you can choose to only enter, if at the time of calculation (as explained before) the expected return is 50% because you are greedy or 15% if you are prudent. Note: you will get less because more people will join after you.

No need to restart the program everyday. Just make sure the transaction to buy is ready/unused(see below for more details.)
# Installation
You need to have a working go(lang) environment in version >=1.9 and clone this repository. I can provide executable if there are requests.

You first need to get the only dependency:
```
$ go get github.com/cryptohazard/coinmarketcap
```

Now you can go in the ```exec``` folder and build the executable( I added the ```-o``` option to change the name):
```
$ cd exec/
$ go build -o hackico main.go
$ ./hackico
```
# Configuration

## Threshold percentage
Choose the thresold at which you want to get by changing it in ```hackico.go```:
```
threshold := 25.37
```

## Time to get in
We actually get in the ICO **87 seconds** before the end. Just change this value if you want to get earlier or later in the file ```hackico.go```:
```
//moneyTime = (blocks average(15) * 5) rounded for margin = 87 sec = 1m27s
moneyTime := 87
```

## Ethereum buying transaction
For now ;-), this program doesn't create the transaction to buy the EOS tokens. So you need to make your own transaction with your favorite ethereum client (geth, parity, myetherwallet, metamask ...).
Just put the raw transaction in a file called **tx** located with the executable. You can see an example of a random transaction already sent during the ICO in the provided **tx** file.

The program uses the *eth_sendRawTransaction* from [etherscan broadcast API](https://etherscan.io/pushTx) to broadcast the transaction.

## Tips
To get a fast confirmation, put the gas limit as low as you can and the gas price higher.

To claim your tokens, use ```claim(uint day)``` instead of ```claimAll()```. It will use less gas ;-).

# Expected logs
When you start the program, you should see information on daily ICO, the period, a countdown, the price:
```
Eos daily ico start time:  2017-07-01 15:00:00 +0200 CEST
Today:  2018-04-25 23:42:58.629965922 +0200 CEST m=+0.034309296
we are on period  312
This period finishes at  2018-04-26 15:00:00 +0200 CEST
Countdown:  15h17m0s left
Current price: 1 eos =  0.02341794846431574 eth
```

If the transaction was broadcasted, you should see:
```
Get rich ... today perhaps
```
You still need to check if you got in at the right period. Due to Ethereum network or your chosen fees, your transaction can get through in the ***next*** period.

**You also need to update the transaction for the next periods**.

If the expected returns are lower than he threshold percentage you put, the transaction is not broadcasted and the program is ready for the next period. You should see:
```
Get rich ... another day
```


# Responsibility/Disclaimer

You are **responsible** for your own transaction. My program will just broadcast it. If you make a mistake in the contract, gas price, gas limit, amount..., I can not do anything about it.

# Eventual roadmap

* clean the code
* add options
* add wallet
* send ether
* wait until it's accepted
* retrieve the eos token from contract
* initiate deal with shapeshift or other exchange API
* send eos
* get ether
* compute ROI


# Tips
If you want because you got lucky with this program:
0xFAcE0f50707B250203FF463fE6C40322ddC459E6

Feel free to open issues, make pull requests...
