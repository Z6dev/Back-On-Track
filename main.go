package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"

	"github.com/Z6dev/Back-On-Track/structs"
	"github.com/Z6dev/Back-On-Track/todoutils"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Must provide Args, use \"help\" to see list of commands")
	}
	var filePath string = "todolist.json"
	todolist, err := todoutils.LoadTodos(filePath)
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {

	case "help":
		fmt.Println("Placeholder help")

	case "list":
		if len(todolist) == 0 {
			fmt.Println("Todo-list empty")
			return
		}
		fmt.Println("______________________________")
		for i := range todolist {
			fmt.Printf("Id: %d\nDescription: %s\nStatus: %s\n", todolist[i].Id, todolist[i].Description, func() string {
				if todolist[i].Status {
					return "Done"
				} else {
					return "In-Progress"
				}
			}())
			fmt.Println("______________________________")
		}

	case "add":
		todolist = append(todolist, structs.TODO{
			Id:          len(todolist),
			Description: os.Args[2],
			Status:      false,
		})

		todoutils.SaveTodos(filePath, todolist)

	case "delete":
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		todolist = slices.Delete(todolist, index, index+1)
		todoutils.SaveTodos(filePath, todolist)
	}
}
