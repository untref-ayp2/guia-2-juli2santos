package main

import (
	"fmt"
	"guia2/queue"
)

func main() {
	q := queue.Queue{}
	q.Enqueue(1)
	q.Enqueue("hola")
	q.Enqueue("Mundo")

	fmt.Println("-----------1------------")

	v, err := q.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
	fmt.Println("-----------2------------")
	v, err = q.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
	fmt.Println("-----------3------------")
	v, err = q.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
	fmt.Println("-----------4------------")
	q.Enqueue(1)
	q.Enqueue("hola")
	q.Enqueue("Mundo")

	v, err = q.Dequeue()
	for err == nil {
		fmt.Println(v)
		v, err = q.Dequeue()
	}

	q.Enqueue(1)
	q.Enqueue(2)

	//Modificamos la cola sin usar los métodos definidos
	//	q[0] = "Hola"
	//	q[1] = "Mundo" estos dos los comento porque con la nueva implementacion no se puede acceder a la lista
	fmt.Println("-----------5------------")
	v, err = q.Dequeue()
	for err == nil {
		fmt.Printf("%s", v)
		v, err = q.Dequeue()
	}
}
