package main

import (
	"encoding/json"
	"fmt"
	"os"
	"io"
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
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		if err == io.EOF {
			return []Task{}, nil
		}
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
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	var maxID int
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	newTask := Task{
		ID:    maxID + 1,
		Title: title,
		Done:  false,
	}

	tasks = append(tasks, newTask)
	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

func ListTasks() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	for _, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[x]"
		}
		fmt.Printf("%d: %s %s\n", task.ID, task.Title, status)
	}
}


func CompleteTask(id int) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Task not found.")
		return
	}
	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

func DeleteTask(id int) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	newTasks := []Task{}
	found := false
	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		} else {
			found = true
		}
	}
	if !found {
		fmt.Println("Task not found.")
		return
	}
	if err := saveTasks(newTasks); err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}
