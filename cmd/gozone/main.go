package main

import (
	"bufio"
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/isudin/gozone/cmd/gozone/commands"
	"github.com/isudin/gozone/internal/domain"
	"github.com/isudin/gozone/internal/repository/sqlc"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	queries := initDatabase()
	// TODO: Move CreateFaction to init_command.go
	queries.CreateFaction(context.Background(), "Test")
	commands.Commands["help"].Exec(nil)
}

func initDatabase() *sqlc.Queries {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("could not connect to database\n", err)
	}
	return sqlc.New(db)
}

func getPlayerData() {
	reader := bufio.NewReader(os.Stdin)
	println("Welcome to the Zone, stalker. What's your name?")
	name, _ := reader.ReadString('\n')
	player := domain.Player{Name: name}
	println("Good hunting, " + player.Name + ".")
	println("What faction do you wish to join?")
}
