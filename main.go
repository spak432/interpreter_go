package main

import (
	"bufio"
	"fmt"
	"interpreter/interpreter"
	"interpreter/lexer"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		s.Scan()
		input := s.Text()
		inter := interpreter.NewInterpreter(lexer.NewLexer(input))
		fmt.Println(inter.Expr())
	}
}
