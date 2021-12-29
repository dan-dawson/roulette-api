package table

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	participants        []Betslip
	TableRequestChannel chan TableRequest
	seed                rand.Source
	winningNumber       int
)

const (
	AddBet = iota
	ClearTable
	NotifyParticipants
	SpinTheWheel
)

type TableRequest struct {
	cmd     int
	betSlip Betslip
}

func TableStart() {
	participants = make([]Betslip, 0)
	TableRequestChannel = make(chan TableRequest)
	seed = rand.NewSource(time.Now().UnixNano())

	tableStateWorker(TableRequestChannel)
}

func tableStateWorker(tableReqCh chan TableRequest) {

	// This ticker will be used to give time for users to place bets. In the real world, the duration would be set via a config file or flag.
	ticker := time.NewTimer(1 * time.Minute)

	for {
		select {
		case <-ticker.C:
			tableReqCh <- TableRequest{
				cmd: SpinTheWheel,
			}
		case req := <-tableReqCh:
			switch req.cmd {
			case AddBet:
				participants = append(participants, req.betSlip)
			case SpinTheWheel:
				randomSeed := rand.New(seed)
				winningNumber = randomSeed.Intn(37)
				fmt.Printf("winning number is %d\n", winningNumber)

				TableRequestChannel <- TableRequest{
					cmd: NotifyParticipants,
				}
			case NotifyParticipants:
				for _, participant := range participants {
					_, winCheck := participant.betNums[winningNumber]

					if winCheck {
						participant.win = true
						participant.userCh <- participant
					} else {
						participant.userCh <- participant
					}
				}
				TableRequestChannel <- TableRequest{
					cmd: ClearTable,
				}
				ticker.Reset(1 * time.Minute)
			case ClearTable:
				participants = make([]Betslip, 0)
			}
		}
	}
}
