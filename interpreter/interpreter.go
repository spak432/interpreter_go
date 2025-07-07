package interpreter

import (
	"interpreter/parser"
)

type Interpreter struct {
	parser *parser.Parser
}

func NewInterpreter(p *parser.Parser) *Interpreter {
	i := &Interpreter{
		parser: p,
	}
	return i
}

func (i Interpreter) Interpret() int {
	return i.parser.Expr().Visit()
}
