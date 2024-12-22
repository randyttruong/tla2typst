package scanner

type TokenType int

const (
	KEYWORD TokenType = iota
	IDENTIFIER
	LITERAL
	OPERATOR
	SPECIAL
	NONE
)

type Token struct {
	tokenType TokenType
	value     string
	pos       int
}

func (t *Token) TokenType() TokenType {
	return t.tokenType
}

func (t *Token) GetValue() string {
	return t.value
}

func (t *Token) GetPos() int {
	return t.pos
}
