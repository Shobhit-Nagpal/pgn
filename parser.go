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

func (p *Parser) ParsePGN() *Game {
	game := &Game{
		tags:  map[string]string{},
		moves: []*Move{},
	}

	for p.currToken.Type != EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			switch v := stmt.(type) {
			case *TagPair:
				game.SetTag(v.Name(), v.Value())
			case *Move:
        game.moves = append(game.moves, v)
			}
		}
	}

	return game
}

func (p *Parser) parseStatement() Stmt {
	switch p.currToken.Type {
	case LBRACKET:
		return p.parseTagPair()
	case INTEGER:
		return p.parseMove()
	default:
		return nil
	}
}

func (p *Parser) parseTagPair() *TagPair {
	tp := &TagPair{
		LBracket: p.currToken,
	}

	if !p.expectPeek(SYMBOL) {
		return nil
	}

	tp.TagName = p.currToken.TokenLiteral()

	if !p.expectPeek(STRING) {
		return nil
	}

	tp.TagValue = p.currToken.TokenLiteral()

	if p.peekTokenIs(RBRACKET) {
		p.nextToken()
		tp.RBracket = p.currToken
	}

	return tp
}

func (p *Parser) parseMove() *Move {
	move := &Move{
		MoveNumber: p.currToken.TokenLiteral(),
	}

	//Zero or more periods
	for p.peekTokenIs(PERIOD) {
		p.nextToken()
	}

	p.nextToken()

	if !p.expectPeek(SYMBOL) || !p.expectPeek(ASTERIX) {
		return nil
	}

	move.MoveWhite = p.currToken.TokenLiteral()
	move.MoveBlack = p.currToken.TokenLiteral()

	if !p.expectPeek(INTEGER) {
		return nil
	}

	return move
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

func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
