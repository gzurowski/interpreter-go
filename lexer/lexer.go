package lexer

import "github.com/gzurowski/monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (current character)
	readPosition int  // current reading position in input (after current character)
	ch           byte // current character
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.readWhile(isWhitespace)

	switch l.ch {
	case '=':
		tok = token.New(token.ASSIGN, string(l.ch))
	case '+':
		tok = token.New(token.PLUS, string(l.ch))
	case '(':
		tok = token.New(token.LPAREN, string(l.ch))
	case ')':
		tok = token.New(token.RPAREN, string(l.ch))
	case '{':
		tok = token.New(token.LBRACE, string(l.ch))
	case '}':
		tok = token.New(token.RBRACE, string(l.ch))
	case ',':
		tok = token.New(token.COMMA, string(l.ch))
	case ';':
		tok = token.New(token.SEMICOLON, string(l.ch))
	case 0:
		tok = token.New(token.EOF, "")
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readWhile(isLetter)
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readWhile(isDigit)
			return tok
		} else {
			tok = token.New(token.ILLEGAL, string(l.ch))
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readWhile(predicte func(byte) bool) string {
	pos := l.position
	for predicte(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}
