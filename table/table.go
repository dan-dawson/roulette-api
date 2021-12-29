package table

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	participants        []Betslip
	TableRequestChannel = make(chan TableRequest)
	seed                rand.Source
	winningNumber       int
)

const (
	AddBet = iota
	ClearParticipants
	NotifyParticipants
)

type TableRequest struct {
	Cmd     int
	BetSlip Betslip
}

// TableStart will build required variables and start the stateWorker.
func TableStart() {
	participants = make([]Betslip, 0)
	seed = rand.NewSource(time.Now().UnixNano())

	tableStateWorker(TableRequestChannel)
}

// tableStateWorker will ensure all manipulation of the participants slice is done in a thread safe manner.
func tableStateWorker(tableReqCh chan TableRequest) {

	// This ticker will be used to give time for users to place bets. In the real world, the duration would be set via a config file or flag.
	ticker := time.NewTimer(1 * time.Minute)

	for {
		select {
		case <-ticker.C:
			go spinTheWheel()
		case req := <-tableReqCh:
			switch req.Cmd {
			case AddBet:
				fmt.Printf("bet received for %v\n", req.BetSlip.betNums)
				participants = append(participants, req.BetSlip)
			case NotifyParticipants:

				for _, participant := range participants {
					_, winCheck := participant.betNums[winningNumber]

					if winCheck {
						participant.Win = true
						participant.userCh <- participant
					} else {
						participant.userCh <- participant
					}
				}
				go clearParticipants()

				ticker.Reset(1 * time.Minute)
			case ClearParticipants:
				participants = make([]Betslip, 0)
			}
		}
	}
}

// Generate a random winning number.
func spinTheWheel() {

	randomSeed := rand.New(seed)
	winningNumber = randomSeed.Intn(37)
	fmt.Printf("winning number is %d\n", winningNumber)
	TableRequestChannel <- TableRequest{
		Cmd: NotifyParticipants,
	}
	return
}

// Clear the list of participants ready for the next game.
func clearParticipants() {
	TableRequestChannel <- TableRequest{
		Cmd: ClearParticipants,
	}
}
