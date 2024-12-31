package scanner

type TokenType int

const (
	// KEYWORDS
	KEYWORD TokenType = iota
	// FILE-RELATED
	EOF
	// CONSTANT OPERATORS
	// + LOGIC
	TRUE
	FALSE
	BOOLEAN
	CHOOSE
	COLON
	BANG
	// + SUBSETS
	SUBSET_CONST // SUBSET (ie, the set of subsets of S)
	UNION_CONST  // UNION (ie, the union of all elements of S)
	// + FUNCTIONS
	DOMAIN // DOMAIN
	EXCEPT // EXCEPT
	// DELIMITERS
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACKET
	RIGHT_BRACKET
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	// + VARIABLE
	VAR
	IDENTIFIER_TOKEN
	// MISCELLANEOUS CONSTRUCTS
	// + CONTROL FLOW
	IF
	THEN
	ELSE
	CASE
	OTHER
	LET
	IN // IN
	// + ACTION OPERATORS
	ENABLED
	UNCHANGED
	// USER-DEFINED OPERATORS
	SLASH_SLASH
	HAT_HAT
	// ARROW_HYPH // FIXME: What is this again?
	AMPERSAND
	DOUBLE_AMPERSAND
	PIPE
	DOUBLE_MOD
	DOLLAR
	DOLLAR_DOLLAR
	AT_AT
	HASH_HASH
	QUESTION
	QUESTION_QUESTION
	BANG_BANG
	DOUBLE_COLON_EQ // ::= // TODO: I am not sure how this one is interpreted, will find later
	// STANDARD OPERATORS
	PLUS     // +
	SUB      // -
	ASTERISK // *
	SLASH    // /
	EQUAL    // = (ie, for equality)
	// EXP -- later implemented in typeset symbols
	RANGE // ..
	MOD   // %
	// STANDARD NAMES
	NATURALS      // Naturals
	INTEGERS      // Integers
	REALS         // Reals
	SEQUENCES     // Sequences
	FINITE_SETS   // FiniteSets
	BAGS          // Bags
	REAL_TIME     // RealTime
	TLC           // TLC
	JSON          // Json
	RANDOMIZATION // Randomization
	// STANDARD MODULE OPERATORS
	// + Naturals, Integers, Reals
	NAT  // Nat
	REAL // Real
	INT  // Int
	INF  // Inf
	// + Sequences
	HEAD       // Head
	SELECT_SEQ // SelectSeq
	SUBSEQ     // SubSeq
	APPEND     // Append
	LEN        // Len
	SEQ        // Seq
	TAIL       // Tail
	// + FiniteSets
	IS_FINITE_SET // IsFiniteSet
	CARDINALITY   // Cardinality
	// + Bags
	BAG_IN          // BagIn
	COPIES_IN       // CopiesIn
	SUB_BAG         // SubBag
	BAG_OF_ALL      // BagOfAll
	EMPTY_BAG       // EmptyBag
	BAG_TO_SET      // BagToSet
	IS_A_BAT        // IsABag
	BAG_CARDINALITY // BagCardinality
	BAG_UNION       // BagUnion
	SET_TO_BAG      // SetToBag
	// + RealTime
	RT_BOUND // RTBound
	RT_NOW   // RTNow
	NOW      // now
	// + TLC
	COLON_ARR    // :>
	DOUBLE_AT    // @@
	PRINT        // Print
	ASSERT       // Assert
	JAVA_TIME    // JavaTime
	PERMUTATIONS // Permutations
	SORT_SEQ     // SortSeq
	// TYPESET SYMBOLS
	AND                // 1. /\\ | \land
	OR                 // 2. \\/ | \lor
	IMPLIES            // 3. =>
	NOT                // 4. ~ | \lnot | \neg
	EQUIVALENT         // 5. <=> | \equiv
	DECL               // 6. ==
	IN_TYPESET         // 7. \in
	NOT_IN             // 8. \notin
	NOT_EQ             // 9. # | /=
	TUP_LEFT           // 10. <<
	TUP_RIGHT          // 11. >>
	SQUARE             // 12. []
	LESS               // 13. <
	GREATER            // 14. >
	DIAMOND            // 15. <>
	LESS_EQ            // 16. \leq | =< | <=
	GREATER_EQ         // 17. \geq | >=
	WEAK_IMPLIES       // 18. ~>
	LESS_LESS          // 19. \ll
	GREATER_GREATER    // 20. \gg
	ARROW_SUP_POSITIVE // 21. -+->
	PREC               // 22. \prec
	SUCC               // 23. \succ
	PIPE_IMPLIES       // 24. |->
	PREC_EQ            // 25. \preceq
	SUCC_EQ            // 26. \succeq
	DIVISION           // 27. \div
	SUBSET_EQ          // 28. \subseteq
	SUPSET_EQ          // 29. \supseteq
	CDOT               // 30. \cdot
	SUBSET             // 31. \subset
	SUPSET             // 32. \supset
	CIRCLE             // 33. \o | \circ
	SQ_SUBSET          // 34. \sqsubset
	SQ_SUPSET          // 35. \sqsupset
	BULLET             // 36. \bullet
	SQ_SUBSET_EQ       // 37. \sqsubseteq
	SQ_SUPSET_EQ       // 38. \sqsupseteq
	STAR               // 39. \star
	BAR_HYPH           // 40. |-
	HYPH_BAR           // 41. -|
	BIG_CIRCLE         // 42. \bigcircle
	BAR_EQ             // 43. |=
	EQ_BAR             // 44. =|
	SIMILAR            // 45. \sim
	RIGHT_ARR          // 46. ->
	LEFT_ARR           // 47. <-
	SIMILAR_EQ         // 48. \simeq
	CAP                // 49. \cap | \intersect
	CUP                // 50. \cup | \union
	ASYMP              // 51. \asymp
	SQ_CAP             // 52. \sqcap
	SQ_CUP             // 53. \sqcup
	APPROX             // 54. \approx
	O_PLUS             // 55. (+) | \oplus
	U_PLUS             // 56. \uplus
	CONG               // 57. \cong
	O_MINUS            // 58. (-) | \ominus
	X                  // 59. \X | \times
	DOT_EQ             // 60. \doteq
	O_DOT              // 61. (.) | \odot
	WR                 // 62. \wr
	EXP                // 63. ^ (ie, x^y)
	O_TIMES            // 64. (\X) | \otimes
	PROPTO             // 65. \propto
	// 66. ^ (ie, x^+, which should already be handled...)
	O_SLASH        // 67. (/)  | \oslash
	STRING_LITERAL // 68. "" (ie, string literals)
	// 69. ^ (ie, x^*, which should be handled)
	EXISTS      // 70. \E
	FORALL      // 71. \A
	HASH        // 72. # for ^ (ie, x^#)
	BOLD_EXISTS // 73. \EE
	BOLD_FORALL // 74. \AA
	PRIME       // 75. '
	BRACKET_V   // 76. ]_v
	TUP_V       // 77. >>_v
	WF_V        // 78. WF_v
	SF          // 79. SF_v
	BAR         // 80. ------ (where length 4+)
	BOT_BAR     // 81. ====== (where length is 4+)

	IDENTIFIER
	NUM_LITERAL
	OPERATOR
	SPECIAL
	DELIMITER
	BLOCK_COMMENT // TODO: Finish comment lexing
	INLINE_COMMENT
	UNASSIGNED
)

