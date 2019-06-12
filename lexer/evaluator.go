package lexer

import "fmt"

var keywords = []string{
	"sum",
	"prod",
	"recip",
	"inv",
}

func (l *Lexer) evaluate(in []lexeme) ([]Token, error) {
	for _, l := range in {
		fmt.Printf("Lexeme %q type %q", l.str, l.typ)
	}
	return nil, nil
}
