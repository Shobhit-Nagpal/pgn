package pgn

type Game struct {
	tags   map[string]string
	moves  []Move
	result string
}

func (g *Game) GetTag(name string) string {
	return g.tags[name]
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
	if number > len(g.moves) {
		return nil
	}

	return &g.moves[number-1]
}
