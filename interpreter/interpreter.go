package interpreter

import (
	"interpreter/lexer"
	"interpreter/token"
)

type Interpreter struct {
	lexer    *lexer.Lexer
	curToken token.Token
}

func NewInterpreter(l *lexer.Lexer) *Interpreter {
	i := &Interpreter{
		lexer:    l,
		curToken: l.GetNextToken(),
	}
	return i
}

func (i *Interpreter) Eat(t token.TokenType) {
	if i.curToken.Type == t {
		i.curToken = i.lexer.GetNextToken()
	}
}

func (i *Interpreter) Factor() int {
	var r int
	if i.curToken.Type == token.LPAREN {
		i.Eat(token.LPAREN)
		r = i.Expr()
		i.Eat(token.RPAREN)
	} else if i.curToken.Type == token.INTEGER {
		r = i.curToken.Value
		i.Eat(token.INTEGER)
	}
	return r
}

func (i *Interpreter) Term() int {
	r := i.Factor()
	for i.curToken.Type == token.MUL || i.curToken.Type == token.DIV {
		if i.curToken.Type == token.MUL {
			i.Eat(token.MUL)
			r *= i.Factor()
		} else {
			i.Eat(token.DIV)
			r /= i.Factor()
		}
	}
	return r
}

func (i *Interpreter) Expr() int {
	r := i.Term()
	for i.curToken.Type == token.PLUS || i.curToken.Type == token.MINUS {
		if i.curToken.Type == token.PLUS {
			i.Eat(token.PLUS)
			r += i.Term()
		} else {
			i.Eat(token.MINUS)
			r -= i.Term()
		}
	}
	return r
}
