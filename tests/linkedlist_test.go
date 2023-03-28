package tests

//Tests Lista Enlazada

import (
	"guia3/linkedlist"
	"testing"
)

func TestInsert(t *testing.T) {
	list := linkedlist.NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	
	if list.Size() != 5 {
		t.Error("Error en Append")
	}

	for i := 0; i < 5; i++ {
		if list.Get(i) != i+1 {
			t.Error("Error en Append")
		}
	}
}

func TestDelete(t *testing.T) {
	list := linkedlist.NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	
	list.Delete(0)
	list.Delete(2)
	list.Delete(3)

	if list.Size() != 2 {
		t.Error("Error en Delete")
	}

	if list.Get(0) != 2 {
		t.Error("Error en Delete")
	}

	if list.Get(1) != 4 {
		t.Error("Error en Delete")
	}
}
	

