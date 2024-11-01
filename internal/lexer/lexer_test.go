package lexer

import (
	"testing"

	"github.com/Shobhit-Nagpal/pgn/internal/token"
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
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACKET, "["},
		{token.SYMBOL, "Event"},
		{token.STRING, "F/S Return Match"},
		{token.RBRACKET, "]"},
		{token.LBRACKET, "["},
		{token.SYMBOL, "Site"},
		{token.STRING, "Belgrade, Serbia JUG"},
		{token.RBRACKET, "]"},
		{token.LBRACKET, "["},
		{token.SYMBOL, "Date"},
		{token.STRING, "1992.11.04"},
		{token.RBRACKET, "]"},
		{token.LBRACKET, "["},
		{token.SYMBOL, "Round"},
		{token.STRING, "29"},
		{token.RBRACKET, "]"},
		{token.LBRACKET, "["},
		{token.SYMBOL, "White"},
		{token.STRING, "Fischer, Robert J."},
		{token.RBRACKET, "]"},
		{token.LBRACKET, "["},
		{token.SYMBOL, "Black"},
		{token.STRING, "Spassky, Boris V."},
		{token.RBRACKET, "]"},
		{token.LBRACKET, "["},
		{token.SYMBOL, "Result"},
		{token.STRING, "1/2-1/2"},
		{token.RBRACKET, "]"},
		{token.INTEGER, "1"},
		{token.PERIOD, "."},
		{token.SYMBOL, "e4"},
		{token.SYMBOL, "e5"},
		{token.INTEGER, "2"},
		{token.PERIOD, "."},
		{token.SYMBOL, "Nf3"},
		{token.SYMBOL, "Nc6"},
		{token.INTEGER, "3"},
		{token.PERIOD, "."},
		{token.SYMBOL, "Bb5"},
		{token.SYMBOL, "a6"},
		{token.INTEGER, "4"},
		{token.PERIOD, "."},
		{token.SYMBOL, "Ba4"},
		{token.SYMBOL, "Nf6"},
		{token.INTEGER, "5"},
		{token.PERIOD, "."},
		{token.SYMBOL, "O-O"},
		{token.SYMBOL, "Be7"},
		{token.INTEGER, "6"},
		{token.PERIOD, "."},
		{token.SYMBOL, "Re1"},
		{token.SYMBOL, "b5"},
		{token.INTEGER, "7"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Bb3"},
		{token.SYMBOL, "d6"},
		{token.INTEGER, "8"},
		{token.PERIOD, "."},
    {token.SYMBOL, "c3"},
		{token.SYMBOL, "O-O"},
		{token.INTEGER, "9"},
		{token.PERIOD, "."},
    {token.SYMBOL, "h3"},
		{token.SYMBOL, "Nb8"},
		{token.INTEGER, "10"},
		{token.PERIOD, "."},
    {token.SYMBOL, "d4"},
		{token.SYMBOL, "Nbd7"},
		{token.INTEGER, "11"},
		{token.PERIOD, "."},
    {token.SYMBOL, "c4"},
		{token.SYMBOL, "c6"},
		{token.INTEGER, "12"},
		{token.PERIOD, "."},
    {token.SYMBOL, "cxb5"},
		{token.SYMBOL, "axb5"},
		{token.INTEGER, "13"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Nc3"},
		{token.SYMBOL, "Bb7"},
		{token.INTEGER, "14"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Bg5"},
		{token.SYMBOL, "b4"},
		{token.INTEGER, "15"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Nb1"},
		{token.SYMBOL, "h6"},
		{token.INTEGER, "16"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Bh4"},
		{token.SYMBOL, "c5"},
		{token.INTEGER, "17"},
		{token.PERIOD, "."},
    {token.SYMBOL, "dxe5"},
		{token.SYMBOL, "Nxe4"},
		{token.INTEGER, "18"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Bxe7"},
		{token.SYMBOL, "Qxe7"},
		{token.INTEGER, "19"},
		{token.PERIOD, "."},
    {token.SYMBOL, "exd6"},
		{token.SYMBOL, "Qf6"},
		{token.INTEGER, "20"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Nbd2"},
		{token.SYMBOL, "Nxd6"},
		{token.INTEGER, "21"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Nc4"},
		{token.SYMBOL, "Nxc4"},
		{token.INTEGER, "22"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Bxc4"},
		{token.SYMBOL, "Nb6"},
		{token.INTEGER, "23"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Ne5"},
		{token.SYMBOL, "Rae8"},
		{token.INTEGER, "24"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Bxf7+"},
		{token.SYMBOL, "Rxf7"},
		{token.INTEGER, "25"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Nxf7"},
		{token.SYMBOL, "Rxe1+"},
		{token.INTEGER, "26"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Qxe1"},
		{token.SYMBOL, "Kxf7"},
		{token.INTEGER, "27"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Qe3"},
		{token.SYMBOL, "Qg5"},
		{token.INTEGER, "28"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Qxg5"},
		{token.SYMBOL, "hxg5"},
		{token.INTEGER, "29"},
		{token.PERIOD, "."},
    {token.SYMBOL, "b3"},
		{token.SYMBOL, "Ke6"},
		{token.INTEGER, "30"},
		{token.PERIOD, "."},
    {token.SYMBOL, "a3"},
		{token.SYMBOL, "Kd6"},
		{token.INTEGER, "31"},
		{token.PERIOD, "."},
    {token.SYMBOL, "axb4"},
		{token.SYMBOL, "cxb4"},
		{token.INTEGER, "32"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Ra5"},
		{token.SYMBOL, "Nd5"},
		{token.INTEGER, "33"},
		{token.PERIOD, "."},
    {token.SYMBOL, "f3"},
		{token.SYMBOL, "Bc8"},
		{token.INTEGER, "34"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Kf2"},
		{token.SYMBOL, "Bf5"},
		{token.INTEGER, "35"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Ra7"},
		{token.SYMBOL, "g6"},
		{token.INTEGER, "36"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Ra6+"},
		{token.SYMBOL, "Kc5"},
		{token.INTEGER, "37"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Ke1"},
		{token.SYMBOL, "Nf4"},
		{token.INTEGER, "38"},
		{token.PERIOD, "."},
    {token.SYMBOL, "g3"},
		{token.SYMBOL, "Nxh3"},
		{token.INTEGER, "39"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Kd2"},
		{token.SYMBOL, "Kb5"},
		{token.INTEGER, "40"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Rd6"},
		{token.SYMBOL, "Kc5"},
		{token.INTEGER, "41"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Ra6"},
		{token.SYMBOL, "Nf2"},
		{token.INTEGER, "42"},
		{token.PERIOD, "."},
    {token.SYMBOL, "g4"},
		{token.SYMBOL, "Bd3"},
		{token.INTEGER, "43"},
		{token.PERIOD, "."},
    {token.SYMBOL, "Re6"},
		{token.SYMBOL, "1-0"},
		{token.ASTERIX, "*"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

    t.Logf("%s", tok.Literal)

		if tok.Type != tt.expectedType {
			t.Fatalf("tests [%d] -- tokentype wrong. expected=%q, got=%q\n", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests [%d] -- literal wrong. expected=%q, got=%q\n", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
