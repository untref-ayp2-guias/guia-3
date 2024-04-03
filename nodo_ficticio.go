package guia3

import "github.com/untref-ayp2/data-structures/types"

type Nodo[T types.Ordered] struct {
	dato T
	sig  *Nodo[T]
	prev *Nodo[T]
}

func NuevoNodo[T types.Ordered](dato T) *Nodo[T] {
	// Implementar
	return nil
}

func (n *Nodo[T]) Data() T {
	// Implementar
	var x T
	return x
}

func (n *Nodo[T]) Next() *Nodo[T] {
	// Implementar
	return nil
}

func (n *Nodo[T]) Prev() *Nodo[T] {
	// Implementar
	return nil
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
