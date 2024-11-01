package pgn

type Move struct {
	MoveNumber      int
	Period          Token
	MoveWhite       string
	MoveBlack       string
	MoveAnnotations []string
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

func (m Move) Annotations() []string {
	return m.MoveAnnotations
}

func (m Move) GetAnnotation(number int) string {
	if number > len(m.MoveAnnotations) {
		return ""
	}

	return m.MoveAnnotations[number-1]
}
