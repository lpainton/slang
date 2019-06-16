package main

import (
	"fmt"

	"github.com/slang/repl"
)

func main() {
	r := repl.New()
	err := r.Run()
	fmt.Printf("error in main: %q\n", err)
}
