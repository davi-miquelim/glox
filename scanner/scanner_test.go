package scanner_test

import (
    "testing"
    "glox/scanner"
)

func TestScanner(t * testing.T) {
    lexer := scanner.NewScanner("3", nil)
    lexer.ScanTokens()
}
