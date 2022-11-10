# Intro

We hope that you find this exercise fun and interesting. There are no trick questions; we want to see your solution to a simple problem with well thought-out and well-structured code. We realise that there are a lot of topics in the brief and that you may not have the experience or time to complete them all.

There is no strict time limit on how long you spend on the test, but we recommend you spend no longer than 3 hours on it.

When you've produced something you're proud of, send it to us (see Submission). We may then invite you to an interview. In this interview you'll discuss your work, so it's worth considering how you'd improve the application further, even if you didn't have time to do it.

# The Brief

Initial considerations

Our product team would like you to help us build a new roulette platform. Currently all the different variations of roulette work in different ways, some with more business logic in the front end than is preferred. The product team aspire to have a single roulette platform which they can concentrate their focus on.

For this technical test we would like you to create a roulette API. This will be an API that receives requests from a user, simulates a game of roulette, and returns the results. A front end user interface is not required, neither is any consideration of any services which you might expect to be shared; examples of these might be account verification or game history. If you do want to include something like this in your code, please write against a stub - there is certainly no need to write a full implementation.

We’d like you to consider:

How bets are placed, how a win or a loss is communicated, how winnings will be shown. Does your implementation allow for single bets, colour bets, odd/even bets, etc.?
What API methods would be useful to other teams writing calling code (how you can make the API easy to use, is each method doing what someone else would expect it to do?).
Testing and maintainability - you should consider what testing is appropriate.
Further considerations
The expectation from the product team is that we produce a fully working system as soon as possible, then continue to add features. We hope that you will think about this expectation as you work.

As you add more features you might want to consider how they would be rolled out into production. Would your design easily allow feature toggles to be implemented?
Could you easily load test your system?
Are there certain parts of the system you'd like to monitor? How would you monitor them?
How would you deploy your system to an environment? Why would it be advantageous to automate this progress?



# Completed roulette-api documentation

A fairly simple roulette API that will take requests, simulate a roulette table spin and send a response to the user if the have won or lost.
This also includes calculations of winnings.

# Getting started

I plan to implement either a config file or the use of flags to configure certain elements of the programme such as; time between roulette spins, port number etc.
- Time between spins is set to 1 minute for now to allow time for bets to be placed on the same game.
- Port number default is `:8080`

For now it is as simple as using `go run roulette.go`
This will start the programme, and you will see console messages stating the winning numbers of each game, when a new bet is placed and any errors with the request.

# Placing a bet

Bets are placed via http requests and require 4 specific values to be valid for the programme.

- The only handler available for now is /bet so all requests must start with `localhost:/8080/bet?`
- `n` are numbers that will be bet on. This assumes that the front-end will populate number combinations correctly. Numbers 0-36 are supported.
- `stake` is the stake placed with the bet. This is only integers at the moment and the programme does not support floats.
- `type` is the type of bet placed to work out the multiplier. Again, this assumes the front-end correctly populates the numbers and amount of numbers. There is currently no error checking for this!

A simple request may be `localhost:8080/bet?type=straight&n=10&stake=100` This will place a straight bet on the number 10 with a stake of £100. 
A straight bet has 35:1 odds meaning the profit from a win would be £3500.

# Supported bet types

- `0` 35:1
- `straight` 35:1
- `row` 17:1
- `split` 17:1
- `street` 11:1
- `corner` 8:1
- `basket` 8:1
- `doublestreet` 5:1
- `column` 2:1
- `dozen` 2:1
- `colour` 1:1
- `oddeven` 1:1
- `highlow` 1:1

Please note:
This programme currently only supports single bets per request.

# Future improvements 

Use of json in requests made to api. Ran short on time in this iteration

Ability to place multiple bets at once either through batching or structured json
