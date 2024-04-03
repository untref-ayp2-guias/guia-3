package guia3

import (
	"cmp"
	"errors"
	dl "github.com/untref-ayp2/data-structures/lists/double-linked-list"
)

type Cola[T cmp.Ordered] struct {
	lista *dl.DoubleLinkedList[T]
}

func NewCola[T cmp.Ordered]() *Cola[T] {
	// Implementar
}

func (c *Cola[T]) Enqueue(dato T) {
	// Implementar
}

func (c *Cola[T]) Dequeue() (T, error) {
	// Implementar
}

func (c *Cola[T]) Front() (T, error) {
	// Implementar
}

func (c *Cola[T]) IsEmpty() bool {
	// Implementar
}

