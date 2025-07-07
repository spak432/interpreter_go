package parser

import (
	"fmt"
	"interpreter/lexer"
	"interpreter/token"
)

type Parser struct {
	lexer    *lexer.Lexer
	curToken token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:    l,
		curToken: l.GetNextToken(),
	}
	return p
}

type ASTNode interface {
	Visit() int
}

type BinOpNode struct {
	tok token.Token
	lt  ASTNode
	rt  ASTNode
}

func (b *BinOpNode) Visit() int {
	switch b.tok.Type {
	case token.PLUS:
		return b.lt.Visit() + b.rt.Visit()
	case token.MINUS:
		return b.lt.Visit() - b.rt.Visit()
	case token.MUL:
		return b.lt.Visit() * b.rt.Visit()
	case token.DIV:
		return b.lt.Visit() / b.rt.Visit()
	default:
		fmt.Println("Visit error")
		return 0
	}
}

type NumNode struct {
	tok token.Token
}

func (n *NumNode) Visit() int {
	return n.tok.Value
}

func (p *Parser) Eat(t token.TokenType) {
	if p.curToken.Type == t {
		p.curToken = p.lexer.GetNextToken()
	}
}

func (p *Parser) Factor() ASTNode {
	curTok := p.curToken
	if curTok.Type == token.LPAREN {
		p.Eat(token.LPAREN)
		result := p.Expr()
		p.Eat(token.RPAREN)
		return result
	} else {
		p.Eat(token.INTEGER)
		return &NumNode{tok: curTok}
	}
}

func (p *Parser) Term() ASTNode {
	node := p.Factor()
	curTok := p.curToken
	for curTok.Type == token.MUL || curTok.Type == token.DIV {
		if curTok.Type == token.MUL {
			p.Eat(token.MUL)
		} else {
			p.Eat(token.DIV)
		}
		node = &BinOpNode{tok: curTok, lt: node, rt: p.Factor()}
		curTok = p.curToken
	}
	return node
}

func (p *Parser) Expr() ASTNode {
	node := p.Term()
	curTok := p.curToken
	for curTok.Type == token.PLUS || curTok.Type == token.MINUS {
		if curTok.Type == token.PLUS {
			p.Eat(token.PLUS)
		} else {
			p.Eat(token.MINUS)
		}
		node = &BinOpNode{tok: curTok, lt: node, rt: p.Term()}
		curTok = p.curToken
	}
	return node
}
