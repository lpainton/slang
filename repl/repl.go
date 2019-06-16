//Package repl implements a basic repl
package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/slang/lexer"
)

const prompt = ":"

//New returns a new REPL with default values
func New() *REPL {
	return &REPL{
		Lexer:  lexer.Lexer{},
		Reader: bufio.NewReader(os.Stdin),
	}
}

//REPL represents the repl component and provides a state for the channel
type REPL struct {
	lexer.Lexer
	*bufio.Reader
}

//Run starts the REPL and returns its standard in and err channels
func (r REPL) Run() error {
	for {
		fmt.Print(":")
		in, err := r.read()
		if err != nil {
			return err
		}

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
	line, err := r.ReadString('\n')
	fmt.Printf("scanned in %q\n", line)
	return line, err
}

func (r REPL) eval(in string) (string, error) {
	tok, err := r.Tokenize(in)
	if err != nil {
		return "", err
	}
	fmt.Printf("tokenizer output: %v\n", tok)
	return in, nil
}

func (r REPL) print(out string) error {
	_, err := fmt.Println(out)
	return err
}
