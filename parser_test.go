package pgn

import "testing"

func TestTagPairs(t *testing.T) {
	tests := []struct {
		input         string
		expectedName  string
		expectedValue string
	}{
		{`[Event "F/S Return Match"]`, "Event", "F/S Return Match"},
		{`[Site "Belgrade, Serbia JUG"]`, "Site", "Belgrade, Serbia JUG"},
		{`[Date "1992.11.04"]`, "Date", "1992.11.04"},
		{`[Round "29"]`, "Round", "29"},
		{`[White "Fischer, Robert J."]`, "White", "Fischer, Robert J."},
		{`[Black "Spassky, Boris V."]`, "Black", "Spassky, Boris V."},
		{`[Result "1/2-1/2"]`, "Result", "1/2-1/2"},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		p := NewParser(l)
		game, _ := p.ParsePGN()
		checkParserErrors(t, p)

		if value, exists := game.TagPairs()[tt.expectedName]; !exists {
			t.Errorf("tag %q not found in game tags", tt.expectedName)
		} else if value != tt.expectedValue {
			t.Errorf("tag %q has wrong value. got=%q, want=%q",
				tt.expectedName, value, tt.expectedValue)
		}

	}
}

func TestGetTag(t *testing.T) {
	tests := []struct {
		input         string
		expectedName  string
		expectedValue string
	}{
		{`[Event "F/S Return Match"]`, "Event", "F/S Return Match"},
		{`[Site "Belgrade, Serbia JUG"]`, "Site", "Belgrade, Serbia JUG"},
		{`[Date "1992.11.04"]`, "Date", "1992.11.04"},
		{`[Round "29"]`, "Round", "29"},
		{`[White "Fischer, Robert J."]`, "White", "Fischer, Robert J."},
		{`[Black "Spassky, Boris V."]`, "Black", "Spassky, Boris V."},
		{`[Result "1/2-1/2"]`, "Result", "1/2-1/2"},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		p := NewParser(l)
		game, _ := p.ParsePGN()
		checkParserErrors(t, p)

		if got := game.GetTag(tt.expectedName); got != tt.expectedValue {
			t.Errorf("GetTag(%q) wrong value. got=%q, want=%q",
				tt.expectedName, got, tt.expectedValue)
		}
	}
}

