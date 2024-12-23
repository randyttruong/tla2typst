package scanner

type TokenType int

const (
	KEYWORD TokenType = iota
	IDENTIFIER
	STRING_LITERAL
	NUM_LITERAL
	OPERATOR
	SPECIAL
	DELIMITER
	BLOCK_COMMENT // TODO: Finish comment lexing
	INLINE_COMMENT
	UNASSIGNED
)

var (
	KEYWORDS map[string]bool = map[string]bool{
		"MODULE":    true,
		"EXTENDS":   true,
		"CONSTANT":  true,
		"VARIABLES": true,
	}

	OPERATORS map[string]bool = map[string]bool{
		// + LOGICAL OPERATORS
		"/\\": true,
		"\\/": true,
		"~":   true,
		"=>":  true,
		"<=>": true,
		// + QUANTIFIERS
		"\\E": true,
		"\\A": true,
		// + SET OPERATORS
		"\\in":       true,
		"\\notin":    true,
		"\\subseteq": true,
		"\\supseteq": true,
		"\\subset":   true,
		"\\supset":   true,
		"\\union":    true,
		"\\cap":      true,
		"\\setminus": true,
		// + ARITHMETIC OPERATORS
		"+":     true,
		"-":     true,
		"*":     true,
		"\\div": true,
		"^":     true,
		// + RELATIONAL OPERATORS
		"=":  true,
		"#":  true,
		">":  true,
		"<":  true,
		"<=": true,
		">=": true,
		// + MISC OPERATORS
		"==":        true,
		"<<":        true,
		">>":        true,
		"..":        true,
		"|->":       true,
		"UNCHANGED": true,
		// + TEMPORAL LOGIC OPERATORS
		"[]": true,
		"<>": true,
		"~>": true,
		// + FUNCTIONAL OPERATORS
		"Seq":    true,
		"Head":   true,
		"Tail":   true,
		"Append": true,
		"Len":    true,
	}

	DELIMITERS = map[string]bool{
		// "(": true,
		// ")": true,
		")":  true,
		" [": true,
		"]":  true,
		"{":  true,
		"}":  true,
		",":  true,
	}

	// + SPECIALS denote ambigious characters
	SPECIALS = map[string]bool{
		// + Left parenthesis serve as both expression delimiters + block comment delimiters
		"(": true,
		// + Forward slashes are start chars for both logical AND and inline comments
		"\\": true,
	}
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
