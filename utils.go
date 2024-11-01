package pgn

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isSpecialChar(ch byte) bool {
	switch ch {
	case '_', '+', '#', '=', ':', '-', '/':
		return true
	default:
		return false
	}
}

func isDigitsOnly(s string) bool {
    for _, c := range s {
        if c < '0' || c > '9' {
            return false
        }
    }
    return true
}
