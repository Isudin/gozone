package main

import (
	"bufio"
	"os"

	"github.com/isudin/gozone/cmd/gozone/commands"
	"github.com/isudin/gozone/internal/domain"
)

func main() {
	commands.Commands["help"].Exec(nil)
}

func getPlayerData() {
	reader := bufio.NewReader(os.Stdin)
	println("Welcome to the Zone, stalker. What's your name?")
	name, _ := reader.ReadString('\n')
	player := domain.Player{Name: name}
	println("Good hunting, " + player.Name + ".")
	println("What faction do you wish to join?")
}
