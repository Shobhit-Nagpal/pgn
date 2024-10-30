package lexer

import "github.com/Shobhit-Nagpal/pgn/internal/token"

type Lexer struct {
	input        string
	position     int  // Current position
	readPosition int  // Position to read (after current position)
	ch           byte // Current character under examination
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition > len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '.':
		tok = newToken(token.PERIOD, l.ch)
	case '*':
		tok = newToken(token.ASTERIX, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '<':
		tok = newToken(token.LANGLE, l.ch)
	case '%':
		tok = newToken(token.PERCENTAGE, l.ch)
	case '>':
		tok = newToken(token.RANGLE, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '$':
		tok.Type = token.NAG
		tok.Literal = l.readNAG()
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) || isDigit(l.ch) {
			tok.Type = token.SYMBOL
			tok.Literal - l.readSymbol()
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNAG() string {
	position := l.position + 1

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readSymbol() string {
	position := l.position

	for isDigit(l.ch) || isLetter(l.ch) || isSpecialChar(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}
