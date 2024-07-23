package lox

import (
	"bufio"
	"fmt"
	"glox/interpreter"
	"glox/parser"
	"glox/scanner"
	"os"
)

var hadError = false
var hadRuntimeError = false

func RunFile(path string) {
	if hadError == true {
	}

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error: Could not read file")
		os.Exit(2)
	}

	run(string(data))
	if hadError == true {
		os.Exit(65)
	}
	if hadRuntimeError == true {
		os.Exit(70)
	}
}

func RunPrompt() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		didScan := scanner.Scan()

		if didScan == false {
			break
		}

		run(scanner.Text())
		hadError = false
	}
}

func run(source string) {
	lexer := scanner.NewScanner(source, nil)
	tokens := lexer.ScanTokens()
	if lexer.HadError == true {
		hadError = true
	}

	loxInterpreter := interpreter.NewInterpreter()
	loxParser := parser.NewParser(tokens)
	exprTree := loxParser.Parse()

	if loxParser.HadError == true {
		hadError = true
	}

	loxInterpreter.Interpret(exprTree)
	if loxInterpreter.HadRuntimeError == true {
		hadRuntimeError = true
	}
}

func Error(line int, where *string, message string) {
	report(line, nil, message)
}

func report(line int, where *string, message string) {
	hadError = true

	if where == nil {
		err := fmt.Errorf("[line %d] Error %s", line, message)
		fmt.Println(err)
		return
	}

	err := fmt.Errorf("[line %d] Error %s: %s", line, message, *where)
	fmt.Println(err)
}
