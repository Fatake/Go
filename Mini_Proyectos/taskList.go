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

func (t *taskList) printList() {
	fmt.Printf("<--Creando funcion dentro del Struct-->\n")
	for _, task := range t.tasks {
		fmt.Println("\nNombre: ", task.name)
		fmt.Println("Description: ", task.description)
	}
}

func (t *taskList) printListCompleted() {
	fmt.Println("<-- Tareas completadas -->")
	for _, task := range t.tasks {
		if task.completed {
			fmt.Println("\nNombre: ", task.name)
			fmt.Println("Description: ", task.description)
			fmt.Println("Completed: ", task.completed)

		}
	}
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
	// fmt.Printf("%+v\n\n", *list.tasks[0])
	// fmt.Printf("%+v\n\n", *list.tasks[1])

	//Marca tarea 2 como completada
	list.tasks[1].markAsCompleted()

	// Muestra tarea 2
	//fmt.Printf("%+v\n\n", *list.tasks[1])

	// Creando tarea
	t3 := &task{
		name:        "Tarea3",
		description: "Manejo de fors",
	}
	// append
	list.appendTask(t3)
	fmt.Printf("<--Imprimiendo lista con for i-->\n")
	// for I
	for i := 0; i < len(list.tasks); i++ {
		fmt.Printf("Index:%d Tarea:%+v\n\n", i, *list.tasks[i])
	}

	fmt.Printf("<--Imprimiendo con foreach-->\n")
	// for each
	for i, tarea := range list.tasks {
		fmt.Printf("Index:%d Tarea:%+v\n\n", i, *tarea)
	}

	list.printList()

	list.printListCompleted()
}
