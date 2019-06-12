//Package repl implements a basic repl
package repl

import (
	"fmt"

	"github.com/slang/lexer"
)

//New returns a new REPL with default values
func New() *REPL {
	return &REPL{
		lexer.Lexer{},
	}
}

//REPL represents the repl component and provides a state for the channel
type REPL struct {
	lexer.Lexer
}

//Run starts the REPL and returns its standard in and err channels
func (r REPL) Run() error {
	for {
		in, err := r.read()
		if err != nil {
			return err
		}
		fmt.Println(in)

		out, err := r.eval(in)
		if err != nil {
			return err
		}

		err = r.print(out)
		if err != nil {
			return err
		}
	}
}

func (r REPL) read() (string, error) {
	var in string
	_, err := fmt.Scanln(&in)
	return in, err
}

func (r REPL) eval(in string) (string, error) {
	tok, err := r.Tokenize(in)
	if err != nil {
		return "", err
	}
	fmt.Println(tok)
	return in, nil
}

func (r REPL) print(out string) error {
	_, err := fmt.Println(out)
	return err
}
