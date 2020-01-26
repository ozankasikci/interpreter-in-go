package main

import (
	"fmt"
	"interpreter-in-go/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", user.Username)
	fmt.Printf("Type some code\n")
	repl.Start(os.Stdin, os.Stdout)
}
