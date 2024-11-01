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
		l := New(tt.input)
		p := NewParser(l)
		game := p.ParsePGN()
		checkParserErrors(t, p)

		if value, exists := game.TagPairs()[tt.expectedName]; !exists {
			t.Errorf("tag %q not found in game tags", tt.expectedName)
		} else if value != tt.expectedValue {
			t.Errorf("tag %q has wrong value. got=%q, want=%q",
				tt.expectedName, value, tt.expectedValue)
		}

    t.Logf("%s, %s", tt.expectedName, game.TagPairs()[tt.expectedName])

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
		l := New(tt.input)
		p := NewParser(l)
		game := p.ParsePGN()
		checkParserErrors(t, p)

		if got := game.GetTag(tt.expectedName); got != tt.expectedValue {
			t.Errorf("GetTag(%q) wrong value. got=%q, want=%q",
				tt.expectedName, got, tt.expectedValue)
		}
	}
}

func TestMoves(t *testing.T) {
	tests := []struct {
		input         string
		expectedMoveNumber  int
		expectedMoveWhite string
		expectedMoveBlack string
	}{
		{"1. e4 e5", 1, "e4", "e5"},
		{"2. Nf3 Nc6", 2, "Nf3", "Nc6"},
		{"3. Bb5 a6", 3, "Bb5", "a6"},
		{"4. Ba4 Nf6", 4, "Ba4", "Nf6"},
		{"12. cxb5 axb5", 12, "cxb5", "axb5"},
		{"24. Bxf7+ Rxf7", 24, "Bxf7+", "Rxf7"},
	}

	for _, tt := range tests {
		l := New(tt.input)
		p := NewParser(l)
		game := p.ParsePGN()
		checkParserErrors(t, p)

		if got := game.GetMove(tt.expectedMoveNumber).White(); got != tt.expectedMoveWhite {
			t.Errorf("GetMove(%q) wrong value. got=%q, want=%q",
				tt.expectedMoveNumber, got, tt.expectedMoveWhite)
		}

    moves := game.Moves()

    t.Logf("%d, %s, %s", tt.expectedMoveNumber, moves[tt.expectedMoveNumber].White(), moves[tt.expectedMoveNumber].Black())
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
