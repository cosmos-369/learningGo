package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordWinsAndRetrievingThem(t *testing.T) {
	database, cleanDataBase := createTempFile(t, `[]`)
	defer cleanDataBase()

	store, err := NewFileSystemPlayerStore(database)
	AssertNoErr(t, err)

	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponceBody(t, response.Body.String(), "3")
	})

	t.Run("get League", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, getNewLeagueRequest())

		AssertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}

		AssertLeague(t, got, want)
	})

}
