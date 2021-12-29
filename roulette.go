package main

import (
	"github.com/Harshmist/roulette-api/api"
	"github.com/Harshmist/roulette-api/table"
	"net/http"
)

func main() {
	go table.TableStart()

	http.HandleFunc("/", api.RequestRouter)
	http.ListenAndServe(":8080", nil)
}
