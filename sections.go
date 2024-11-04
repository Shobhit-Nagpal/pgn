package pgn

import "fmt"

type stmt interface {
	Type() string
}

// Move

type Move struct {
	MoveNumber       int
	MoveWhite        string
	MoveBlack        string
	WhiteAnnotations []string
	BlackAnnotations []string
}

func (m Move) Number() int {
	return m.MoveNumber
}

func (m Move) White() string {
	return m.MoveWhite
}

func (m Move) Black() string {
	return m.MoveBlack
}

func (m Move) Type() string {
	return MOVE
}

func (m Move) GetAnnotations(color string) []string {
	if color == "White" {
		return m.WhiteAnnotations
	}

	if color == "Black" {
		return m.BlackAnnotations
	}

	return []string{}
}

func (m Move) String() string {
	return fmt.Sprintf("%d. %s %s", m.MoveNumber, m.MoveWhite, m.MoveBlack)
}

//Tag Pair

type TagPair struct {
	LBracket token
	TagName  string
	TagValue string
	RBracket token
}

func (tp TagPair) Name() string {
	return tp.TagName
}

func (tp TagPair) Value() string {
	return tp.TagValue
}

func (tp TagPair) Stringify() string {
	return fmt.Sprintf("%s%s %s%s", tp.LBracket.TokenLiteral(), tp.TagName, tp.TagValue, tp.RBracket.TokenLiteral())
}

func (tp TagPair) Type() string {
	return TAG_PAIR
}

// Game Termination

type gameTermination struct {
	TerminationValue string
}

func (gt gameTermination) Value() string {
	return gt.TerminationValue
}

func (gt gameTermination) Type() string {
	return TERMINATION
}
