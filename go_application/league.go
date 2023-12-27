package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}

	return nil
}

func NewLeague(r io.Reader) ([]Player, error) {
	var players []Player
	err := json.NewDecoder(r).Decode(&players)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return players, err
}
