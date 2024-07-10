package scanner

import (
	"glox/token"
	"testing"
)

func TestScanTokens(t *testing.T) {
	var tests = []struct {
		source string
		tokens *[]token.Token
		want   *[]int
	}{
		{"(", nil, &[]int{token.LeftParen, token.Eof}},
		{"// a comment \n", nil, nil},
		{"\"test string\"", nil, &[]int{token.String, token.Eof}},
		{"9", nil, &[]int{token.Number, token.Eof}},
		{"9999.99", nil, &[]int{token.Number, token.Eof}},
		{"ab", nil, &[]int{token.String, token.Eof}},
		{")", nil, &[]int{token.RightParen, token.Eof}},
		{"{", nil, &[]int{token.LeftBrace, token.Eof}},
		{"}", nil, &[]int{token.RightBrace, token.Eof}},
		{"()", nil, &[]int{token.LeftBrace, token.RightBrace, token.Eof}},
		{".", nil, &[]int{token.Dot, token.Eof}},
		{";", nil, &[]int{token.SemiColon, token.Eof}},
		{"*", nil, &[]int{token.Star, token.Eof}},
		{"==", nil, &[]int{token.EqualEqual, token.Eof}},
		{"=", nil, &[]int{token.Equal, token.Eof}},
		{"!=", nil, &[]int{token.BangEqual, token.Eof}},
		{">=", nil, &[]int{token.GreaterEqual, token.Eof}},
		{"<=", nil, &[]int{token.LessEqual, token.Eof}},
		{"/* a comment \n\n this /* is nice */ */ 1234", nil, &[]int{token.Number, token.Eof}},
		{"/* a comment \n\n this /* is nice */ */", nil, &[]int{token.Eof}},
		{"/* a comment \n\n this is nice */ 1234", nil, &[]int{token.Number, token.Eof}},
	}

	for _, tt := range tests {
		lexer := NewScanner(tt.source, tt.tokens)
		lexer.ScanTokens()
		tokens := lexer.Tokens

		if tt.want == nil {
			return
		}

		tokensValue := *tokens
		want := *tt.want
		for i := range *tt.want {
			if want[i] != tokensValue[i].TokenType {
				t.Errorf("[ScanTokens]: expected: %d, got: %d", want[i], tokensValue[i].TokenType)
			}
		}

	}
}
