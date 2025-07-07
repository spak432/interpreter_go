package lexer

import (
	"interpreter/token"
	"strings"
)

type Lexer struct {
	input   string
	curPos  int
	curRune rune
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:   input,
		curPos:  0,
		curRune: []rune(input)[0],
	}
	return l
}

func (l *Lexer) advance() {
	l.curPos += 1
	if l.curPos > len([]rune(l.input))-1 {
		l.curRune = 0
	} else {
		l.curRune = []rune(l.input)[l.curPos]
	}
}

func (l *Lexer) skipWhiteSpace() {
	whitespace := " \n\t\r"
	for strings.ContainsRune(whitespace, l.curRune) && l.curPos < len([]rune(l.input))-1 {
		l.advance()
	}
}

func (l *Lexer) integer() int {
	var num int = 0
	for l.curRune >= 48 && l.curRune <= 57 {
		num += 10*num + int(l.curRune-'0')
		l.advance()
	}
	return num
}

func (l *Lexer) GetNextToken() token.Token {

	l.skipWhiteSpace()

	if l.curPos > len([]rune(l.input))-1 {
		return token.Token{token.EOF, 0}
	}
	if l.curRune >= 48 && l.curRune <= 57 {
		num := l.integer()
		return token.Token{token.INTEGER, num}
	}
	if l.curRune == '+' {
		l.advance()
		return token.Token{token.PLUS, '+'}
	} else if l.curRune == '-' {
		l.advance()
		return token.Token{token.MINUS, '-'}
	} else if l.curRune == '*' {
		l.advance()
		return token.Token{token.MUL, '*'}
	} else if l.curRune == '/' {
		l.advance()
		return token.Token{token.DIV, '/'}
	} else if l.curRune == '(' {
		l.advance()
		return token.Token{token.LPAREN, '('}
	} else if l.curRune == ')' {
		l.advance()
		return token.Token{token.RPAREN, ')'}
	}

	return token.Token{token.EOF, 0}
}
