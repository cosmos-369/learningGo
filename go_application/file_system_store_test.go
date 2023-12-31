package poker

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from reader", func(t *testing.T) {
		database, cleanDataBase := CreateTempFile(t, `[
			{"Name":"Cleo", "Wins":10},
			{"Name":"Chris", "Wins":33}]`)
		defer cleanDataBase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoErr(t, err)

		got := store.GetLeague()

		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}
		AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDataBase := CreateTempFile(t, `[
			{"Name":"Cleo", "Wins":10},
			{"Name":"Chris", "Wins":33}]`)
		defer cleanDataBase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoErr(t, err)

		got := store.GetPlayerScore("Cleo")

		AssertPlayerScore(t, got, 10)
	})

	t.Run("store win of an existing player", func(t *testing.T) {
		database, cleanDataBase := CreateTempFile(t, `[
			{"Name":"Cleo", "Wins":10},
			{"Name":"Chris", "Wins":33}]`)
		defer cleanDataBase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoErr(t, err)

		store.RecordWin("Cleo")
		store.RecordWin("Cleo")

		got := store.GetPlayerScore("Cleo")
		AssertPlayerScore(t, got, 12)
	})

	t.Run("stores wins of a new player", func(t *testing.T) {
		database, cleanDataBase := CreateTempFile(t, `[
			{"Name":"Cleo", "Wins":10},
			{"Name":"Chris", "Wins":33}]`)
		defer cleanDataBase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoErr(t, err)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		AssertPlayerScore(t, got, want)
	})

	t.Run("runs with an empty file", func(t *testing.T) {
		database, cleanDataBase := CreateTempFile(t, ``)
		defer cleanDataBase()

		_, err := NewFileSystemPlayerStore(database)

		AssertNoErr(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoErr(t, err)

		got := store.GetLeague()

		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		AssertLeague(t, got, want)

		got = store.GetLeague()
		AssertLeague(t, got, want)
	})
}

func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tempfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create a temp file, %v", err)
	}

	tempfile.Write([]byte(initialData))

	removeFile := func() {
		tempfile.Close()
		os.Remove(tempfile.Name())
	}

	return tempfile, removeFile
}
