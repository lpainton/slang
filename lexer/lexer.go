//Package lexer implements a lexer for slang
package lexer

import "fmt"

//LexemeType represents a both a class and a lexer state
type LexemeType = int

//TokenType represents the final token type
type TokenType = int

const (
	//unspecified is the default token type
	unspecified LexemeType = iota
	//identifier starts with an alphabetic character
	identifier LexemeType = iota
	//leftSeparator is a sequence of left parens only
	leftSeparator LexemeType = iota
	//rightSeparator is a sequence of right parens only
	rightSeparator LexemeType = iota
	//stringLiteral starts and ends with a quote
	stringLiteral LexemeType = iota
	//numberLiteral starts with a digit
	numberLiteral LexemeType = iota
)

const (
	//Invalid is the default token type
	Invalid TokenType = iota
	//Keyword ::= “add” | “sub” | “mul” | “div”
	Keyword TokenType = iota
	//Variable ::= <any identifier not in Keyword>
	Variable TokenType = iota
	//LeftParen ::= "("
	LeftParen TokenType = iota
	//RightParen ::= ")"
	RightParen TokenType = iota
	//String ::= <any stringLiteral>
	String TokenType = iota
	//Integer ::= "0"..."9" | <Integer> + "0"..."9"
	Integer TokenType = iota
	//Float ::= <integer> + "." + <integer>
	Float TokenType = iota
)

//Lexer holds binds to methods for the lexer
type Lexer struct {
	scanState LexemeType
}

type lexeme struct {
	typ LexemeType
	str string
}

//Token is the final product of the lexer and binds a type to a string value
type Token struct {
	Typ TokenType
	Str string
}

func (l *Lexer) scan(in string) ([]lexeme, error) {
	for _, c := range in {
		fmt.Println(c)
	}
	return nil, nil
}

//Tokenize tokenizes a string, producing a list of tokens
func (l Lexer) Tokenize(in string) ([]Token, error) {
	lex, err := l.scan(in)
	if err != nil {
		return nil, err
	}
	return l.evaluate(lex)
}
