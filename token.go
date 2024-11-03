package pgn

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func (t Token) TokenLiteral() string {
	return t.Literal
}

const (
	PERIOD     = "."
	ASTERIX    = "*"
	PERCENTAGE = "%"
	STRING     = "STRING"
	INTEGER    = "INTEGER"
	ILLEGAL    = "ILLEGAL"

	LBRACKET = "["
	RBRACKET = "]"

	LPAREN = "("
	RPAREN = ")"

	LANGLE = "<"
	RANGLE = ">"

	NAG    = "NAG"
	SYMBOL = "SYMBOL"

	EOF = "EOF"
)
