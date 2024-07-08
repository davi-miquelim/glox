package main

import (
    "fmt"
    "os"
    "glox/lox"
)


func main() {
    args := os.Args[1:]
    
    if len(args) > 1 {
        fmt.Println("Usage: glox [script]") 
        os.Exit(64)
    } else if len(args) == 1 {
        lox.RunFile(args[0])
    } else {
        lox.RunPrompt()
    }
}
