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
	value string 
	pos int 
}  

func (t *Token) TokenType() TokenType { 
	return t.tokenType
} 

func (t *Token) value() string { 
	return t.value
}

func (t *Token) pos() string ( 
	return t.pos
)

