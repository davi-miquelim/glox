package scanner_test

import (
    "testing"
    "glox/scanner"
)

func TestScanTokens(t * testing.T) {
    lexer := scanner.NewScanner("3", nil)
    lexer.ScanTokens()

    if scanner.HadError == true {
        t.Errorf("ScanTokens")
    }
}
