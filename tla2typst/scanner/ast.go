package scanner

type BinaryOperator int

const (
	UNASSIGNED_BINOP BinaryOperator = iota
	// + PRIMITIVES
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

type Expr interface {
	accept(v Visitor)
}

type Visitor interface {
	visitBinaryExpr(b *Binary)
	visitUnaryExpr(u *Unary)
	visitGrouping(g *Grouping)
	visitComprehension(c *Comprehension)
}

type Walker struct {
	Visitor
}

func (w *Walker) visitBinaryExpr(b *Binary)           {}
func (w *Walker) visitUnaryExpr(u *Unary)             {}
func (w *Walker) visitGrouping(g *Grouping)           {}
func (w *Walker) visitComprehension(c *Comprehension) {}

type Binary struct {
	Expr
	op    BinaryOperator
	left  Expr
	right Expr
}

func (b *Binary) accept(v Visitor) {
	v.visitBinaryExpr(b)
}

type Unary struct {
	Expr
	op    UnaryOperator
	right Expr
}

func (u *Unary) accept(v Visitor) {
	v.visitUnaryExpr(u)
}

type Grouping struct {
	Expr
}

func (g *Grouping) accept(v Visitor) {
	v.visitGrouping(g)
}

type Comprehension struct {
	Expr
}

func (c *Comprehension) accept(v Visitor) {
	v.visitComprehension(c)
}
