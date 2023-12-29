package poker_test

import (
	"fmt"
	poker "go_application"
	"testing"
	"time"
)

func TestGame(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)
		game.Start(5)

		cases := []ScheduledAlert{
			{0 * time.Minute, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)
		game.Start(7)

		cases := []ScheduledAlert{
			{0 * time.Minute, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})
}

func checkSchedulingCases(t *testing.T, cases []ScheduledAlert, blindAlerter *SpyBlindAlerter) {
	t.Helper()
	for i, c := range cases {
		t.Run(fmt.Sprintf("%d scheduled for %v", c.amount, c.at), func(t *testing.T) {

			if len(blindAlerter.alerts) < i {
				t.Fatalf("alert %d was not scheduled, %v", i, blindAlerter.alerts)
			}

			alert := blindAlerter.alerts[i]
			assertScheduledAlert(t, alert, c)
		})
	}
}

func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewTexasHoldem(dummyBlindAlerter, store)
	winner := "Ruth"

	game.Finish(winner)
	poker.AssertPlayerWin(t, store, winner)
}
