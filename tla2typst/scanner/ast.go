package scanner

import "fmt"

type LiteralType int

const (
	UNASSIGNED_LITERAL_TYPE LiteralType = iota
)

type Primary int

const (
	UNASSIGNED_PRIMARY Primary = iota
	TRUE
	FALSE
	NIL
	NUMERIC_LITERAL_TYPE
	STRING_LITERAL_TYPE
	BOOL_LITERAL_TYPE

	// DELIMITERS
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACKET
	RIGHT_BRACKET
	LEFT_BRACE
	RIGHT_BRACE

	// VARIABLE
	VAR
	IDENTIFIER_TOKEN

	// CONTROL FLOW
	IF
	ELIF
	ELSE
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
	SUCH_THAT
	IMPLIED_SIMILARITY // ~> TODO: Is this really the correct symbol?
	// + RELATIONAL
	EQ
	EQUIVALENT
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

type Decl interface {
}

type varDecl struct {
	Decl
}

type funcDecl struct {
	Decl
}

type Stmt interface {
}

type ExprStmt struct {
	Stmt

	expr Expr
}

func (p *Parser) parse() []Expr {
	stmts := []Expr{}

	decl, _ := p.declaration()

	stmts = append(stmts, decl)
	for !p.reachedEnd() {
	}

	return stmts
}

func (p *Parser) statement() (Expr, error) {
	return nil, nil
}

// CONTROL FLOW
func (p *Parser) ifStatement() (Expr, error) {
	return nil, nil
}

func (p *Parser) elifStatment() (Expr, error) {
	return nil, nil
}

func (p *Parser) elseStatement() (Expr, error) {
	return nil, nil
}

func (p *Parser) exprStatement() (Expr, error) {
	return nil, nil
}

func (p *Parser) lor() (Expr, error) {
	return nil, nil
}

func (p *Parser) land() (Expr, error) {
	return nil, nil
}

func (p *Parser) implies() (Expr, error) {
	return nil, nil
}

func (p *Parser) suchThat() (Expr, error) {
	return nil, nil
}

// alright so we will definitely need statements for conditionals

// TODO: make it so that the match functions return errors,
func (p *Parser) declaration() (Expr, error) {
	exists, err := p.matchP(VAR)

	if err != nil {
		// TODO: figure out how to throw errors
		// p.synchronize()
		return nil, nil
	}

	if exists {
		return p.varDecl(), nil
	}

	return nil, nil
}

// TODO: write this
func (p *Parser) synchronize() {
}

func (p *Parser) varDecl() Expr {
	name, _ := p.consume(IDENTIFIER_TOKEN, "Expected a variable name")

	var initializer Expr = nil

	if exists, _ := p.matchBinOp(EQ); exists {
		initializer = p.expression()
	}

	return &Var{
		name: name,
		expr: initializer,
	}
}

type Spec struct { // NOTE: This is the root of the AST
}

type Var struct {
	Expr // TODO: Change this

	name *Token // TODO: should be a string?
	expr Expr
}

type Expr interface {
	accept(v Visitor) any
}

func (p *Parser) matchBinOp(bin ...BinaryOperator) (bool, error) {

	for idx := range bin {
		if p.checkBinOp(bin[idx]) {
			p.advance()

			return true, nil
		}
	}

	return false, nil
}

func (p *Parser) matchUnOp(un ...UnaryOperator) (bool, error) { //
	for idx := range un {
		if p.checkUnOp(UnaryOperator(un[idx])) {
			p.advance()

			return true, nil
		}
	}

	return false, nil

}

func (p *Parser) matchP(P ...Primary) (bool, error) {
	for idx := range P {
		if p.checkP(P[idx]) {
			p.advance()

			return true, nil
		}
	}

	return false, nil
}

// TODO: bruh wtf is this
// why should i have to check the tokentype man, since
// i have a fuckton of different operators.
func (p *Parser) checkBinOp(tokType BinaryOperator) bool {
	if p.reachedEnd() {
		return false
	}

	return true
	// return (p.scanner.stream[p.idx].tokenType == tokType)
}

func (p *Parser) checkUnOp(tokType UnaryOperator) bool {
	if p.reachedEnd() {
		return false
	}

	return true
	// return (p.scanner.stream[p.idx].tokenType == tokType)
}

func (p *Parser) checkP(tokType Primary) bool {
	if p.reachedEnd() {
		return false
	}

	return true
}

func (p *Parser) reachedEnd() bool {
	return p.idx == len(p.scanner.stream)
}

func (p *Parser) advance() *Token {
	if !p.reachedEnd() {
		p.idx++
	}

	return p.previous()
}

func (p *Parser) consume(pType Primary, msg string) (*Token, error) {

	if p.checkP(pType) {
		return p.advance(), nil
	}

	return nil, fmt.Errorf(msg)
}

func (p *Parser) previous() *Token {
	return p.scanner.stream[p.idx-1]
}

// Precedence rules and shit
func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparison()

	for exists, _ := p.matchBinOp(EQ, EQUIVALENT, NOT_EQ); exists; {
		op := p.previous()
		right := p.comparison() // should be a excpr

		return &Binary{
			op:    op,
			left:  expr,
			right: right,
		}
	}

	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.term()

	for exists, _ := p.matchBinOp(GREATER, LESS, GREATER_EQ, LESS_EQ, LESS_LESS, GREATER_GREATER); exists; {
		op := p.previous()
		right := p.term()

		return &Binary{
			op:    op,
			left:  expr,
			right: right,
		}
	}

	return expr
}

