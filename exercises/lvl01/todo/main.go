package main

import (
	"fmt"
	"os"
	"todo/cli"
	"todo/todo"
)

func main() {
	store := todo.NewStore()
	processor := cli.NewProcessor(store)

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <command> [arguments]")
		os.Exit(1)
	}

	command := cli.Parse(os.Args[1:])

	if err := processor.Process(command); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
