# roulette-api

A fairly simple roulette API that will take requests, simulate a roulette table spin and send a response to the user if the have won or lost.
This also includes calculations of winnings.

# Getting started

I plan to implement either a config file or the use of flags to configure certain elements of the programme such as; time between roulette spins, port number etc.
- Time between spins is set to 1 minute for now to allow time for bets to be placed on the same game.
- Port number default is `:8080`

For now it is as simple as using `go run roulette.go`
This will start the programme and you will see console messages stating the winning numbers of each game and also when a new bet is placed.

# Placing a bet

Bets are placed via http requests and require 4 specific values to be valid for the programme.

- The only handler available for now is /bet so all requests must start with `localhost:/8080/bet?`
- `n` are numbers that will be bet on. This assumes that the front-end will populate number combinations correctly. Numbers 0-36 are supported.
- `stake` is the stake placed with the bet. This is only integers at the moment and the programme does not support floats.
- `type` is the type of bet placed to work out the multiplier. Again, this assumes the front-end correctly populates the numbers and amount of numbers. There is currently no error checking for this!

A simple request may be `localhost:8080/bet?type=straight&n=10&stake=100` This will place a straight bet on the number 10 with a stake of £100. 
A straight bet has 35:1 odds meaning the profit from a win would be £3500.

# Supported bet types

- `0`
- `straight`
- `row`
- `split`
- `street`
- `corner`
- `basket`
- `doublestreet`
- `column`
- `dozen`
- `colour`
- `oddeven`
- `highlow`

Please note:
This programme currently only supports single bets.
