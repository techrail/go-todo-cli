package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/techrail/todo-cli/todo"
)

func main() {
	fmt.Printf("This is a CLI app to manage todo lists. It is written in golang.")

	f, err := os.OpenFile("/Users/vaibhavkaushal/code/Techrail/ToDo/todos.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Error when opening file: %v", err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("Some fatal error: %v", err)
		}
	}(f)

	// Read the contents of the existing file...
	dat, err := os.ReadFile("/Users/vaibhavkaushal/code/Techrail/ToDo/todos.json")
	// fmt.Print(string(dat))

	var existingTodos []todo.Todo

	// ... and try to cast it in a slice of todos
	err = json.Unmarshal(dat, &existingTodos)
	if err != nil {
		fmt.Printf("\n Fatal error when reading file: %v", err)
	}

	// ... and print them!
	for _, t := range existingTodos {
		t.Print()
		fmt.Printf("####")
	}

	// Let's now create the loop to read new ToDos and save them to the todos.json file
	todos := todo.InputLoop(existingTodos)
	for _, t := range todos {
		t.Print()
		fmt.Printf("####")
	}

	// Let's convert the list of Todos to JSON
	combinedTodos := append(existingTodos, todos...)
	todoJsonBytes, err := json.MarshalIndent(combinedTodos, "", "    ")
	if err != nil {
		fmt.Printf("Error when marshalling to JSON: %v", err)
		return
	}

	todoJsonContents := string(todoJsonBytes)

	//Rewriting the whole todo list into a single JSON object instead of creating new object everytime
	//Might use a different approach with seek, but this is simpler(?)s
	err = f.Truncate(0)

	if err != nil {
		fmt.Printf("\n Error while truncating the file.")
	}

	_, err = f.WriteString(todoJsonContents)
	if err != nil {
		fmt.Printf("Error when writing to file: %v\n", err)
	}
	fmt.Printf("Wrote Todo contents to file")

	// t.Print()
}
