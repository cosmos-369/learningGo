package poker_test

import (
	"encoding/json"
	"fmt"
	poker "go_application"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

const ten = time.Millisecond * 10

func TestGETPlayers(t *testing.T) {
	store := poker.StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		[]string{},
		nil,
	}
	server := mustMakePlayerServer(t, &store, &GameSpy{})
	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusOK)
		poker.AssertResponceBody(t, response.Body.String(), "20")
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusOK)
		poker.AssertResponceBody(t, response.Body.String(), "10")
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

	store := poker.StubPlayerStore{
		map[string]int{},
		[]string{},
		nil,
	}

	server := mustMakePlayerServer(t, &store, &GameSpy{})
	t.Run("it records win when POST", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		responce := httptest.NewRecorder()

		server.ServeHTTP(responce, request)

		poker.AssertStatus(t, responce, http.StatusOK)
		poker.AssertPlayerWin(t, &store, "Pepper")
	})
}

func TestLeague(t *testing.T) {
	store := poker.StubPlayerStore{}
	server := mustMakePlayerServer(t, &store, &GameSpy{})

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		responce := httptest.NewRecorder()

		server.ServeHTTP(responce, request)

		var got []poker.Player

		err := json.NewDecoder(responce.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse responce from server %q into slice of Player, %v", responce.Body, err)
		}

		poker.AssertStatus(t, responce, http.StatusOK)
	})

	t.Run("it returns league table as JSON", func(t *testing.T) {
		wantedLeague := []poker.Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := poker.StubPlayerStore{nil, nil, wantedLeague}
		server := mustMakePlayerServer(t, &store, &GameSpy{})

		request := getNewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		poker.AssertStatus(t, response, http.StatusOK)
		poker.AssertLeague(t, got, wantedLeague)
		poker.AssertContentType(t, response, poker.JsonContentType)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		store := &poker.StubPlayerStore{}
		server := mustMakePlayerServer(t, store, &GameSpy{})

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusOK)
	})

	t.Run("start the game with 3 player and declrae Ruth as winner", func(t *testing.T) {

		wantedBlindAlert := "Blind is 100"
		winner := "Ruth"

		store := &poker.StubPlayerStore{}
		game := &GameSpy{BlindAlert: []byte(wantedBlindAlert)}
		server := httptest.NewServer(mustMakePlayerServer(t, store, game))
		ws := mustDailWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")

		defer server.Close()
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		assertGameStartedWith(t, game, 3)
		assertGameFinishCalledWith(t, game, winner)

		within(t, ten, func() { assertWebSocketGotMsg(t, ws, wantedBlindAlert) })
	})
}

func assertWebSocketGotMsg(t testing.TB, ws *websocket.Conn, want string) {
	_, msg, _ := ws.ReadMessage()

	if string(msg) != want {
		t.Errorf("got blind alert %q, want %q", string(msg), want)
	}
}
func within(t testing.TB, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{}, 1)
	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case <-time.After(d):
		t.Error("time out")
	case <-done:
	}
}
func getLeagueFromResponse(t testing.TB, body io.Reader) (league []poker.Player) {
	t.Helper()
	league, _ = poker.NewLeague(body)
	return
}

func getNewLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func newGameRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/game", nil)
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

func mustMakePlayerServer(t testing.TB, store poker.PlayerStore, game *GameSpy) *poker.PlayerServer {
	t.Helper()
	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		t.Fatalf("problem creataing player server, %v", err)
	}
	return server
}

func writeWSMessage(t testing.TB, ws *websocket.Conn, message string) {
	t.Helper()
	if err := ws.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not sent winner over ws connection, %v", err)
	}
}

func mustDailWS(t testing.TB, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("could not open ws connection on %s %v", url, err)
	}
	return ws
}
