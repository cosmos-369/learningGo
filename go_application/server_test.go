package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerScores{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		[]string{},
		nil,
	}
	server := NewPlayerServer(&store)
	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponceBody(t, response.Body.String(), "20")
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponceBody(t, response.Body.String(), "10")
	})

	t.Run("return 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got status %d, want status %d", got, want)
		}
	})
}

func TestScoreWins(t *testing.T) {

	store := StubPlayerScores{
		map[string]int{},
		[]string{},
		nil,
	}

	server := NewPlayerServer(&store)
	t.Run("it records win when POST", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		responce := httptest.NewRecorder()

		server.ServeHTTP(responce, request)

		AssertStatus(t, responce.Code, http.StatusOK)
		AssertPlayerWin(t, &store, "Pepper")
	})
}

func TestLeague(t *testing.T) {
	store := StubPlayerScores{}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		responce := httptest.NewRecorder()

		server.ServeHTTP(responce, request)

		var got []Player

		err := json.NewDecoder(responce.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse responce from server %q into slice of Player, %v", responce.Body, err)
		}

		AssertStatus(t, responce.Code, http.StatusOK)
	})

	t.Run("it returns league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerScores{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := getNewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertLeague(t, got, wantedLeague)
		AssertContentType(t, response, jsonContentType)
	})
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	league, _ = NewLeague(body)
	return
}

func getNewLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}
