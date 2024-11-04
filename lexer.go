package pgn

type lexer struct {
	input        string
	position     int  // Current position
	readPosition int  // Position to read (after current position)
	ch           byte // Current character under examination
}

func newLexer(input string) *lexer {
	l := &lexer{
		input: input,
	}

	l.readChar()
	return l
}

func (l *lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *lexer) NextToken() token {
	var tok token

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

func newToken(tokenType tokenType, ch byte) token {
	return token{Type: tokenType, Literal: string(ch)}
}

func (l *lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

func (l *lexer) readNAG() string {
  position := l.position + 1

  l.readChar()

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *lexer) readSymbolOrInteger() (string, tokenType) {
	flag := false
	position := l.position

	for isDigit(l.ch) || isLetter(l.ch) || isSpecialChar(l.ch) {
		if l.peekChar() == '.' || l.peekChar() == '*' || l.peekChar() == '$' {
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

func (l *lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
