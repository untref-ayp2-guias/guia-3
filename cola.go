package guia3

import (
	dl "github.com/untref-ayp2/data-structures/lists/double_linked_list"
	"github.com/untref-ayp2/data-structures/types"
)

type Cola[T types.Ordered] struct {
	lista *dl.DoubleLinkedList[T]
}

func NewCola[T types.Ordered]() *Cola[T] {
	// Implementar
	return nil
}

func (c *Cola[T]) Enqueue(dato T) {
	// Implementar
}

func (c *Cola[T]) Dequeue() (T, error) {
	// Implementar
	var x T
	return x, nil

}

func (c *Cola[T]) Front() (T, error) {
	// Implementar
	var x T
	return x, nil
}

func (c *Cola[T]) IsEmpty() bool {
	// Implementar
	return false
}
