package queue

import (
	"errors"
	"guia2/stack"
)

// Queue implementa una cola genérica sobre un arreglo dinámico.
type Queue struct {
	items []any
}

// Enqueue agrega un elemento a la cola. O(1)
func (q *Queue) Enqueue(v any) {
	q.items = append(q.items, v)
}

// Dequeue saca un elemento de la cola. O(1)
func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	head := (q.items)[0]
	q.items = (q.items)[1:]
	return head, nil
}

// Front devuelve el elemento del frente de la cola. O(1)
func (q *Queue) Front() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return (q.items)[0], nil
}

// IsEmpty verifica si la cola esta vacia. O(1)
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

type QueueS struct {
	stack, stackAux stack.Stack
}

func (q *QueueS) Enqueue(v any) {
	//TODO
	q.stack.Push(v)
}

func (q *QueueS) Dequeue() (any, error) {
	//TODO
	if q.stack.IsEmpty() && q.stackAux.IsEmpty() {
		return nil, errors.New("Queue empty")
	}
	if q.stackAux.IsEmpty() {
		// este for se ejecuta mientras la pila main tenga items, los sacac y los pushea a la aux para invertirlos
		for !q.stack.IsEmpty() {
			v, _ := q.stack.Pop()
			q.stackAux.Push(v)
		}
	}

	return q.stackAux.Pop()
}

func (q *QueueS) IsEmpty() bool {
	//TODO
	return q.stack.IsEmpty() && q.stackAux.IsEmpty()
}

func (q *QueueS) Front() (any, error) {
	//TODO
	if q.stack.IsEmpty() && q.stackAux.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	if q.stackAux.IsEmpty() {
		for !q.stack.IsEmpty() {
			v, _ := q.stack.Pop()
			q.stackAux.Push(v)
		}
	}

	return q.stackAux.Top()

}
