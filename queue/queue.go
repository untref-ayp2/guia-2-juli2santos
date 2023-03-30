package queue

import "errors"

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
}

func (q *QueueS) Enqueue(v any) {
	//TODO

}

func (q *QueueS) Dequeue() (any, error) {
	//TODO
	return 0, nil
}

func (q *QueueS) IsEmpty() bool {
	//TODO
	return false
}

func (q *QueueS) Front() (any, error) {
	//TODO
	return 0, nil
}