func TestMoves(t *testing.T) {
	tests := []struct {
		input              string
		expectedMoveNumber int
		expectedMoveWhite  string
		expectedMoveBlack  string
	}{
		{"1. e4 $69 e5", 1, "e4", "e5"},
		{"2. Nf3 Nc6", 2, "Nf3", "Nc6"},
		{"3. Bb5 a6", 3, "Bb5", "a6"},
		{"4. Ba4 Nf6", 4, "Ba4", "Nf6"},
		{"12. cxb5 axb5", 12, "cxb5", "axb5"},
		{"24. Bxf7+ Rxf7", 24, "Bxf7+", "Rxf7"},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		p := NewParser(l)
		game, _ := p.ParsePGN()
		checkParserErrors(t, p)

		if got := game.GetMove(tt.expectedMoveNumber).White(); got != tt.expectedMoveWhite {
			t.Errorf("GetMove(%q) wrong value. got=%q, want=%q",
				tt.expectedMoveNumber, got, tt.expectedMoveWhite)
		}

	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestMovesWithNoPeriods(t *testing.T) {
	tests := []struct {
		input              string
		expectedMoveNumber int
		expectedMoveWhite  string
		expectedMoveBlack  string
	}{
		{"1 e4 e5", 1, "e4", "e5"},
		{"2 Nf3 Nc6", 2, "Nf3", "Nc6"},
		{"3 Bb5 a6", 3, "Bb5", "a6"},
		{"4 Ba4 Nf6", 4, "Ba4", "Nf6"},
		{"12 cxb5 axb5", 12, "cxb5", "axb5"},
		{"24 Bxf7+ Rxf7", 24, "Bxf7+", "Rxf7"},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		p := NewParser(l)
		game, _ := p.ParsePGN()
		checkParserErrors(t, p)

		if got := game.GetMove(tt.expectedMoveNumber).White(); got != tt.expectedMoveWhite {
			t.Errorf("GetMove(%q) wrong value. got=%q, want=%q",
				tt.expectedMoveNumber, got, tt.expectedMoveWhite)
		}

	}
}

func TestMovesWithThreePeriods(t *testing.T) {
	tests := []struct {
		input              string
		expectedMoveNumber int
		expectedMoveWhite  string
		expectedMoveBlack  string
	}{
		{"1... e4 e5", 1, "e4", "e5"},
		{"2... Nf3 Nc6", 2, "Nf3", "Nc6"},
		{"3... Bb5 a6", 3, "Bb5", "a6"},
		{"4... Ba4 Nf6", 4, "Ba4", "Nf6"},
		{"12... cxb5 axb5", 12, "cxb5", "axb5"},
		{"24... Bxf7+ Rxf7", 24, "Bxf7+", "Rxf7"},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		p := NewParser(l)
		game, _ := p.ParsePGN()
		checkParserErrors(t, p)

		if got := game.GetMove(tt.expectedMoveNumber).White(); got != tt.expectedMoveWhite {
			t.Errorf("GetMove(%q) wrong value. got=%q, want=%q",
				tt.expectedMoveNumber, got, tt.expectedMoveWhite)
		}

	}
}

func TestCompletePGN(t *testing.T) {
	input := `
  [Event "Live Chess"]
  [Site "Chess.com"]
  [Date "2024.11.03"]
  [Round "?"]
  [White "shbhtngpl"]
  [Black "Michal_Chmara_2002"]
  [Result "1-0"]
  [TimeControl "180+2"]
  [WhiteElo "1616"]
  [BlackElo "1674"]
  [Termination "shbhtngpl won by resignation"]
  [Link "https://www.chess.com/game/live/124369775413"]

  1. d4 Nf6 2. c4 g6 3. Nf3 Bg7 4. g3 O-O 5. Bg2 d6 6. O-O c5 7. e3 cxd4 8. exd4
  Bg4 9. Nbd2 Qc7 10. h3 Bd7 11. b3 Nc6 12. Bb2 Rac8 13. Re1 Rfe8 14. d5 Nb4 15.
  Ne4 Nxe4 16. Bxg7 Nxg3 17. Bc3 Qb6 18. Bxb4 Qxb4 19. fxg3 Qc5+ 20. Qd4 Qc7 21.
  Qh4 e6 22. Ng5 h5 23. Qf4 Rf8 24. Qf6 exd5 25. Bxd5 Bc6 26. Rac1 Bxd5 27. cxd5
  Qb6+ 28. Kh2 Rxc1 29. Rxc1 Qe3 30. Rf1 Qe2+ 31. Rf2 Qa6 32. Ne6 1-0
  `

	expectedTags := map[string]string{
		"Event":       "Live Chess",
		"Site":        "Chess.com",
		"Date":        "2024.11.03",
		"Round":       "?",
		"White":       "shbhtngpl",
		"Black":       "Michal_Chmara_2002",
		"Result":      "1-0",
		"TimeControl": "180+2",
		"WhiteElo":    "1616",
		"BlackElo":    "1674",
		"Termination": "shbhtngpl won by resignation",
		"Link":        "https://www.chess.com/game/live/124369775413",
	}

	expectedMoves := map[int]*Move{
		1: &Move{
			MoveNumber: 1,
			MoveWhite:  "d4",
			MoveBlack:  "Nf6",
		},
		2: &Move{
			MoveNumber: 2,
			MoveWhite:  "c4",
			MoveBlack:  "g6",
		},
		3: &Move{
			MoveNumber: 3,
			MoveWhite:  "Nf3",
			MoveBlack:  "Bg7",
		},
		4: &Move{
			MoveNumber: 4,
			MoveWhite:  "g3",
			MoveBlack:  "O-O",
		},
		5: &Move{
			MoveNumber: 5,
			MoveWhite:  "Bg2",
			MoveBlack:  "d6",
		},
		6: &Move{
			MoveNumber: 6,
			MoveWhite:  "O-O",
			MoveBlack:  "c5",
		},
		7: &Move{
			MoveNumber: 7,
			MoveWhite:  "e3",
			MoveBlack:  "cxd4",
		},
		8: &Move{
			MoveNumber: 8,
			MoveWhite:  "exd4",
			MoveBlack:  "Bg4",
		},
		9: &Move{
			MoveNumber: 9,
			MoveWhite:  "Nbd2",
			MoveBlack:  "Qc7",
		},
		10: &Move{
			MoveNumber: 10,
			MoveWhite:  "h3",
			MoveBlack:  "Bd7",
		},
		11: &Move{
			MoveNumber: 11,
			MoveWhite:  "b3",
			MoveBlack:  "Nc6",
		},
		12: &Move{
			MoveNumber: 12,
			MoveWhite:  "Bb2",
			MoveBlack:  "Rac8",
		},
		13: &Move{
			MoveNumber: 13,
			MoveWhite:  "Re1",
			MoveBlack:  "Rfe8",
		},
		14: &Move{
			MoveNumber: 14,
			MoveWhite:  "d5",
			MoveBlack:  "Nb4",
		},
		15: &Move{
			MoveNumber: 15,
			MoveWhite:  "Ne4",
			MoveBlack:  "Nxe4",
		},
		16: &Move{
			MoveNumber: 16,
			MoveWhite:  "Bxg7",
			MoveBlack:  "Nxg3",
		},
		17: &Move{
			MoveNumber: 17,
			MoveWhite:  "Bc3",
			MoveBlack:  "Qb6",
		},
		18: &Move{
			MoveNumber: 18,
			MoveWhite:  "Bxb4",
			MoveBlack:  "Qxb4",
		},
		19: &Move{
			MoveNumber: 19,
			MoveWhite:  "fxg3",
			MoveBlack:  "Qc5+",
		},
		20: &Move{
			MoveNumber: 20,
			MoveWhite:  "Qd4",
			MoveBlack:  "Qc7",
		},
		21: &Move{
			MoveNumber: 21,
			MoveWhite:  "Qh4",
			MoveBlack:  "e6",
		},
		22: &Move{
			MoveNumber: 22,
			MoveWhite:  "Ng5",
			MoveBlack:  "h5",
		},
		23: &Move{
			MoveNumber: 23,
			MoveWhite:  "Qf4",
			MoveBlack:  "Rf8",
		},
		24: &Move{
			MoveNumber: 24,
			MoveWhite:  "Qf6",
			MoveBlack:  "exd5",
		},
		25: &Move{
			MoveNumber: 25,
			MoveWhite:  "Bxd5",
			MoveBlack:  "Bc6",
		},
		26: &Move{
			MoveNumber: 26,
			MoveWhite:  "Rac1",
			MoveBlack:  "Bxd5",
		},
		27: &Move{
			MoveNumber: 27,
			MoveWhite:  "cxd5",
			MoveBlack:  "Qb6+",
		},
		28: &Move{
			MoveNumber: 28,
			MoveWhite:  "Kh2",
			MoveBlack:  "Rxc1",
		},
		29: &Move{
			MoveNumber: 29,
			MoveWhite:  "Rxc1",
			MoveBlack:  "Qe3",
		},
		30: &Move{
			MoveNumber: 30,
			MoveWhite:  "Rf1",
			MoveBlack:  "Qe2+",
		},
		31: &Move{
			MoveNumber: 31,
			MoveWhite:  "Rf2",
			MoveBlack:  "Qa6",
		},
		32: &Move{
			MoveNumber: 32,
			MoveWhite:  "Ne6",
			MoveBlack:  "",
		},
	}

	l := NewLexer(input)
	p := NewParser(l)
	game, _ := p.ParsePGN()
	checkParserErrors(t, p)

	for expectedKey, expectedValue := range expectedTags {
		if value, exists := game.TagPairs()[expectedKey]; !exists {
			t.Errorf("tag %q not found in game tags", expectedKey)
		} else if value != expectedValue {
			t.Errorf("tag %q has wrong value. got=%q, want=%q",
				expectedKey, value, expectedValue)
		}
	}

	for expectedMoveNumber, expectedValue := range expectedMoves {
		if got := game.GetMove(expectedMoveNumber).White(); got != expectedValue.White() {
			t.Errorf("GetMove(%q) wrong value. got=%q, want=%q",
				expectedMoveNumber, got, expectedValue.White())
		}

		if got := game.GetMove(expectedMoveNumber).Black(); got != expectedValue.Black() {
			t.Errorf("GetMove(%q) wrong value. got=%q, want=%q",
				expectedMoveNumber, got, expectedValue.Black())
		}

	}

}
