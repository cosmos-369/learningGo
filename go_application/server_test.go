package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerScores struct {
	score    map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerScores) GetPlayerScore(name string) int {
	score := s.score[name]
	return score
}

func (s *StubPlayerScores) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerScores) GetLeague() League {
	return s.league
}
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

		assertStatus(t, response.Code, http.StatusOK)
		assertResponceBody(t, response.Body.String(), "20")
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponceBody(t, response.Body.String(), "10")
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

		assertStatus(t, responce.Code, http.StatusOK)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner, got %q, want %q", store.winCalls[0], player)
		}
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

		assertStatus(t, responce.Code, http.StatusOK)
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

		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, response, jsonContentType)
	})
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	league, _ = NewLeague(body)
	return
}

func assertLeague(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func getNewLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponceBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get the correct status, got %d , want %d", got, want)
	}
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != jsonContentType {
		t.Errorf("responce did not have content type application/json, got %v", response.Result().Header)
	}
}
