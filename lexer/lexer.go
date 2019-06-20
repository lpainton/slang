//Package lexer implements a lexer for slang
package lexer

import (
	"bytes"
	"fmt"
	"unicode"
)

//LexemeType represents both a class and a lexer state
type LexemeType = int

const (
	//unspecified is the default token type
	unspecified LexemeType = iota
	//identifier starts with an alphabetic character
	identifier LexemeType = iota
	//leftSeparator is a sequence of left parens only
	leftSeparator LexemeType = iota
	//rightSeparator is a sequence of right parens only
	rightSeparator LexemeType = iota
	//numberLiteral starts with a digit
	numberLiteral LexemeType = iota
	//operator starts with a symbol that isn't a paren, letter or number
	operator LexemeType = iota
)

//Lexer holds binds to methods for the lexer
type Lexer struct{}

type lexeme struct {
	typ LexemeType
	str string
}

func scan(in string) ([]lexeme, error) {
	var (
		buf   bytes.Buffer
		state LexemeType
		out   []lexeme
	)

	seq := bytes.NewBufferString(in)
	fmt.Printf("processing sequence: %q\n", seq.String())

	var i int
	for r, _, err := seq.ReadRune(); err == nil; r, _, err = seq.ReadRune() {
		switch {
		case unicode.IsSpace(r):
			switch state {
			case unspecified:
			case identifier, leftSeparator, rightSeparator, numberLiteral, operator:
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
			case identifier:
				buf.WriteRune(r)
			default:
				return nil, fmt.Errorf("unexpected letter %q at index %d", r, i)
			}
		case unicode.IsNumber(r):
			switch state {
			case unspecified:
				state = numberLiteral
				fallthrough
			case identifier, numberLiteral:
				buf.WriteRune(r)
			default:
				return nil, fmt.Errorf("unexpected number %q at index %d", r, i)
			}
		case r == '(':
			switch state {
			case unspecified:
				state = leftSeparator
				fallthrough
			case leftSeparator:
				buf.WriteRune(r)
			default:
				return nil, fmt.Errorf("unexpected left paren at index %d", i)
			}
		case r == ')':
			switch state {
			case unspecified:
				state = rightSeparator
				fallthrough
			case rightSeparator:
				buf.WriteRune(r)
			default:
				return nil, fmt.Errorf("unexpected right paren at index %d", i)
			}
		case unicode.IsSymbol(r) || unicode.IsPunct(r):
			switch r {
			case '.', '+', '-', '*', '/', '%', '=', '&', '|', '^', '!', '<', '>':
				switch state {
				case unspecified:
					state = operator
					fallthrough
				case operator:
					buf.WriteRune(r)
				default:
					return nil, fmt.Errorf("unexpected symbol %q at index %d", r, i)
				}
			default:
				return nil, fmt.Errorf("unrecognized symbol %q at index %d", r, i)
			}
		default:
			return nil, fmt.Errorf("encountered invalid character %q at index %d", r, i)
		}
		i++
	}
	return out, nil
}

//Tokenize tokenizes a string, producing a list of tokens
func (l Lexer) Tokenize(in string) ([]Token, error) {
	fmt.Printf("tokenize %s\n", in)
	lex, err := scan(in)
	if err != nil {
		return nil, err
	}
	fmt.Println(lex)
	return evaluate(lex)
}
