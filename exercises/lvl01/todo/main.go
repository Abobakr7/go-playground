package main

import (
	"fmt"
	"todo/todo"
)

func main() {
	store := todo.NewStore()

	store.Add("Learn Go")
	store.Add("Build Todo CLI")
	store.Add("Study slices")

	for _, t := range store.List() {
		fmt.Println(t)
	}
	
	err := store.Delete(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, t := range store.List() {
		fmt.Println(t)
	}
}
