package table

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Betslip struct {
	userCh        chan Betslip
	betNums       map[int]struct{}
	betType       string
	WinMultiplier int
	Stake         int
	Win           bool
}

// BuildBetSlip will be used to generate the betslip used by the participants slice.
func BuildBetSlip(userChan chan Betslip, r *http.Request) (Betslip, error) {
	var err error
	query := r.URL.Query()

	numbers, err := parseBetNumbers(query)
	if err != nil {
		return Betslip{}, err
	}

	betType, err := parseBetType(query)
	if err != nil {
		return Betslip{}, err
	}
	multiplier := setBetMultiplier(betType)
	if multiplier < 1 {
		return Betslip{}, fmt.Errorf("multiplier set to 0")
	}
	stake, err := parseStake(query)
	if err != nil {
		return Betslip{}, err
	}

	newBetSlip := Betslip{
		userCh:        userChan,
		betNums:       numbers,
		betType:       betType,
		WinMultiplier: multiplier,
		Stake:         stake,
		Win:           false,
	}
	return newBetSlip, nil
}

func parseBetNumbers(values url.Values) (map[int]struct{}, error) {
	numMap := make(map[int]struct{}, 0)

	numbers, present := values["n"]

	if !present {
		return numMap, fmt.Errorf("bet number selection not present")
	}

	for _, v := range numbers {
		numInt, err := strconv.Atoi(v)
		if err != nil || numInt > 36 || numInt < 0 {
			return numMap, fmt.Errorf("error parsing bet selection numbers")
		}
		numMap[numInt] = struct{}{}
	}
	return numMap, nil
}

func parseBetType(values url.Values) (string, error) {
	betType, present := values["type"]

	if !present || len(betType) > 1 {
		return "", fmt.Errorf("no bet type found")
	}

	return betType[0], nil
}

func setBetMultiplier(betType string) int {

	switch betType {
	case "0":
		return 35
	case "straight":
		return 35
	case "row":
		return 17
	case "split":
		return 17
	case "street":
		return 11
	case "corner":
		return 8
	case "basket":
		return 8
	case "doublestreet":
		return 5
	case "column":
		return 2
	case "dozen":
		return 2
	case "oddeven":
		return 1
	case "colour":
		return 1
	case "highlow":
		return 1
	default:
		return 0
	}
}

func parseStake(values url.Values) (int, error) {
	stake, present := values["stake"]

	if !present || len(stake) > 1 {
		return 0, fmt.Errorf("incorrect stake format")
	}

	return strconv.Atoi(stake[0])
}
