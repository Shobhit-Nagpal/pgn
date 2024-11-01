package pgn

import (
	"fmt"
)

type TagPair struct {
	LBracket Token
	TagName  string
	TagValue string
	RBracket Token
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

// The same tag name should not appear more than once in a tag pair section.
//Probably will use a hashmap for this

/* Some tag values may be composed of a sequence of items. For example, a
consultation game may have more than one player for a given side. When this
occurs, the single character ":" (colon) appears between adjacent items.
Because of this use as an internal separator in strings, the colon should not
otherwise appear in a string.
*/
// Check for colon, get players, offer a method to get number of players and their names
