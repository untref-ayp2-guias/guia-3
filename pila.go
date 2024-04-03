package guia3

import "github.com/untref-ayp2/data-structures/types"

type Pila[T types.Ordered] struct {
	// Implementar
}

func NewPila[T types.Ordered]() *Pila[T] {
	// Implementar
	return nil
}

func (p *Pila[T]) Push(dato T) {
	// Implementar
}

func (p *Pila[T]) Pop() (T, error) {
	// Implementar
	var x T
	return x, nil
}

func (p *Pila[T]) Top() (T, error) {
	// Implementar
	var x T
	return x, nil
}

func (p *Pila[T]) IsEmpty() bool {
	// Implementar
	return false
}
