package main

import (
	"fmt"
	"net/http"

	"github.com/Harshmist/roulette-api/api"
	"github.com/Harshmist/roulette-api/table"
)

func main() {

	fmt.Print(`
░█▀▄░█▀█░█░█░█░░░█▀▀░▀█▀░▀█▀░█▀▀░░░░░█▀█░█▀█░▀█▀
░█▀▄░█░█░█░█░█░░░█▀▀░░█░░░█░░█▀▀░▄▄▄░█▀█░█▀▀░░█░
░▀░▀░▀▀▀░▀▀▀░▀▀▀░▀▀▀░░▀░░░▀░░▀▀▀░░░░░▀░▀░▀░░░▀▀▀
`)
	go table.TableStart()

	http.HandleFunc("/", api.RequestRouter)
	http.ListenAndServe(":8080", nil)
}
