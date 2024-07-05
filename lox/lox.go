package lox

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

var hadError = false

func main(args []string) {
    if len(args) > 1 {
        fmt.Println("Usage: glox [script]") 
        os.Exit(64)
    } else if len(args) == 1 {
        runFile(args[0])
    } else {
        runPrompt()
    }
}

func runFile(path string) {
    if hadError == true {
        os.Exit(65)
    }

    data, err := os.ReadFile(path)   

    if err != nil {
        panic(err)
    }

    run(string(data))
}

func runPrompt() {
    for {
        scanner := bufio.NewScanner(os.Stdin)
        fmt.Println("> ")

        didScan := scanner.Scan()

        if didScan == false {
            break
        }

        run(scanner.Text())
        hadError = false
    }
}

func run(source string) {
    scanner := bufio.NewScanner(strings.NewReader(source))

    for scanner.Scan() {
        fmt.Println(scanner.Text())
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

