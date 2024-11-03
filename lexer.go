package pgn

type Lexer struct {
	input        string
	position     int  // Current position
	readPosition int  // Position to read (after current position)
	ch           byte // Current character under examination
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '.':
		tok = newToken(PERIOD, l.ch)
	case '*':
		tok = newToken(ASTERIX, l.ch)
	case '[':
		tok = newToken(LBRACKET, l.ch)
	case ']':
		tok = newToken(RBRACKET, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case '<':
		tok = newToken(LANGLE, l.ch)
	case '%':
		tok = newToken(PERCENTAGE, l.ch)
	case '>':
		tok = newToken(RANGLE, l.ch)
	case '"':
		tok.Type = STRING
		tok.Literal = l.readString()
	case '$':
		tok.Type = NAG
		tok.Literal = l.readNAG()
	case 0:
		tok.Type = EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) || isDigit(l.ch) {
			tok.Literal, tok.Type = l.readSymbolOrInteger()
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
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

func (l *Lexer) readSymbolOrInteger() (string, TokenType) {
	flag := false
	position := l.position

	for isDigit(l.ch) || isLetter(l.ch) || isSpecialChar(l.ch) {
		if l.peekChar() == '.' || l.peekChar() == '*' {
			flag = true
			break
		}

		l.readChar()
	}

	var tokenLiteral string

	if flag {
		tokenLiteral = l.input[position:l.readPosition]
	} else {
		tokenLiteral = l.input[position:l.position]
	}

	isInteger := isDigitsOnly(tokenLiteral)

	if isInteger {
		return tokenLiteral, INTEGER
	}

	return tokenLiteral, SYMBOL
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