func (p *Parser) term() Expr {
	expr := p.factor()

	for exists, _ := p.matchBinOp(); exists; {
		op := p.previous()
		right := p.factor()

		return &Binary{
			op:    op,
			left:  expr,
			right: right,
		}
	}

	return expr
}

func (p *Parser) factor() Expr {
	expr := p.unary()

	for exists, _ := p.matchBinOp(); exists; {
		op := p.previous()
		right := p.unary()

		return &Binary{
			op:    op,
			left:  expr,
			right: right,
		}
	}

	return expr
}

func (p *Parser) unary() Expr {
	if exists, _ := p.matchUnOp(); exists {
		op := p.previous()
		right := p.primary()

		return &Unary{
			op:    op,
			right: right,
		}
	}

	return p.primary()
}

func (p *Parser) primary() Expr {

	if exists, _ := p.matchP(TRUE); exists {
		return &Literal{
			value: "true", // TODO: Change this
		}
	}

	if exists, _ := p.matchP(FALSE); exists {
		return &Literal{
			value: "false", // TODO: Change this
		}
	}

	// TODO: Can you have null values in TLA+?
	// if p.matchP(NIL) {
	//  return &Literal{
	//    value: "nil",
	//  }
	// }

	if exists, _ := p.matchP(NUMERIC_LITERAL_TYPE, STRING_LITERAL_TYPE); exists {
		return &Literal{
			value: p.previous().value,
		}
	}

	if exists, _ := p.matchP(IDENTIFIER_TOKEN); exists {
		return &Var{
			name: p.previous(),
		}
	}

	if exists, _ := p.matchP(LEFT_PAREN); exists {
		expr := p.expression()

		_, err := p.consume(RIGHT_PAREN, "Expected a right parenthesis here")

		if err != nil {
			fmt.Println("Something went wrong, got %v", err) // TODO: figure out error handling lol
		}

		return &Grouping{
			expr: expr,
		}
	}

	if exists, _ := p.matchP(LEFT_BRACKET); exists {
		expr := p.expression()

		return &Grouping{
			expr: expr,
		}
	}

	if exists, _ := p.matchP(LEFT_BRACE); exists {
		expr := p.expression()

		return &Grouping{
			expr: expr,
		}
	}

	// should not get here
	return &Literal{}
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

type Binary struct {
	Expr

	op    *Token
	left  Expr
	right Expr
}

func (b *Binary) accept(v Visitor) any {
	return v.visitBinaryExpr(b)
}

type Unary struct {
	Expr

	op    *Token
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
