package poker_test

import (
	poker "go_application"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordWinsAndRetrievingThem(t *testing.T) {
	database, cleanDataBase := poker.CreateTempFile(t, `[]`)
	defer cleanDataBase()

	store, err := poker.NewFileSystemPlayerStore(database)
	poker.AssertNoErr(t, err)

	server := mustMakePlayerServer(t, store, &GameSpy{})
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		poker.AssertStatus(t, response, http.StatusOK)
		poker.AssertResponceBody(t, response.Body.String(), "3")
	})

	t.Run("get League", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, getNewLeagueRequest())

		poker.AssertStatus(t, response, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []poker.Player{
			{"Pepper", 3},
		}

		poker.AssertLeague(t, got, want)
	})

}
