package main

import (
	"bufio"
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
	queries := initDBConnection()
	cmds := commands.InitCommands(queries)

	// First argument is program name
	if len(os.Args) < 2 {
		cmds["help"].Exec(nil)
		os.Exit(0)
	}

	// Second argument should be command name
	cmd := cmds[os.Args[1]]
	if cmd.Name == "" {
		log.Fatal("command not found")
	}

	cmd.Exec(nil)
}

func initDBConnection() *sqlc.Queries {
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
