package main

import (
	"fmt"
	poker "go_application"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.NewFileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("lets player poker")
	fmt.Println("Type {Name} wins to record a win")
	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.StdOutAlerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
