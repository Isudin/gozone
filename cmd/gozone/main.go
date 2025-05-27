package main

import "github.com/isudin/gozone/internal/domain"

func main() {
	println("Welcome to the Zone, stalker.")
	player := domain.Player{Name: "Test"}
	println("Good hunting, " + player.Name)
}
