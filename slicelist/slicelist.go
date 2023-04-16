package slicelist

import "fmt"

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
	if position == (len(lista.slice) - 1) {
		lista.Append(value)
		return
	}
	if position <= len(lista.slice) {

		lista.slice = append(lista.slice[:position+1], lista.slice[position:]...)
		lista.slice[position] = value
	}
}

func (lista *Slice[T]) Remove(value T) {

	if len(lista.slice) == 0 {

		return
	}

	for i := 0; i < (len(lista.slice)); i++ {

		if lista.slice[i] == value {

			lista.slice = append(lista.slice[:i], lista.slice[i+1:]...)
		}

	}
}

func (lista *Slice[T]) String() string {

	return fmt.Sprintf("%v", lista.slice)
}

//Search busca el primer nodo que contenga el valor recibido
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
