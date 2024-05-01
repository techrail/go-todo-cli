package todo

import (
	`bufio`
	`fmt`
	`log`
	`os`
	`strings`
	`time`
)

func ReadATodo() Todo {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nEnter the title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("\nEnter the description: ")
	desc, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return Todo{
		Title:     strings.Trim(title, "\n"),
		Desc:      strings.Trim(desc, "\n"),
		Done:      false,
		CreatedAt: time.Now(),
	}
}

func InputLoop(existingTodos []Todo) []Todo {
	todos := make([]Todo, 0)
	for {
		fmt.Printf("\nPlease choose an option:\n")
		fmt.Printf("1. Create a new Todo\n")
		fmt.Printf("2. Select a new Todo\n")
		fmt.Printf("0. Exit the loop\n")
		fmt.Printf("Enter your choice: ")
		var input string
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Printf("E#1VUOMM - %v\n", err)
		}
		fmt.Printf("Input was: %v\n", input)
		fmt.Printf("Email was invalid - does not have the @ character in it")

		if input == "0" {
			break
		} else if input == "1" {
			t := ReadATodo()
			// t.Print()
			todos = append(todos, t)
		} else if input == "2" {
			for i, t := range existingTodos {
				t.PrintIdAndTitle(i + 1)
			}
			var selectedTodoNumber int
			fmt.Printf("\n Please enter the ID of todo to select it: ")
			_, err := fmt.Scanf("%d", &selectedTodoNumber)
			if err != nil {
				fmt.Printf("E#1VUQF1 - %v\n", err)
			}
			selectedTodoIndex := selectedTodoNumber - 1

			option := SingleTodoActions(existingTodos[selectedTodoIndex])
			switch option {
			case "1":
				fmt.Printf("\nYou want to mark the todo `%v` as done!", existingTodos[selectedTodoIndex].Title)
			case "2":
				fmt.Printf("\nYou want to edit the title of the Todo `%v`", existingTodos[selectedTodoIndex].Title)
			default:
				fmt.Printf("\n Invalid Selection!")
			}
		} else {
			fmt.Println("You selected an invalid option. Try again")
		}
	}
	return todos
}

func SingleTodoActions(todo Todo) string {
	fmt.Printf("\n You have selected: %v", todo.Title)
	fmt.Printf("\n1. Mark as done")
	fmt.Printf("\n2. Edit the title")
	fmt.Printf("\n3. Edit the description")
	fmt.Printf("\n0. Go back to main loop")
	fmt.Printf("\nPlease select the option: ")
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		log.Fatalln("Fatal error:", err)
	}

	fmt.Printf("Input was: %v\n", input)
	return input
}
