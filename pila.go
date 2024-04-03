package guia3

import (
	"cmp"
	"errors"

	sl "github.com/untref-ayp2/data-structures/lists/single-linked-list"
)

type Pila[T cmp.Ordered] struct {
	// Implementar
}

func NewPila[T cmp.Ordered]() *Pila[T] {
	// Implementar
}

func (p *Pila[T]) Push(dato T) {
	// Implementar
}

func (p *Pila[T]) Pop() (T, error) {
	// Implementar
}

func (p *Pila[T]) Top() (T, error) {
	// Implementar
}

func (p *Pila[T]) IsEmpty() bool {
	// Implementar
}


