package scanner_test

import (
	"glox/scanner"
	"glox/token"
	"testing"
)

func TestScanTokens(t *testing.T) {
	var tests = []struct {
		source string
		tokens *[]token.Token
	}{
        {"(", nil},
        {"// a comment \n", nil},
        {"\"test string\"", nil},
        {"9", nil},
        {"9999.99", nil},
        {"ab", nil},
        {")", nil},
        {"{", nil},
        {"}", nil},
        {"()", nil},
        {".", nil},
        {";", nil},
        {"*", nil},
        {"==", nil},
        {"=", nil},
        {"=!", nil},
        {"=>", nil},
        {">=", nil},
        {"/* a comment \n\n this /* is nice */ */ 1234", nil},
        {"/* a comment \n\n this /* is nice */ */", nil},
        {"/* a comment \n\n this is nice */ 1234", nil},
    }

    for _, tt := range tests {
        lexer := scanner.NewScanner(tt.source, tt.tokens)
        lexer.ScanTokens()

        if lexer.HadError == true {
            t.Errorf("ScanTokens: source: %s", tt.source)
        }
    }

}
