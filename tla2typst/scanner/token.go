package scanner

type TokenType int

const (
	// KEYWORDS
	KEYWORD TokenType = iota
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
	ARROW_HYPH
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
	LESS_EQ            // 16. \leq | =< | >=
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
	SQ_SECT            // 52. \sqcap
	SQ_CAP             // 53. \sqcap
	SQ_CUP             // 54. \sqcup
	APPROX             // 55. \approx
	O_PLUS             // 56. (+) | \oplus
	U_PLUS             // 57. \uplus
	CONG               // 58. \cong
	O_MINUS            // 59. (-) | \ominus
	X                  // 60. \X | \times
	DOT_EQ             // 61. \doteq
	O_DOT              // 62. (.) | \odot
	WR                 // 63. \wr
	EXP                // 64. ^ (ie, x^y)
	O_TIMES            // 65. (\X) | \otimes
	PROPTO             // 66. \propto
	// 67. ^ (ie, x^+, which should already be handled...)
	O_SLASH        // 68. (/)  | \oslash
	STRING_LITERAL // 69. "" (ie, string literals)
	// 70. ^ (ie, x^*, which should be handled)
	EXISTS      // 71. \E
	FORALL      // 72. \A
	HASH        // 73. # for ^ (ie, x^#)
	BOLD_EXISTS // 74. \EE
	BOLD_FORALL // 75. \AA
	PRIME       // 76. '
	BRACKET_V   // 77. ]_v
	TUP_V       // 78. >>_v
	WF_V        // 79. WF_v
	SF          // 80. SF_v
	BAR         // 81. ------ (where length 4+)
	BOT_BAR     // 82. ====== (where length is 4+)

	IDENTIFIER
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
