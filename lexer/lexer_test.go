package lexer

import (
	"testing"

	"github.com/gzurowski/monkey/token"
	"github.com/stretchr/testify/assert"
)

func TestNextToken(t *testing.T) {
	// Arrange
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for _, tt := range tests {
		// Act
		got := lexer.NextToken()

		// Assert
		assert.Equal(t, tt.expectedType, got.Type, "wrong token type")
		assert.Equal(t, tt.expectedLiteral, got.Literal, "wrong literal")
	}
}
