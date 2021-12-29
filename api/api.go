package api

import (
	"fmt"
	"net/http"
	"strings"
	"table"
)

func RequestRouter(w http.ResponseWriter, r *http.Request) {
	route := parseRoute(r)

	switch route {
	case "bet":
		betSlip, err := table.BuildBetSlip(r)

		if err != nil {
			// simulate logging
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}

	}
}

func parseRoute(r *http.Request) string {
	parts := strings.Split(r.URL.Path, "/")

	return parts[len(parts)-1]
}
