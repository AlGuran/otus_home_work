package main

import (
	"log"
	"os"
)

func main() {
	var argLen int = 3

	if len(os.Args) < argLen {
		log.Fatal("Too few arguments")
	}

	var directory string = os.Args[1]
	var envArgs []string = os.Args[2:]

	env, err := ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(RunCmd(envArgs, env))
}
