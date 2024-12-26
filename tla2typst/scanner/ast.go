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


}

}


}

}

}

}


}




		}

	}


}

}

}
}
}

}
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
