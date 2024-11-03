package pgn

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
  [Event "F/S Return Match"]
  [Site "Belgrade, Serbia JUG"]
  [Date "1992.11.04"]
  [Round "29"]
  [White "Fischer, Robert J."]
  [Black "Spassky, Boris V."]
  [Result "1/2-1/2"]

  1. e4 e5 2. Nf3 Nc6 3. Bb5 a6 4. Ba4 Nf6 5. O-O Be7 6. Re1 b5 7. Bb3 d6 8. c3
  O-O 9. h3 Nb8 10. d4 Nbd7 11. c4 c6 12. cxb5 axb5 13. Nc3 Bb7 14. Bg5 b4 15.
  Nb1 h6 16. Bh4 c5 17. dxe5 Nxe4 18. Bxe7 Qxe7 19. exd6 Qf6 20. Nbd2 Nxd6 21.
  Nc4 Nxc4 22. Bxc4 Nb6 23. Ne5 Rae8 24. Bxf7+ Rxf7 25. Nxf7 Rxe1+ 26. Qxe1 Kxf7
  27. Qe3 Qg5 28. Qxg5 hxg5 29. b3 Ke6 30. a3 Kd6 31. axb4 cxb4 32. Ra5 Nd5 33.
  f3 Bc8 34. Kf2 Bf5 35. Ra7 g6 36. Ra6+ Kc5 37. Ke1 Nf4 38. g3 Nxh3 39. Kd2 Kb5
  40. Rd6 Kc5 41. Ra6 Nf2 42. g4 Bd3 43. Re6 1/2-1/2  
  `

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LBRACKET, "["},
		{SYMBOL, "Event"},
		{STRING, "F/S Return Match"},
		{RBRACKET, "]"},
		{LBRACKET, "["},
		{SYMBOL, "Site"},
		{STRING, "Belgrade, Serbia JUG"},
		{RBRACKET, "]"},
		{LBRACKET, "["},
		{SYMBOL, "Date"},
		{STRING, "1992.11.04"},
		{RBRACKET, "]"},
		{LBRACKET, "["},
		{SYMBOL, "Round"},
		{STRING, "29"},
		{RBRACKET, "]"},
		{LBRACKET, "["},
		{SYMBOL, "White"},
		{STRING, "Fischer, Robert J."},
		{RBRACKET, "]"},
		{LBRACKET, "["},
		{SYMBOL, "Black"},
		{STRING, "Spassky, Boris V."},
		{RBRACKET, "]"},
		{LBRACKET, "["},
		{SYMBOL, "Result"},
		{STRING, "1/2-1/2"},
		{RBRACKET, "]"},
		{INTEGER, "1"},
		{PERIOD, "."},
		{SYMBOL, "e4"},
		{SYMBOL, "e5"},
		{INTEGER, "2"},
		{PERIOD, "."},
		{SYMBOL, "Nf3"},
		{SYMBOL, "Nc6"},
		{INTEGER, "3"},
		{PERIOD, "."},
		{SYMBOL, "Bb5"},
		{SYMBOL, "a6"},
		{INTEGER, "4"},
		{PERIOD, "."},
		{SYMBOL, "Ba4"},
		{SYMBOL, "Nf6"},
		{INTEGER, "5"},
		{PERIOD, "."},
		{SYMBOL, "O-O"},
		{SYMBOL, "Be7"},
		{INTEGER, "6"},
		{PERIOD, "."},
		{SYMBOL, "Re1"},
		{SYMBOL, "b5"},
		{INTEGER, "7"},
		{PERIOD, "."},
		{SYMBOL, "Bb3"},
		{SYMBOL, "d6"},
		{INTEGER, "8"},
		{PERIOD, "."},
		{SYMBOL, "c3"},
		{SYMBOL, "O-O"},
		{INTEGER, "9"},
		{PERIOD, "."},
		{SYMBOL, "h3"},
		{SYMBOL, "Nb8"},
		{INTEGER, "10"},
		{PERIOD, "."},
		{SYMBOL, "d4"},
		{SYMBOL, "Nbd7"},
		{INTEGER, "11"},
		{PERIOD, "."},
		{SYMBOL, "c4"},
		{SYMBOL, "c6"},
		{INTEGER, "12"},
		{PERIOD, "."},
		{SYMBOL, "cxb5"},
		{SYMBOL, "axb5"},
		{INTEGER, "13"},
		{PERIOD, "."},
		{SYMBOL, "Nc3"},
		{SYMBOL, "Bb7"},
		{INTEGER, "14"},
		{PERIOD, "."},
		{SYMBOL, "Bg5"},
		{SYMBOL, "b4"},
		{INTEGER, "15"},
		{PERIOD, "."},
		{SYMBOL, "Nb1"},
		{SYMBOL, "h6"},
		{INTEGER, "16"},
		{PERIOD, "."},
		{SYMBOL, "Bh4"},
		{SYMBOL, "c5"},
		{INTEGER, "17"},
		{PERIOD, "."},
		{SYMBOL, "dxe5"},
		{SYMBOL, "Nxe4"},
		{INTEGER, "18"},
		{PERIOD, "."},
		{SYMBOL, "Bxe7"},
		{SYMBOL, "Qxe7"},
		{INTEGER, "19"},
		{PERIOD, "."},
		{SYMBOL, "exd6"},
		{SYMBOL, "Qf6"},
		{INTEGER, "20"},
		{PERIOD, "."},
		{SYMBOL, "Nbd2"},
		{SYMBOL, "Nxd6"},
		{INTEGER, "21"},
		{PERIOD, "."},
		{SYMBOL, "Nc4"},
		{SYMBOL, "Nxc4"},
		{INTEGER, "22"},
		{PERIOD, "."},
		{SYMBOL, "Bxc4"},
		{SYMBOL, "Nb6"},
		{INTEGER, "23"},
		{PERIOD, "."},
		{SYMBOL, "Ne5"},
		{SYMBOL, "Rae8"},
		{INTEGER, "24"},
		{PERIOD, "."},
		{SYMBOL, "Bxf7+"},
		{SYMBOL, "Rxf7"},
		{INTEGER, "25"},
		{PERIOD, "."},
		{SYMBOL, "Nxf7"},
		{SYMBOL, "Rxe1+"},
		{INTEGER, "26"},
		{PERIOD, "."},
		{SYMBOL, "Qxe1"},
		{SYMBOL, "Kxf7"},
		{INTEGER, "27"},
		{PERIOD, "."},
		{SYMBOL, "Qe3"},
		{SYMBOL, "Qg5"},
		{INTEGER, "28"},
		{PERIOD, "."},
		{SYMBOL, "Qxg5"},
		{SYMBOL, "hxg5"},
		{INTEGER, "29"},
		{PERIOD, "."},
		{SYMBOL, "b3"},
		{SYMBOL, "Ke6"},
		{INTEGER, "30"},
		{PERIOD, "."},
		{SYMBOL, "a3"},
		{SYMBOL, "Kd6"},
		{INTEGER, "31"},
		{PERIOD, "."},
		{SYMBOL, "axb4"},
		{SYMBOL, "cxb4"},
		{INTEGER, "32"},
		{PERIOD, "."},
		{SYMBOL, "Ra5"},
		{SYMBOL, "Nd5"},
		{INTEGER, "33"},
		{PERIOD, "."},
		{SYMBOL, "f3"},
		{SYMBOL, "Bc8"},
		{INTEGER, "34"},
		{PERIOD, "."},
		{SYMBOL, "Kf2"},
		{SYMBOL, "Bf5"},
		{INTEGER, "35"},
		{PERIOD, "."},
		{SYMBOL, "Ra7"},
		{SYMBOL, "g6"},
		{INTEGER, "36"},
		{PERIOD, "."},
		{SYMBOL, "Ra6+"},
		{SYMBOL, "Kc5"},
		{INTEGER, "37"},
		{PERIOD, "."},
		{SYMBOL, "Ke1"},
		{SYMBOL, "Nf4"},
		{INTEGER, "38"},
		{PERIOD, "."},
		{SYMBOL, "g3"},
		{SYMBOL, "Nxh3"},
		{INTEGER, "39"},
		{PERIOD, "."},
		{SYMBOL, "Kd2"},
		{SYMBOL, "Kb5"},
		{INTEGER, "40"},
		{PERIOD, "."},
		{SYMBOL, "Rd6"},
		{SYMBOL, "Kc5"},
		{INTEGER, "41"},
		{PERIOD, "."},
		{SYMBOL, "Ra6"},
		{SYMBOL, "Nf2"},
		{INTEGER, "42"},
		{PERIOD, "."},
		{SYMBOL, "g4"},
		{SYMBOL, "Bd3"},
		{INTEGER, "43"},
		{PERIOD, "."},
		{SYMBOL, "Re6"},
		{SYMBOL, "1/2-1/2"},
		{EOF, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests [%d] -- tokentype wrong. expected=%q, got=%q\n", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests [%d] -- literal wrong. expected=%q, got=%q\n", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
