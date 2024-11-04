package pgn

import (
	"fmt"
	"log"
	"strconv"
)

type parser struct {
	l *lexer

	errors []string

	currToken token
	peekToken token
}

func newParser(l *lexer) *parser {
	p := &parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *parser) ParsePGN() (*Game, error) {
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
			case *gameTermination:
				if v.Value() != game.GetTag("Result") {
					p.errors = append(p.errors, "Game termination marker does not match game result in tag pair")
				}
				game.SetResult(v.Value())
			}
		}
	}

	if len(p.Errors()) > 0 {
		return nil, fmt.Errorf("parsing errors: %v", p.Errors())
	}

	return game, nil
}

func (p *parser) parseStatement() stmt {
	switch p.currToken.Type {
	case LBRACKET:
		return p.parseTagPair()
	case INTEGER:
		return p.parseMove()
	case SYMBOL:
		if isGameResult(p.currToken.TokenLiteral()) {
			gt := &gameTermination{TerminationValue: p.currToken.TokenLiteral()}
			p.nextToken()
			return gt
		}
		return nil
	default:
		return nil
	}
}

func (p *parser) parseTagPair() *TagPair {
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

func (p *parser) parseMove() *Move {

	moveNumInt, err := strconv.Atoi(p.currToken.TokenLiteral())
	if err != nil {
		log.Fatalf("Couldn't convert string to integer for moves: %s", p.currToken.TokenLiteral())
	}

	move := &Move{
		MoveNumber:       moveNumInt,
		WhiteAnnotations: []string{},
		BlackAnnotations: []string{},
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

	for p.peekTokenIs(NAG) {
		p.nextToken()
		move.WhiteAnnotations = append(move.WhiteAnnotations, p.currToken.TokenLiteral())
	}

	p.nextToken()

	if isGameResult(p.currToken.TokenLiteral()) {
		return move
	}

	move.MoveBlack = p.currToken.TokenLiteral()

	for p.peekTokenIs(NAG) {
		p.nextToken()
		move.BlackAnnotations = append(move.BlackAnnotations, p.currToken.TokenLiteral())
	}

	p.nextToken()

	return move
}

func (p *parser) Errors() []string {
	return p.errors
}

func (p *parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *parser) currTokenIs(t tokenType) bool {
	return p.currToken.Type == t
}

func (p *parser) peekTokenIs(t tokenType) bool {
	return p.peekToken.Type == t
}

func (p *parser) peekError(t tokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead\n", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *parser) expectPeek(t tokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
