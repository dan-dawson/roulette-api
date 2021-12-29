package main

import (
	"fmt"
	"github.com/Harshmist/roulette-api/api"
	"github.com/Harshmist/roulette-api/table"
	"net/http"
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
