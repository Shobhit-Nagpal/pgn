package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	PERIOD  = "."
	ASTERIX = "*"
	STRING  = "STRING"
	INTEGER = "INTEGER"

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
