package main

import (
	"fmt"
	"os"

	"github.com/hereisjohnny2/monkey-lang/repl"
)

func main() {
	fmt.Println("Monkey Lang REPL - v0.1.0")
	repl.Start(os.Stdin, os.Stdout)
}
