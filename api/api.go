package api

import (
	"fmt"
	"github.com/Harshmist/roulette-api/table"
	"net/http"
	"strings"
)

// RequestRouter will allow paths to be wildcards meaning new functionality will not require coding a new handler each time.
func RequestRouter(w http.ResponseWriter, r *http.Request) {
	route := parseRoute(r)
	userResponse := make(chan table.Betslip)

	switch route {
	case "bet":
		betSlip, err := table.BuildBetSlip(userResponse, r)

		if err != nil {
			// simulate logging
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)

			return
		}
		table.TableRequestChannel <- table.TableRequest{
			Cmd:     table.AddBet,
			BetSlip: betSlip,
		}

		// Wait for win/loss response from the table.
		response := checkBetslip(userResponse)

		fmt.Fprintln(w, response)
	}
}

func parseRoute(r *http.Request) string {
	parts := strings.Split(r.URL.Path, "/")

	return parts[len(parts)-1]
}

func checkBetslip(userResponse chan table.Betslip) string {
	var response string

	for slip := range userResponse {
		if slip.Win == true {
			winnings := slip.Stake * slip.WinMultiplier
			close(userResponse)
			response = fmt.Sprintf("Congratulations! You have won £%d from your £%d stake.\n", winnings, slip.Stake)
		} else {
			close(userResponse)
			response = fmt.Sprintf("Bad luck! You didn't win this time")
		}

	}
	return response
}
