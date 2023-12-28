package poker_test

import (
	poker "go_application"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {

	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins")
		playerStore := &poker.StubPlayerScores{}
		cli := poker.NewCLI(playerStore, in)

		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins")
		playerStore := &poker.StubPlayerScores{}
		cli := poker.NewCLI(playerStore, in)

		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})
}
