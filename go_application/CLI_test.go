package poker_test

import (
	"bytes"
	"fmt"
	poker "go_application"
	"strings"
	"testing"
	"time"
)

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

type ScheduledAlert struct {
	at     time.Duration
	amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, ScheduledAlert{
		duration,
		amount,
	})
}

var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyStdOut = &bytes.Buffer{}
var dummyPlayerStore = &poker.StubPlayerStore{}

func TestCLI(t *testing.T) {

	t.Run("start the game with 7 players and record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("7\nChris wins\n")
		game := &GameSpy{}
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()
		got := game.FinishedWith
		if game.FinishedWith != "Chris" {
			t.Errorf("wrong winner recorded, got %q, want %q", got, "Chris")
		}
	})

	t.Run("start the game with 5 players and record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")
		game := &GameSpy{}
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()
		got := game.FinishedWith
		if game.FinishedWith != "Cleo" {
			t.Errorf("wrong winner recorded, got %q, want %q", got, "Cleo")
		}
	})

	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")

		game := &GameSpy{}
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)

		if game.StartedWith != 7 {
			t.Errorf("wanted to start with 7 players, but got %d", game.StartedWith)
		}
	})

	t.Run("it prints an error when non numberic value is entered and does not start the game", func(t *testing.T) {
		in := strings.NewReader("Pies\n")
		stdout := &bytes.Buffer{}
		game := &GameSpy{}

		c := poker.NewCLI(in, stdout, game)
		c.PlayPoker()

		if game.StartCalled {
			t.Error("game should not have started")
		}

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}

func assertScheduledAlert(t testing.TB, got, want ScheduledAlert) {
	t.Helper()
	if got.amount != want.amount {
		t.Errorf("amount got %d, want %d", got.amount, want.amount)
	}

	if got.at != want.at {
		t.Errorf("got scheduled time %v, want %v", got.at, want.at)
	}
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}
