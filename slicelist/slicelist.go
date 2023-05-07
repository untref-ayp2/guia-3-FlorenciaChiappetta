package slicelist

import (
	"errors"
	"fmt"
)

type Slice[T comparable] struct {
	slice []T
}

func NewSliceList[T comparable]() *Slice[T] {
	return &Slice[T]{}
}

func (lista *Slice[T]) Append(value T) {
	lista.slice = append(lista.slice, value)
}

func (lista *Slice[T]) Prepend(value T) {
	lista.slice = append([]T{value}, lista.slice...) // agrega un slice a otro slice --> slice = append(slice, anotherSlice...)
}

func (lista *Slice[T]) InsertAt(value T, position int) {
	if position < 0 {
		return
	}
	if position == 0 {
		lista.Prepend(value)
		return
	}
	if position == len(lista.slice) {
		lista.Append(value)
		return
	}
	if position <= len(lista.slice) {

		lista.slice = append(lista.slice[:position], lista.slice[position-1:]...)
		lista.slice[position-1] = value
	}
}

func (lista *Slice[T]) Remove(value T) {

	// if len(lista.slice) == 0 {

	// 	return
	// }

	// for i := 0; i < (len(lista.slice)); i++ {

	// 	if lista.slice[i] == value {

	// 		lista.slice = append(lista.slice[:i], lista.slice[i+1:]...)
	// 	}

	// }

	if len(lista.slice) == 0 {
		return
	}
	for i := 0; i < len(lista.slice); i++ {

		if lista.slice[i] == value && i == 0 {
			lista.slice = append(lista.slice[1:len(lista.slice)])
		}

		if lista.slice[i] == value && i == len(lista.slice) {
			lista.slice = append(lista.slice[:len(lista.slice)])
		}

		if lista.slice[i] == value {
			lista.slice = append(lista.slice[:len(lista.slice)], lista.slice[len(lista.slice)-1:]...)
		}
	}
}

func (lista *Slice[T]) String() string {

	return fmt.Sprintf("%v", lista.slice)
}

// Search busca el primer nodo que contenga el valor recibido
// y devuelve su posición en la lista o -1 si no lo encuentra
func (lista *Slice[T]) Search(value T) int {
	if len(lista.slice) == 0 {
		return -1
	}
	for i := 0; i < (len(lista.slice)); i++ {
		if lista.slice[i] == value {
			return i
		}
	}
	return -1
}

// Get devuelve el valor del nodo en la posición recibida
// o un valor nulo si la posición es inválida
func (lista *Slice[T]) Get(posicion int) T {

	if posicion < 0 || posicion > len(lista.slice)-1 {
		var t T
		return t
	}
	return lista.slice[posicion]
}

func (lista *Slice[T]) size() int {

	return len(lista.slice)
}

type node[T comparable] struct {
	value T
	next  *node[T]
}

// newNode crea un nuevo nodo, con el valor recibido
// y el puntero al siguiente nodo en nil
func newNode[T comparable](value2 T) *node[T] {
	return &node[T]{value: value2, next: nil}
}

type OrLinkedList[T comparable] struct {
	head   *node[T] // puntero al primer nodo
	tail   *node[T] // puntero al último nodo
	tamaño int
}

func NewLinkedList[T comparable]() *OrLinkedList[T] {
	return &OrLinkedList[T]{head: nil, tail: nil, tamaño: 0}

}

func (l *OrLinkedList[T]) Insert(value2 T) error {
	NuevoNodo := newNode[T](value2)

	if l.head == nil {
		return nil
	}
	posActual := l.head
	for posActual != nil {
		if posActual.value == value2 {
			return errors.New("La posicion ya existe ")
		}

		if posActual.value < value2 && value2 < posActual.next.next.value {
			NuevoNodo.next = posActual.next
			posActual.next = NuevoNodo

		}
		posActual = posActual.next
	}
	return nil
}
