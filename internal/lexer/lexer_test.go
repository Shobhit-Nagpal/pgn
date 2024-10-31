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

  1. e4 e5 2. Nf3 Nc6
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
