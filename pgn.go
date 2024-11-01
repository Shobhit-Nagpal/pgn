package pgn

import (
	"errors"
	"fmt"
)

type Game struct {
	tags   map[string]string
	moves  map[int]*Move
	result string
}

func (g *Game) GetTag(name string) string {
	return g.tags[name]
}

func (g *Game) SetTag(tag, value string) error {
	if _, exists := g.tags[tag]; exists {
		return errors.New(fmt.Sprintf("%s tag already exists in the tag pair section", tag))
	}

	g.tags[tag] = value

	return nil
}

func (g *Game) TagPairs() map[string]string {
	return g.tags
}

func (g *Game) Event() string {
	return g.tags["Event"]
}

func (g *Game) Site() string {
	return g.tags["Site"]
}

func (g *Game) Round() string {
	return g.tags["Round"]
}

func (g *Game) Date() string {
	return g.tags["Date"]
}

func (g *Game) White() string {
	return g.tags["White"]
}

func (g *Game) Black() string {
	return g.tags["Black"]
}

func (g *Game) Result() string {
	return g.result
}

func (g *Game) GetMove(number int) *Move {
	return g.moves[number]
}

func (g *Game) SetMove(number int, move *Move) {
	g.moves[number] = move
}

func (g *Game) Moves() map[int]*Move {
	return g.moves
}
