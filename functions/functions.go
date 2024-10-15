package main

import "fmt"

func addTask(tasks []string, task string) []string {
	tasks = append(tasks, task)
	fmt.Println("Task added : ", task)
	return tasks
}

func ViewTask(tasks []string) {

	if len(tasks) == 0 {
		fmt.Println("There are no tasks available")
		return
	}

	fmt.Println("Your tasks are : ")
	for index, task := range tasks {
		fmt.Println(index+1, ": ", task)
	}
}

func DeleteTask(taskNo int, tasks []string) ([]string, bool) {

	if taskNo <= 0 || taskNo > len(tasks) {
		fmt.Println("Invalid task no entered...")
		return tasks, false
	}

	task := tasks[taskNo-1]
	tasks = append(tasks[:taskNo-1], tasks[taskNo:]...)

	fmt.Println("Deleted task : ", task)

	return tasks, true
}

func main() {
	var tasks []string
	for {
		fmt.Println("Enter your choice: ")
		fmt.Println("1 to Add task")
		fmt.Println("2 to View task")
		fmt.Println("3 to delete task")
		fmt.Println("4 to exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter the task : ")
			var task string
			fmt.Scanln(&task)
			tasks = addTask(tasks, task)

		case 2:
			ViewTask(tasks)

		case 3:
			var taskNo int
			fmt.Print("Enter the task number to be deleted : ")
			fmt.Scanln(&taskNo)
			var res bool
			tasks, res = DeleteTask(taskNo, tasks)

			if res {
				fmt.Println("Deletion is successfull")
			} else {
				fmt.Println("Deletion failed")
			}

		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice\nPlease enter a valid choice")
		}
	}
}
