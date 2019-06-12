//Package lexer implements a lexer for slang
package lexer

import (
	"bytes"
	"fmt"
	"unicode"
)

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
	//decimalLiterals are numbers with a single decimal
	decimalLiteral LexemeType = iota
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
type Lexer struct{}

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
	var (
		seq   = bytes.NewBufferString(in)
		buf   bytes.Buffer
		state LexemeType
		out   []lexeme
	)

	var i int
	for r, _, err := seq.ReadRune(); err != nil; r, _, err = seq.ReadRune() {
		switch {
		case unicode.IsSpace(r):
			switch state {
			case unspecified:
			case identifier | leftSeparator | rightSeparator |
				numberLiteral | decimalLiteral:
				out = append(out, lexeme{typ: state, str: buf.String()})
				buf.Reset()
				state = unspecified
			default:
				return nil, fmt.Errorf("unexpected space at index %d", i)
			}
		case unicode.IsLetter(r):
			switch state {
			case unspecified:
				state = identifier
				fallthrough
			case identifier | stringLiteral:
				buf.WriteRune(r)
			default:
				return nil, fmt.Errorf("unexpected letter %q at index %d", r, i)
			}
		case unicode.IsNumber(r):
			switch state {
			case unspecified:
				state = numberLiteral
				fallthrough
			case identifier | stringLiteral |
				numberLiteral | decimalLiteral:
				buf.WriteRune(r)
			default:
				return nil, fmt.Errorf("unexpected number %q at index %d", r, i)
			}
		case r == '.':
			switch state {
			case numberLiteral:
				state = decimalLiteral
				fallthrough
			case stringLiteral | decimalLiteral:
				buf.WriteRune(r)
			default:
				return nil, fmt.Errorf("unexpected '.' at index %d", i)
			}
		case r == '(':
			switch state {
			case unspecified:
				state = leftSeparator
				fallthrough
			case leftSeparator | stringLiteral:
				buf.WriteRune(r)
			default:
				return nil, fmt.Errorf("unexpected left paren at index %d", i)
			}
		case r == ')':
			switch state {
			case unspecified:
				state = rightSeparator
				fallthrough
			case rightSeparator | stringLiteral:
				buf.WriteRune(r)
			default:
				return nil, fmt.Errorf("unexpected right paren at index %d", i)
			}
		case r == '"':
			switch state {
			case unspecified:
				state = stringLiteral
			case stringLiteral:

			}
		default:
			return nil, fmt.Errorf("encountered invalid character %q at index %d", r, i)
		}
		i++
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
