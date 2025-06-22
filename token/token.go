package token

type TokenType string

type Token struct {
	Type  TokenType
	Value int
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF               = "EOF"

	INTEGER = "INTEGER"

	PLUS  = "+"
	MINUS = "-"
	MUL   = "*"
	DIV   = "/"

	LPAREN = "("
	RPAREN = ")"
)
