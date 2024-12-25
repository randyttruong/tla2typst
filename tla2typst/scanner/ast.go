package scanner

type BinaryOperator int

const (
	UNASSIGNED_BINOP BinaryOperator = iota
	// + PRIUnaryOperator = iotaUnaryOperator = iotaMITIVES
	ASSIGN
	// + ARITHEMTIC
	ADD
	SUB
	MULTIPLY     // represented as cdot or asterisk
	FLOOR_DIVIDE // int divide
	DIVIDE       // float divide
	MOD
	EXP
	// + GROUPING
	RANGE
	// + LOGICAL
	AND
	OR
	IMPLIES
	EQUIVALENT
	SUCH_THAT
	IMPLIED_SIMILARITY // ~> TODO: Is this really the correct symbol?
	// + RELATIONAL
	GREATER
	LESS
	GREATER_EQ
	LESS_EQ
	NOT_EQ
	LESS_LESS
	GREATER_GREATER
	// + TEMPORAL
	LEADS_TO
	// + INTER SET
	SUBJUNCTION
	DISJUNCTION
	IN
	NOT_IN
	SET_MINUS
	// + INTRA SET
	SUBSET
	SUBSET_EQ
	SUPSET
	SUPSET_EQ
	SQ_SUBSET
	SQ_SUBSET_EQ
	SQ_SUPSET
	SQ_SUPSET_EQ
)

type UnaryOperator int

const (
	UNASSIGNED_UNOP UnaryOperator = iota
	// + TYPECASTING
	NAT
	REAL
	INT
	INFINITY
	// + LOGICAL
	NEGATE
	// + QUANTIFIERs
	FORALL
	EXISTS
	// + TEMPORAL
	ALWAYS
	EVENTUALLY
	WEAK_FAIRNESS
	STRONG_FAIRNESS
	// + RECORDS
	FIELD
)

type Expr struct {
}

type BinaryOperation struct {
	Expr
	op    BinaryOperator
	left  Expr
	right Expr
}

type UnaryOperation struct {
	Expr
	op    UnaryOperator
	right Expr
}

type Grouping struct {
	Expr
}

type Comprehension struct {
	Grouping
}
