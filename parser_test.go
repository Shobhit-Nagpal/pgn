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
		l := newLexer(tt.input)
		p := newParser(l)
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
		l := newLexer(tt.input)
		p := newParser(l)
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
		l := newLexer(tt.input)
		p := newParser(l)
		game, _ := p.ParsePGN()
		checkParserErrors(t, p)

		if got := game.GetMove(tt.expectedMoveNumber).White(); got != tt.expectedMoveWhite {
			t.Errorf("GetMove(%q) wrong value. got=%q, want=%q",
				tt.expectedMoveNumber, got, tt.expectedMoveWhite)
		}

	}
}

func checkParserErrors(t *testing.T, p *parser) {
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
		l := newLexer(tt.input)
		p := newParser(l)
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
		l := newLexer(tt.input)
		p := newParser(l)
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

	l := newLexer(input)
	p := newParser(l)
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

func TestCompletePGNWithComments(t *testing.T) {
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

  1. { {First move} } d4 Nf6 2. c4 g6 3. Nf3 Bg7 4. g3 O-O 5. Bg2 d6 6. O-O c5 7. e3 cxd4 8. exd4
  Bg4 9. Nbd2 Qc7 10. h3 Bd7 11. b3 Nc6 12. Bb2 Rac8 13. Re1 Rfe8 14. d5 Nb4 15.
  Ne4 Nxe4 16. Bxg7 Nxg3 17. Bc3 Qb6 18. Bxb4 Qxb4 19. fxg3 Qc5+ 20. Qd4 Qc7 21.
  Qh4 e6 22. Ng5 h5 23. Qf4 Rf8 24. Qf6 { BRILLIANT MOVE } exd5 25. Bxd5 Bc6 26. Rac1 Bxd5 27. cxd5
  Qb6+ 28. Kh2 Rxc1 29. Rxc1 Qe3 30. Rf1 Qe2+ 31. Rf2 Qa6 32. Ne6 1-0 { WHITE WINS }
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

	l := newLexer(input)
	p := newParser(l)
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

func TestCompletePGNWithRestOfLineComments(t *testing.T) {
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

  1. d4 ; First move
  Nf6 2. c4 g6 3. Nf3 Bg7 4. g3 O-O 5. Bg2 d6 6. O-O c5 7. e3 cxd4 8. exd4
  Bg4 9. Nbd2 Qc7 10. h3 Bd7 11. b3 Nc6 12. Bb2 Rac8 13. Re1 Rfe8 14. d5 Nb4 15.
  Ne4 Nxe4 16. Bxg7 Nxg3 17. Bc3 Qb6 18. Bxb4 Qxb4 19. fxg3 Qc5+ 20. Qd4 Qc7 21.
  Qh4 e6 22. Ng5 h5 23. Qf4 Rf8 24. Qf6 exd5 ; BRILLIANT MOVE
  25. Bxd5 Bc6 26. Rac1 Bxd5 27. cxd5
  Qb6+ 28. Kh2 Rxc1 29. Rxc1 Qe3 30. Rf1 Qe2+ 31. Rf2 Qa6 32. Ne6 1-0 ; WHITE WINS
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

	l := newLexer(input)
	p := newParser(l)
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

func TestLichessGameWithRestOfLineComments(t *testing.T) {
	input := `[Event "casual blitz game"]
[Site "https://lichess.org/FPn6uGag"]
[Date "2025.09.21"]
[White "Mihir504"]
[Black "Anonymous"]
[Result "1-0"]
[GameId "FPn6uGag"]
[UTCDate "2025.09.21"]
[UTCTime "15:46:43"]
[WhiteElo "1770"]
[BlackElo "?"]
[Variant "Standard"]
[TimeControl "180+0"]
[ECO "D94"]
[Opening "Grünfeld Defense: Three Knights Variation, Burille Variation"]
[Termination "Time forfeit"]
[Annotator "lichess.org"]

1. d4 Nf6 2. c4 d5 3. Nc3 g6 4. Nf3 Bg7 5. e3 O-O ; D94 Grünfeld Defense: Three Knights Variation, Burille Variation
6. Be2 Nc6 7. O-O Bf5 8. h3 a6 9. a3 dxc4 10. Bxc4 b5 11. Bd3 Bxd3 12. Qxd3 Na5 13. b4 Nb3 14. Rb1 Nxc1 15. Rbxc1 Rc8 16. e4 e6 17. e5 Nd5 18. Nxd5 Qxd5 19. Rc5 Qd7 20. Rfc1 Rfe8 21. Qc3 Re7 22. Rc6 Bh6 23. Rc2 Kg7 24. Rxa6 f5 25. exf6+ Kxf6 26. d5+ Kf7 27. dxe6+ Qxe6 28. Rxe6 Rxe6 29. Qb3 Ke7 30. Nd4 Re1+ 31. Kh2 Bf4+ 32. g3 Bd6 33. Nc6+ Kf8 34. Qf3+ Kg8 35. Qf8+ 1-0 ; White wins on time.
`

	expectedTags := map[string]string{
		"Event":       "casual blitz game",
		"Site":        "https://lichess.org/FPn6uGag",
		"Date":        "2025.09.21",
		"White":       "Mihir504",
		"Black":       "Anonymous",
		"Result":      "1-0",
		"GameId":      "FPn6uGag",
		"UTCDate":     "2025.09.21",
		"UTCTime":     "15:46:43",
		"WhiteElo":    "1770",
		"BlackElo":    "?",
		"Variant":     "Standard",
		"TimeControl": "180+0",
		"ECO":         "D94",
		"Opening":     "Grünfeld Defense: Three Knights Variation, Burille Variation",
		"Termination": "Time forfeit",
		"Annotator":   "lichess.org",
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
			MoveBlack:  "d5",
		},
		3: &Move{
			MoveNumber: 3,
			MoveWhite:  "Nc3",
			MoveBlack:  "g6",
		},
		4: &Move{
			MoveNumber: 4,
			MoveWhite:  "Nf3",
			MoveBlack:  "Bg7",
		},
		5: &Move{
			MoveNumber: 5,
			MoveWhite:  "e3",
			MoveBlack:  "O-O",
		},
		6: &Move{
			MoveNumber: 6,
			MoveWhite:  "Be2",
			MoveBlack:  "Nc6",
		},
		7: &Move{
			MoveNumber: 7,
			MoveWhite:  "O-O",
			MoveBlack:  "Bf5",
		},
		8: &Move{
			MoveNumber: 8,
			MoveWhite:  "h3",
			MoveBlack:  "a6",
		},
		9: &Move{
			MoveNumber: 9,
			MoveWhite:  "a3",
			MoveBlack:  "dxc4",
		},
		10: &Move{
			MoveNumber: 10,
			MoveWhite:  "Bxc4",
			MoveBlack:  "b5",
		},
		11: &Move{
			MoveNumber: 11,
			MoveWhite:  "Bd3",
			MoveBlack:  "Bxd3",
		},
		12: &Move{
			MoveNumber: 12,
			MoveWhite:  "Qxd3",
			MoveBlack:  "Na5",
		},
		13: &Move{
			MoveNumber: 13,
			MoveWhite:  "b4",
			MoveBlack:  "Nb3",
		},
		14: &Move{
			MoveNumber: 14,
			MoveWhite:  "Rb1",
			MoveBlack:  "Nxc1",
		},
		15: &Move{
			MoveNumber: 15,
			MoveWhite:  "Rbxc1",
			MoveBlack:  "Rc8",
		},
		16: &Move{
			MoveNumber: 16,
			MoveWhite:  "e4",
			MoveBlack:  "e6",
		},
		17: &Move{
			MoveNumber: 17,
			MoveWhite:  "e5",
			MoveBlack:  "Nd5",
		},
		18: &Move{
			MoveNumber: 18,
			MoveWhite:  "Nxd5",
			MoveBlack:  "Qxd5",
		},
		19: &Move{
			MoveNumber: 19,
			MoveWhite:  "Rc5",
			MoveBlack:  "Qd7",
		},
		20: &Move{
			MoveNumber: 20,
			MoveWhite:  "Rfc1",
			MoveBlack:  "Rfe8",
		},
		21: &Move{
			MoveNumber: 21,
			MoveWhite:  "Qc3",
			MoveBlack:  "Re7",
		},
		22: &Move{
			MoveNumber: 22,
			MoveWhite:  "Rc6",
			MoveBlack:  "Bh6",
		},
		23: &Move{
			MoveNumber: 23,
			MoveWhite:  "Rc2",
			MoveBlack:  "Kg7",
		},
		24: &Move{
			MoveNumber: 24,
			MoveWhite:  "Rxa6",
			MoveBlack:  "f5",
		},
		25: &Move{
			MoveNumber: 25,
			MoveWhite:  "exf6+",
			MoveBlack:  "Kxf6",
		},
		26: &Move{
			MoveNumber: 26,
			MoveWhite:  "d5+",
			MoveBlack:  "Kf7",
		},
		27: &Move{
			MoveNumber: 27,
			MoveWhite:  "dxe6+",
			MoveBlack:  "Qxe6",
		},
		28: &Move{
			MoveNumber: 28,
			MoveWhite:  "Rxe6",
			MoveBlack:  "Rxe6",
		},
		29: &Move{
			MoveNumber: 29,
			MoveWhite:  "Qb3",
			MoveBlack:  "Ke7",
		},
		30: &Move{
			MoveNumber: 30,
			MoveWhite:  "Nd4",
			MoveBlack:  "Re1+",
		},
		31: &Move{
			MoveNumber: 31,
			MoveWhite:  "Kh2",
			MoveBlack:  "Bf4+",
		},
		32: &Move{
			MoveNumber: 32,
			MoveWhite:  "g3",
			MoveBlack:  "Bd6",
		},
		33: &Move{
			MoveNumber: 33,
			MoveWhite:  "Nc6+",
			MoveBlack:  "Kf8",
		},
		34: &Move{
			MoveNumber: 34,
			MoveWhite:  "Qf3+",
			MoveBlack:  "Kg8",
		},
		35: &Move{
			MoveNumber: 35,
			MoveWhite:  "Qf8+",
			MoveBlack:  "",
		},
	}

	l := newLexer(input)
	p := newParser(l)
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
			t.Errorf("GetMove(%d) white move wrong value. got=%q, want=%q",
				expectedMoveNumber, got, expectedValue.White())
		}

		if got := game.GetMove(expectedMoveNumber).Black(); got != expectedValue.Black() {
			t.Errorf("GetMove(%d) black move wrong value. got=%q, want=%q",
				expectedMoveNumber, got, expectedValue.Black())
		}
	}
}
