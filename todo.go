package main

import (
	"encoding/json"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const dataFile = "todo.json"

func loadTasks() ([]Task, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(tasks)
}

func nextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func AddTask(title string) {
	tasks := loadTasks()

	var maxID int
	for _, task = range tasks{
		if task.ID > maxID{
			maxID = task.ID
		}
	}
	
	newTask := Task{
		ID: maxID + 1,
		Title: title
		Done: false,
	}

	tasks = append(tasks,newTask)
	saveTasks(task)
}

func ListTasks() {
    tasks := loadTasks()

    for _, task := range tasks {
        status := "[ ]"
        if task.Completed {
            status = "[x]"
        }
        fmt.Printf("%d: %s %s\n", task.ID, task.Title, status)
    }
}


func CompleteTask(id int) {
	panic("unimplemented")
}

func DeleteTask(id int) {
	panic("unimplemented")
}
