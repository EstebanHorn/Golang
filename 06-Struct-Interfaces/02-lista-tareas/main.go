package main

import "fmt"

type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) appendTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) deleteTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

type Task struct {
	name        string
	description string
	completed   bool
}

func (t *Task) print() {
	fmt.Printf("Tarea: %s \nDescripcion: %s \nCompletado: %t \n",
		t.name, t.description, t.completed)
}

func (t *Task) toComplet() {
	t.completed = true
}

func main() {

	task1 := Task{
		name:        "Curso GO",
		description: "Completar curso go",
	}

	task2 := Task{
		name:        "Curso React",
		description: "Completar React",
		completed:   true,
	}

	task3 := Task{
		name:        "Curso HTML",
		description: "Completar Html",
	}

	list := TaskList{}

	list.appendTask(&task1)
	list.appendTask(&task2)
	list.appendTask(&task3)

	for i, task := range list.tasks {
		fmt.Print(i, " ")
		task.print()
		fmt.Print("\n")
	}
}
