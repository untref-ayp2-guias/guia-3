package guia3

import (
	"cmp"
)

type Nodo[T cmp.Ordered] struct {
	dato T
	sig  *Nodo[T]
	prev *Nodo[T]
}

func NuevoNodo[T cmp.Ordered](dato T) *Nodo[T] {
	// Implementar
}

func (n *Nodo[T]) Data() T {
	// Implementar
}

func (n *Nodo[T]) Next() *Nodo[T] {
	// Implementar
}

func (n *Nodo[T]) Prev() *Nodo[T] {
	// Implementar
}

func (n *Nodo[T]) SetData(dato T) {
	// Implementar
}

func (n *Nodo[T]) SetNext(sig *Nodo[T]) {
	// Implementar
}

func (n *Nodo[T]) SetPrev(prev *Nodo[T]) {
	// Implementar
}

func (n *Nodo[T]) InsertAfter(dato T) {
	// Implementar
}

func (n *Nodo[T]) InsertBefore(dato T) {
	// Implementar
}

func (n *Nodo[T]) Remove() {
	// Implementar
}

func (n *Nodo[T]) RemoveNext() {
	// Implementar
}

func (n *Nodo[T]) RemovePrev() {
	// Implementar
}

