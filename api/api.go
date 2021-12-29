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

	switch route {
	case "bet":
		betSlip, err := table.BuildBetSlip(r)

		if err != nil {
			// simulate logging
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)

			return
		}

	}
}

func parseRoute(r *http.Request) string {
	parts := strings.Split(r.URL.Path, "/")

	return parts[len(parts)-1]
}
