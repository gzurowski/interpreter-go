package lexer

type Lexer struct {
	input        string
	position     int  // current position in input (current character)
	readPosition int  // current reading position in input (after current character)
	ch           byte // current character
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	return lexer
}
