package main

import (
	"fmt"
	"os"
	"os/user"

	"example.com/cat_interpreter/src/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the CA+ programming language!\n",
		user.Username)
	fmt.Printf("Type some demands of my capabilities\n")
	repl.Start(os.Stdin, os.Stdout)
}
