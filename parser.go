package pgn

import (
	"fmt"
)

type Parser struct {
	l *Lexer

	errors []string

	currToken Token
	peekToken Token
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

  p.nextToken()
  p.nextToken()

	return p
}

func (p *Parser) ParsePGN() {
  return 
}

func (p *Parser) Errors() []string {
    return p.errors
}

func (p *Parser) nextToken() {
  p.currToken = p.peekToken
  p.peekToken = p.l.NextToken()
}

func (p *Parser) currTokenIs(t TokenType) bool {
    return p.currToken.Type == t
}

func (p *Parser) peekTokenIs(t TokenType) bool {
    return p.peekToken.Type == t
}

func (p *Parser) peekError(t TokenType) {
    msg := fmt.Sprintf("expected next token to be %s, got %s instead\n", t, p.peekToken.Type)
    p.errors = append(p.errors, msg)
}
