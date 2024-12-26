package scanner

import "fmt"

type Visitor interface {
	visitBinaryExpr(b *Binary) any
	visitUnaryExpr(u *Unary) any
	visitGrouping(g *Grouping) any
	visitComprehension(c *Comprehension) any
	visitIdentifier(i *Identifier) any
	visitLiteral(l *Literal) any
}

// TODO: Start working on recursive descent algorithm.
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
