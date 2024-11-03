package pgn

import (
	"fmt"
	"log"
	"strconv"
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
		moves: map[int]*Move{},
	}

	for p.currToken.Type != EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			switch v := stmt.(type) {
			case *TagPair:
				game.SetTag(v.Name(), v.Value())
			case *Move:
				game.SetMove(v.Number(), v)
			case *GameTermination:
				game.SetResult(v.Value())
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
	case SYMBOL:
		if isGameResult(p.currToken.TokenLiteral()) {
      gt := &GameTermination{TerminationValue: p.currToken.TokenLiteral()}
      p.nextToken()
      return gt
		}
		return nil
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

	if p.expectPeek(RBRACKET) {
		tp.RBracket = p.currToken
	}

	p.nextToken()

	return tp
}

func (p *Parser) parseMove() *Move {

	moveNumInt, err := strconv.Atoi(p.currToken.TokenLiteral())
	if err != nil {
		log.Fatalf("Couldn't convert string to integer for moves: %s", p.currToken.TokenLiteral())
	}

	move := &Move{
		MoveNumber: moveNumInt,
	}

	//Zero or more periods
	for p.peekTokenIs(PERIOD) {
		p.nextToken()
	}

	if !p.expectPeek(SYMBOL) {
		return nil
	}

	if isGameResult(p.currToken.TokenLiteral()) {
		return move
	}

	move.MoveWhite = p.currToken.TokenLiteral()
	p.nextToken()

	if isGameResult(p.currToken.TokenLiteral()) {
		return move
	}

	move.MoveBlack = p.currToken.TokenLiteral()
	p.nextToken()

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
