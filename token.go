package pgn

type tokenType string

type token struct {
	Type    tokenType
	Literal string
}

func (t token) TokenLiteral() string {
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

	LBRACE = "{"
	RBRACE = "}"

	SEMICOLON = ";"

	LANGLE = "<"
	RANGLE = ">"

	NAG    = "NAG"
	SYMBOL = "SYMBOL"

	MISC = "MISC"

	EOF = "EOF"
)