// NOTE: This is probably an unsustainable way of
// adding keywords, operators, etc., so in the future,
// it makes sense to add a loader to autogenerate the maps.
var (
	KEYWORDS map[string]bool = map[string]bool{
		"MODULE":    true,
		"EXTENDS":   true,
		"CONSTANT":  true,
		"VARIABLES": true,
	}

	OPERATORS map[string]bool = map[string]bool{
		// CONSTANT OPERATORS
		// + LOGIC
		"TRUE":    true,
		"FALSE":   true,
		"BOOLEAN": true,
		"CHOOSE":  true,
		":":       true,
		"!":       true,
		// + SUBSETS
		"SUBSET": true,
		"UNION":  true,
		// + FUNCTIONS
		"DOMAIN": true,
		"EXCEPT": true,
		// MISCELLANEOUS CONSTRUCTS
		// + CONTROL FLOW
		"IF":    true,
		"THEN":  true,
		"ELSE":  true,
		"CASE":  true,
		"OTHER": true,
		"LET":   true,
		"IN":    true,
		// + ACTION OPERATORS
		"ENABLED":   true,
		"UNCHANGED": true,
		// USER-DEFINED OPERATORS
		"//":  true,
		"^^":  true,
		"&":   true,
		"&&":  true,
		"|":   true,
		"%%":  true,
		"$":   true,
		"$$":  true,
		"##":  true,
		"?":   true,
		"??":  true,
		"!!":  true,
		"::=": true,
		// STANDARD OPERATORS
		"+": true,
		"-": true,
		"*": true,
		// "/":   true, // NOTE: This is located in specials to avoid misparse with /\\
		"=":   true,
		"..":  true,
		"...": true,
		"%":   true,
		// STANDARD MODULE NAMES
		"Naturals":      true,
		"Integers":      true,
		"Reals":         true,
		"Sequences":     true,
		"FiniteSets":    true,
		"Bags":          true,
		"RealTime":      true,
		"TLC":           true,
		"Json":          true,
		"Randomization": true,
		// + Naturals, Integers, Reals
		"Nat":  true,
		"Real": true,
		"Int":  true,
		"Inf":  true,
		// + Sequences
		"Head":      true,
		"SelectSeq": true,
		"SubSeq":    true,
		"Append":    true,
		"Len":       true,
		"Seq":       true,
		"Tail":      true,
		// + FiniteSets
		"IsFiniteSet": true,
		"Cardinality": true,
		// + Bags
		"BagIn":          true,
		"CopiesIn":       true,
		"SubBag":         true,
		"BagOfAll":       true,
		"EmptyBag":       true,
		"BagToSet":       true,
		"IsABag":         true,
		"BagCardinality": true,
		"BagUnion":       true,
		"SetToBag":       true,
		// + RealTime
		"RTBound": true,
		"RTNow":   true,
		"now":     true,
		// + TLC
		"<:":           true,
		":>":           true,
		"@@":           true, // NOTE: @@ also referenced in User-Defined Operator Symbols (and is omitted)
		"Print":        true,
		"Assert":       true,
		"JavaTime":     true,
		"Permutations": true,
		"SortSeq":      true,
		// TYPESET SYMBOLS
		"/\\":          true,
		"\\land":       true,
		"\\/":          true,
		"\\lor":        true,
		"=>":           true,
		"~":            true,
		"\\lnot":       true,
		"\\neg":        true,
		"<=>":          true,
		"\\equiv":      true,
		"==":           true,
		"\\in":         true,
		"\\notin":      true,
		"#":            true,
		"/=":           true,
		"<<":           true,
		">>":           true,
		"[]":           true,
		"<":            true,
		">":            true,
		"<>":           true,
		"\\leq":        true,
		"=<":           true,
		"<=":           true,
		"\\geq":        true,
		">=":           true,
		"~>":           true,
		"\\ll":         true,
		"\\gg":         true,
		"-+->":         true,
		"\\prec":       true,
		"\\succ":       true,
		"|->":          true,
		"\\preceq":     true,
		"\\succeq":     true,
		"\\div":        true,
		"\\subseteq":   true,
		"\\supseteq":   true,
		"\\cdot":       true,
		"\\subset":     true,
		"\\supset":     true,
		"\\o":          true,
		"\\circ":       true,
		"\\sqsubset":   true,
		"\\sqsupset":   true,
		"\\bullet":     true,
		"\\sqsubseteq": true,
		"\\sqsupseteq": true,
		"\\star":       true,
		"|-":           true,
		"-|":           true,
		"\\bigcircle":  true,
		"|=":           true,
		"=|":           true,
		"\\sim":        true,
		"->":           true,
		"<-":           true,
		"\\simeq":      true,
		"\\cap":        true,
		"\\intersect":  true,
		"\\cup":        true,
		"\\union":      true,
		"\\asymp":      true,
		"\\sqcap":      true,
		"\\sqcup":      true,
		"\\approx":     true,
		"(+)":          true,
		"\\oplus":      true,
		"\\uplus":      true,
		"\\cong":       true,
		"\\(-)":        true,
		"\\ominus":     true,
		"\\X":          true,
		"\\times":      true,
		"\\doteq":      true,
		"\\(.)":        true,
		"\\odot":       true,
		"\\wr":         true,
		"^":            true,
		"(\\X)":        true,
		"\\otimes":     true,
		"\\propto":     true,
		"\\(/)":        true,
		"\\oslash":     true,
		"\\E":          true,
		"\\A":          true,
		"\\EE":         true,
		"\\AA":         true,
		"'":            true,
		"_v":           true,
		"WF":           true,
		"SF":           true,
		"----":         true,
		"====":         true,
		"\\setminus":   true,
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
		"/":  true,
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
