package main

import (
	"bufio"
	"fmt"
	"interpreter/interpreter"
	"interpreter/lexer"
	"interpreter/parser"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		s.Scan()
		input := s.Text()
		inter := interpreter.NewInterpreter(parser.NewParser(lexer.NewLexer(input)))
		fmt.Println(inter.Interpret())
	}
}
