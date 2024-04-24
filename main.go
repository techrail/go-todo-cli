package main

import (
	`encoding/json`
	`fmt`
	`time`

	`github.com/techrail/todo-cli/todo`
)

func main() {
	fmt.Printf("This is a CLI app to manage todo lists. It is written in golang.")

	t := todo.Todo{
		Title:     "First todo ever",
		Desc:      "Just testing the app",
		Done:      true,
		CreatedAt: time.Now(),
	}

	if t.Validate() != nil {
		fmt.Printf("\nTodo is not valid. Error: %v\n", t.Validate())
	} else {
		fmt.Printf("\nTodo is valid\n")
	}

	todoBytes, err := json.Marshal(t)
	if err != nil {
		fmt.Printf("\nSomething went wrong while encoding to JSON: %v\n", err)
	}

	todoJsonString := string(todoBytes)
	fmt.Printf("\nJSON String: %v\n", todoJsonString)

	// Let's now create the loop to read new ToDos and
	t.Print()
}
