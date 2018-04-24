# Hackico
This (*Not so quick and not so dirty*) project help you profit from market discrepancy in ICOs like EOS. EOS has a new batch of tokens sold everyday, for 350 period of 23 Hours. The token can already be traded, while the daily ICO is ongoing.
This create an opportunity to profit: the daily ICO price won't always be aligned with the market price due to various reason (FOMO, finishing ICO time compared to investors timezone, loss of popularity, spam attacks...).
We use two strategies:
* we enter the ICO *before* the last minute. The price is decided at the end so most of people join at the end. You can see nice chart here: http://eosscan.io/
* we choose an expected return before we get in. For instance you can choose to only enter, if at the time of calculation (as explained before) the expected return is 50% because you are greedy or 15% if you are prudent. Note: you will get less because more people will join after you.

# Installation
You need to have a working go environment and clone this repository. I can provide executable if there are requests.
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

## Ethereum buying transaction
For now ;-), this program doesn't create the transaction to buy the EOS tokens. So you need to make your own transaction with your favorite ethereum client (geth, parity, myetherwallet, metamask ...).
Just put the raw transaction in a file called **tx** located with the executable.
*Small advice: to get a fast confirmation, put the gas limit as low as you can and the gas price higher.*

# Responsibility

You are **responsible** for your own transaction. My program will just broadcast it. If you make a mistake in the contract, gas price, gas limit, amount, I can do anything about it.

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
