//Package lexer implements a lexer for slang
package lexer

import "bytes"

//TokenType represents a both a class and a lexer state
type TokenType = int

const (
	//UNSPECIFIED is the default token type
	UNSPECIFIED TokenType = iota
	//IDENTIFIER are name bindings that resolve to values
	IDENTIFIER TokenType = iota
	//LITERAL resolves to values directly
	LITERAL TokenType = iota
	//SEPERATOR is a paren
	SEPERATOR TokenType = iota
)

//Lexer holds binds to state and methods for the lexer
type Lexer struct {
	state TokenType
}

//Token is the final form of a lexemes
type Token struct {
	//Type is the type of the token
	Type TokenType
	//String holds the actual munched substring
	String string
}

//Tokenize tokenizes a string, producing a list of tokens
func (l Lexer) Tokenize(s string) ([]Token, error) {
	var (
		tok []Token
		buf bytes.Buffer
	)

	for _, c := range s {
		buf.WriteRune(c)
	}
	return tok, nil
}
