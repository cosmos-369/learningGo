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
	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
