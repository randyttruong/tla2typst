package scanner

import "fmt"

type LiteralType int

const (
	UNASSIGNED_LITERAL_TYPE LiteralType = iota
	NUMERIC_LITERAL_TYPE
	STRING_LITERAL_TYPE
	BOOL_LITERAL_TYPE
)

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
	accept(v Visitor) any
}

type Identifier struct {
	Expr

	value string
}

func (i *Identifier) accept(v Visitor) any {
	return v.visitIdentifier(i)
}

type Literal struct {
	Expr

	value string
}

func (l *Literal) accept(v Visitor) any {
	return v.visitLiteral(l)
}

type Visitor interface {
	visitBinaryExpr(b *Binary) any
	visitUnaryExpr(u *Unary) any
	visitGrouping(g *Grouping) any
	visitComprehension(c *Comprehension) any
	visitIdentifier(i *Identifier) any
	visitLiteral(l *Literal) any
}

type Walker struct {
	Visitor
}

func (w *Walker) visitBinaryExpr(b *Binary)           {}
func (w *Walker) visitUnaryExpr(u *Unary)             {}
func (w *Walker) visitGrouping(g *Grouping)           {}
func (w *Walker) visitComprehension(c *Comprehension) {}

type PrettyPrinter struct {
	Visitor
}

func (p *PrettyPrinter) prettyPrint(op any, exprs ...Expr) string {
	final_string := ""

	final_string += fmt.Sprintf("(%v ", op)

	for idx, expr := range exprs {
		var next string

		if idx == len(exprs)-1 {
			next += fmt.Sprintf("%v", expr.accept(p))
		} else {
			next += fmt.Sprintf("%v ", expr.accept(p))
		}

		final_string += next
	}

	final_string += fmt.Sprintf(")")

	return final_string
}

func (p *PrettyPrinter) visitBinaryExpr(b *Binary) any {
	return p.prettyPrint(b.op, b.left, b.right)
}

func (p *PrettyPrinter) visitUnaryExpr(u *Unary) any {
	return p.prettyPrint(u.op, u.right)
}
func (p *PrettyPrinter) visitGrouping(g *Grouping) any {
	return p.prettyPrint("grouping: ", g.expr)
}
func (p *PrettyPrinter) visitComprehension(c *Comprehension) any {
	return p.prettyPrint("this is a comprehension lol i dont even know what to put here just yet")
}

func (p *PrettyPrinter) visitIdentifier(i *Identifier) any {
	return i.value
}
func (p *PrettyPrinter) visitLiteral(l *Literal) any {
	return fmt.Sprintf("%v", l.value)
}

type Binary struct {
	Expr

	op    string
	left  Expr
	right Expr
}

func (b *Binary) accept(v Visitor) any {
	return v.visitBinaryExpr(b)
}

type Unary struct {
	Expr

	op    string
	right Expr
}

func (u *Unary) accept(v Visitor) any {
	return v.visitUnaryExpr(u)
}

type Grouping struct {
	Expr
	expr Expr
}

func (g *Grouping) accept(v Visitor) any {
	return v.visitGrouping(g)
}

type Comprehension struct {
	Expr
}

func (c *Comprehension) accept(v Visitor) any {
	return v.visitComprehension(c)
}
