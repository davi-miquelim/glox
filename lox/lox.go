package lox

import (
	"bufio"
	"fmt"
	"glox/scanner"
	"os"
)

var hadError = false

func RunFile(path string) {
	if hadError == true {
		os.Exit(65)
	}

	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	run(string(data))
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
	lexer.ScanTokens()

	if lexer.HadError == true {
		hadError = true
	}
}

func Error(line int, message string) {
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
