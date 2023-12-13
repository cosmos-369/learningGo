package main

import (
	"testing"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		defination := "this is just a test"
		err := dictionary.Add(word, defination)

		assertError(t, err, nil)
		assertDefination(t, dictionary, word, defination)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dictionary := Dictionary{word: defination}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefination(t, dictionary, word, defination)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		newDefination := "updated test"

		dictionary := Dictionary{word: defination}
		err := dictionary.Update(word, newDefination)

		assertError(t, err, nil)
		assertDefination(t, dictionary, word, newDefination)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, defination)
		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	word := "test"
	defination := "this is just a test"

	dictionary := Dictionary{word: defination}
	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertDefination(t testing.TB, dictionary Dictionary, word string, defination string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, defination)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
