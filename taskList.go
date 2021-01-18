package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) appendTask(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) removeTask(index int) {
	/**
	* ... es el operador ellipis que permite extraer todos los elementos
	* de un arreglo, ya que append solo acepta elementos de un objeto
	 */
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) markAsCompleted() {
	t.completed = true
}

func (t *task) updateName(name string) {
	t.name = name
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func main() {
	t1 := &task{
		name:        "Tarea 1",
		description: "Ejemplo de manejo de punteros ",
	}
	t2 := &task{
		name:        "Tarea 2",
		description: "Ejemplo de lista",
	}
	list := &taskList{
		tasks: []*task{
			t1, t2,
		},
	} // crea una lista definida

	// Muestra las dos tareas iniciales
	fmt.Printf("%+v\n\n", *list.tasks[0])
	fmt.Printf("%+v\n\n", *list.tasks[1])

	//Marca tarea 2 como completada
	list.tasks[1].markAsCompleted()

	// Muestra tarea 2
	fmt.Printf("%+v\n\n", *list.tasks[1])

	// Creando tarea
	t3 := &task{
		name:        "Tarea3",
		description: "Manejo de fors",
	}
	// append
	list.appendTask(t3)
	fmt.Printf("%+v\n\n", *list.tasks[2])

	// elimina tarea2
	list.removeTask(1)
	fmt.Printf("%+v\n\n", *list.tasks[1])
}
