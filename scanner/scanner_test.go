package scanner_test

import (
	"glox/scanner"
	"testing"
)

func TestScanTokens(t *testing.T) {
	var tests = []struct {
		source string
		tokens []interface{}
	}{
        {"9", nil},
        {"9999.99", nil},
        {"ab", nil},
        {"(", nil},
        {")", nil},
        {"{", nil},
        {"}", nil},
        {"()", nil},
        {"[", nil},
        {"]", nil},
        {".", nil},
        {";", nil},
        {"*", nil},
        {"==", nil},
        {"=!", nil},
        {"=>", nil},
        {">=", nil},
        {"// a comment \n", nil},
    }

    for _, tt := range tests {
        lexer := scanner.NewScanner(tt.source, tt.tokens)
        lexer.ScanTokens()

        if lexer.HadError == true {
            t.Errorf("ScanTokens")
        }
    }

}
