//
//
//
//https://tutorialedge.net/courses/go-data-structures-course/21-implementing-queues-in-go/
package main

import (
	"fmt"
)

type Queue struct {
	Elements []*Vertex
}

func (q *Queue) Enqueue(elem *Vertex) {
	q.Elements = append(q.Elements, elem)
}

func (q *Queue) Dequeue() []*Vertex {
	for i := 1; i < q.GetLength(); i++ {
		q.Elements[i-1] = q.Elements[i]
	}

	if q.GetLength() >= 1 {
		q.Elements = q.Elements[:len(q.Elements)-1]
	}
	return q.Elements
}

func (q *Queue) Peek() *Vertex {
	if len(q.Elements) == 0 {
		err := fmt.Errorf("empty queue")
		fmt.Println(err.Error())
		return nil
	}
	return q.Elements[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue) GetLength() int {
	return len(q.Elements)
}

func (q *Queue) Print() {

	for _, e := range q.Elements {
		fmt.Printf(" %v ", e.nodeId)
	}
	fmt.Println()
}
