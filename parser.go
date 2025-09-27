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
				if v != nil {
					game.SetTag(v.Name(), v.Value())
				}
			case *Move:
				if v != nil {
					game.SetMove(v.Number(), v)
				}
			case *gameTermination:
				if v != nil {
					if v.Value() != game.GetTag("Result") {
						p.errors = append(p.errors, "Game termination marker does not match game result in tag pair")
					}
					game.SetResult(v.Value())
				}
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
	case INTEGER, LBRACE, SEMICOLON:
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

	if p.currTokenIs(LBRACE) || p.currTokenIs(SEMICOLON) {
		p.parseComments()
	}

	if p.peekTokenIs(EOF) {
		p.nextToken()
		return nil
	}

	moveNumInt, err := strconv.Atoi(p.currToken.TokenLiteral())
	if err != nil {
		errMsg := fmt.Sprintf("Could not convert string to integer for moves: %s", p.currToken.TokenLiteral())
		p.errors = append(p.errors, errMsg)
		moveNumInt = -1
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

	if p.peekTokenIs(LBRACE) || p.peekTokenIs(SEMICOLON) {
		p.nextToken()
		p.parseComments()
	}

	if !p.expectPeek(SYMBOL) {
		return nil
	}

	if isGameResult(p.currToken.TokenLiteral()) {
		return move
	}

	move.MoveWhite = p.currToken.TokenLiteral()

	if p.peekTokenIs(LBRACE) || p.peekTokenIs(SEMICOLON) {
		p.nextToken()
		p.parseComments()
	}

	for p.peekTokenIs(NAG) {
		p.nextToken()
		move.WhiteAnnotations = append(move.WhiteAnnotations, p.currToken.TokenLiteral())
	}

	p.nextToken()

	if p.currTokenIs(LBRACE) || p.currTokenIs(SEMICOLON) {
		p.parseComments()
	}

	if isGameResult(p.currToken.TokenLiteral()) {
		return move
	}

	move.MoveBlack = p.currToken.TokenLiteral()

	for p.peekTokenIs(NAG) {
		p.nextToken()
		move.BlackAnnotations = append(move.BlackAnnotations, p.currToken.TokenLiteral())
	}

	p.nextToken()

	if p.currTokenIs(LBRACE) || p.currTokenIs(SEMICOLON) {
		p.parseComments()
	}

	if p.currTokenIs(NEWLINE) {
		p.nextToken()
	}

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

func (p *parser) parseComments() {
	parsingComments := true

	for parsingComments {
		if p.currTokenIs(LBRACE) {
			p.parseComment()
		} else if p.currTokenIs(SEMICOLON) {
			p.parseRestOfLineComment()
		} else {
			parsingComments = false
		}
	}
}

func (p *parser) parseComment() {
	// Current token is LBRACE
	nestedBracesCount := 0
	for !p.peekTokenIs(RBRACE) || nestedBracesCount != 0 {
		if p.peekTokenIs(LBRACE) {
			nestedBracesCount++
		}

		if p.peekTokenIs(RBRACE) {
			nestedBracesCount--
		}

		p.nextToken()
	}

	p.nextToken()
}

func (p *parser) parseRestOfLineComment() {
	// Current token is SEMICOLON
	p.l.SetReadNewLine(true)
	for !p.currTokenIs(NEWLINE) {
		p.nextToken()
	}
	p.l.SetReadNewLine(false)
}
