package poker

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	score    map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.score[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func AssertPlayerScore(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("invalid player score, got %d, want %d", got, want)
	}
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner, got %q, want %q", store.winCalls[0], winner)
	}
}

func AssertLeague(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertResponceBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get the correct status, got %d , want %d", got, want)
	}
}

func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != jsonContentType {
		t.Errorf("responce did not have content type application/json, got %v", response.Result().Header)
	}
}

func AssertNoErr(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
