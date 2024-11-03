package pgn

import (
	"testing"
)

func TestGame_GetTag(t *testing.T) {
	game := &Game{
		tags: map[string]string{
			"CustomTag": "CustomValue",
		},
	}

	if got := game.GetTag("CustomTag"); got != "CustomValue" {
		t.Errorf("Game.GetTag() = %v, want %v", got, "CustomValue")
	}

	// Test non-existent tag
	if got := game.GetTag("NonExistent"); got != "" {
		t.Errorf("Game.GetTag() for non-existent tag = %v, want empty string", got)
	}
}

func TestGame_TagPairs(t *testing.T) {
	tags := map[string]string{
		"Tag1": "Value1",
		"Tag2": "Value2",
	}
	game := &Game{
		tags: tags,
	}

	got := game.TagPairs()
	if len(got) != len(tags) {
		t.Errorf("Game.TagPairs() length = %v, want %v", len(got), len(tags))
	}
	for k, v := range tags {
		if gotValue := got[k]; gotValue != v {
			t.Errorf("Game.TagPairs()[%v] = %v, want %v", k, gotValue, v)
		}
	}
}

func TestGame_StandardTags(t *testing.T) {
	game := &Game{
		tags: map[string]string{
			"Event": "Test Event",
			"Site":  "Test Site",
			"Round": "1",
			"Date":  "2024.01.01",
			"White": "Player 1",
			"Black": "Player 2",
		},
	}

	tests := []struct {
		name     string
		got      string
		expected string
		fn       func() string
	}{
		{"Event", game.Event(), "Test Event", game.Event},
		{"Site", game.Site(), "Test Site", game.Site},
		{"Round", game.Round(), "1", game.Round},
		{"Date", game.Date(), "2024.01.01", game.Date},
		{"White", game.White(), "Player 1", game.White},
		{"Black", game.Black(), "Player 2", game.Black},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fn(); got != tt.expected {
				t.Errorf("Game.%s() = %v, want %v", tt.name, got, tt.expected)
			}
		})
	}

	// Test missing tags return empty string
	emptyGame := &Game{
		tags: map[string]string{},
	}
	if got := emptyGame.Event(); got != "" {
		t.Errorf("Game.Event() with no tags = %v, want empty string", got)
	}
}

func TestGame_Result(t *testing.T) {
	game := &Game{
		result: "1-0",
	}

	if got := game.Result(); got != "1-0" {
		t.Errorf("Game.Result() = %v, want %v", got, "1-0")
	}

	// Test empty result
	emptyGame := &Game{}
	if got := emptyGame.Result(); got != "" {
		t.Errorf("Game.Result() empty game = %v, want empty string", got)
	}
}

func TestGame_Moves(t *testing.T) {
	moves := map[int]*Move{
		1: {MoveNumber: 1, MoveWhite: "e4", MoveBlack: "e5"},
		2: {MoveNumber: 2, MoveWhite: "Nf3", MoveBlack: "Nc6"},
	}
	game := &Game{
		moves: moves,
	}

	got := game.Moves()
	if len(got) != len(moves) {
		t.Errorf("Game.Moves() length = %v, want %v", len(got), len(moves))
	}
	for k, v := range moves {
		if gotMove := got[k]; gotMove != v {
			t.Errorf("Game.Moves()[%v] = %v, want %v", k, gotMove, v)
		}
	}
}

func TestGame_GetMove(t *testing.T) {
	move := &Move{MoveNumber: 1, MoveWhite: "e4", MoveBlack: "e5"}
	game := &Game{
		moves: map[int]*Move{
			1: move,
		},
	}

	if got := game.GetMove(1); got != move {
		t.Errorf("Game.GetMove() = %v, want %v", got, move)
	}

	// Test non-existent move
	if got := game.GetMove(999); got != nil {
		t.Errorf("Game.GetMove() for non-existent move = %v, want nil", got)
	}
}
