package linkedlist

import (
	"errors"
	"fmt"
)

// node es el nodo de la lista enlazada
// contiene un valor y un puntero al siguiente nodo
// el valor es de tipo genérico, comparable
type node[T any] struct {
	value T
	next  *node[T]
}

// newNode crea un nuevo nodo, con el valor recibido
// y el puntero al siguiente nodo en nil
func newNode[T comparable](value T) *node[T] {
	return &node[T]{value: value, next: nil}
}

/************************************************************/

// LinkedList es la lista enlazada simple
// contiene punteros al primer nodo y al último
type LinkedList[T comparable] struct {
	head   *node[T] // puntero al primer nodo
	tail   *node[T] // puntero al último nodo
	tamaño int
}

// NewLinkedList crea una nueva lista enlazada, vacía
// En nuestra implementación representamos la lista vacía
// con un puntero al primer nodo en nil
// y un puntero al último nodo en nil
// O(1)
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{head: nil, tail: nil}
}

// Append agrega un nuevo nodo, con el valor recibido, al final de la lista
// O(1)
func (l *LinkedList[T]) Append(value T) {
	newNode := newNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	l.tail.next = newNode
	l.tail = newNode
	l.tamaño++

}

// Prepend agrega un nuevo nodo, con el valor recibido,
// al inicio de la lista
// O(1)
func (l *LinkedList[T]) Prepend(value T) {
	newNode := newNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	newNode.next = l.head
	l.head = newNode
	l.tamaño++
}

// InsertAt agrega un nuevo nodo, con el valor recibido,
// en la posición recibida.
// Si la posición es inválida, no hace nada
// O(n)
func (l *LinkedList[T]) InsertAt(value T, position int) {
	if position < 0 {
		return
	}
	newNode := newNode(value)
	if position == 0 {
		l.Prepend(value)
		return
	}
	current := l.head
	for current != nil && position > 1 {
		current = current.next
		position--
	}
	if current == nil {
		return
	}
	newNode.next = current.next
	current.next = newNode
	l.tamaño++
}

// Remove elimina el primer nodo que contenga el valor recibido
// O(n)
func (l *LinkedList[T]) Remove(value T) {
	if l.head == nil {
		return // no hay nada que eliminar
	}
	if l.head.value == value {
		l.head = l.head.next
		l.tamaño--
		return

	}
	current := l.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			return
		}
		current = current.next

	}
	l.tamaño--
}

// String devuelve una representación en cadena de la lista
// en el formato [1 2 3].
// Se puede usar para imprimir la lista con fmt.Println
// O(n)
func (l *LinkedList[T]) String() string {
	if l.head == nil {
		return "[]"
	}
	current := l.head
	result := "["
	for current != nil {
		result += fmt.Sprintf("%v", current.value)
		if current.next != nil {
			result += " "
		}
		current = current.next
	}
	result += "]"
	return result
}

// Search busca el primer nodo que contenga el valor recibido
// y devuelve su posición en la lista o -1 si no lo encuentra
// O(n)

func (l *LinkedList[T]) Search(value T) int {
	if l.head == nil {
		return -1
	}
	current := l.head
	position := 0
	for current != nil {
		if current.value == value {
			return position
		}
		current = current.next
		position++
	}
	return -1
}

// Get devuelve el valor del nodo en la posición recibida
// o un valor nulo si la posición es inválida
// O(n)
func (l *LinkedList[T]) Get(position int) (T, error) {
	if l.head == nil {
		var t T
		return t, errors.New("Lista vacía")
	}
	current := l.head
	for current != nil && position > 1 {
		current = current.next
		position--
	}
	if current == nil {
		var t T
		return t, errors.New("Posición inválida")
	}
	return current.value, nil
}

// Size devuelve la cantidad de nodos en la lista
// O(n)
func (l *LinkedList[T]) Size() int {
	if l.head == nil {
		return 0
	}
	current := l.head
	position := 0
	for current != nil {
		current = current.next
		position++
	}
	return position
}

func (l *LinkedList[T]) SizeFlor() int {
	return l.tamaño // HAY QUE CORREGIR DONDE DECREMENTO LA VARIABLE EN EL METODO REMOVE.
}

// Escribir una función que reciba como parámetros dos listas del mismo tipo
//y devuelva una lista resultante de concatenar la segunda lista al final de la primera

func (l3 *LinkedList[T]) ConcatenarLista(l1, l2 *LinkedList[T]) *LinkedList[T] {
	l1.tail.next = l2.head
	l3 = l1
	return l3
}

// 4. Escribir una función que reciba dos listas del mismo tipo
// y devuelva la lista que resulta de intercalar uno a uno los elementos de ambas listas.
func (l *LinkedList[T]) IntercalarElementos(l1, l2 *LinkedList[T]) *LinkedList[T] {

	if l1.head == nil && l2.head == nil {
		return nil
	}
	if l1.head != nil && l2.head == nil {
		return l1
	}
	if l1.head == nil && l2.head != nil {
		return l2
	}

	posicionActual := l1.head
	posicionDestino := l2.head

	i := l1.Size()
	j := l2.Size()

	if i > j {
		for posicionActual != nil {
			l.Append(posicionActual.value)

			if posicionDestino != nil {
				l.Append(posicionDestino.value)
			}
			posicionActual = posicionActual.next
			posicionDestino = posicionDestino.next
		}
		return l
	}

	if i < j {
		for posicionDestino != nil {

			if posicionActual != nil {
				l.Append(posicionActual.value)
			}
			l.Append(posicionDestino.value)

			posicionActual = posicionActual.next
			posicionDestino = posicionDestino.next
		}
		return l
	}

	if i == j {
		for posicionActual != nil && posicionDestino != nil {
			l.Append(posicionActual.value)
			l.Append(posicionDestino.value)

			posicionActual = posicionActual.next
			posicionDestino = posicionDestino.next
		}
		return l
	}

	return nil
}

//Implementar una pila y una cola usando la lista enlazada simple

// PILA
// agrega un valor al tope de la pila
func (pila *LinkedList[T]) PilaPush(value T) {
	pila.Append(value)
}

// devuelve el valor del tope de la pila y lo elimina
func (pila *LinkedList[T]) PilaPop() T {
	posicion, _ := pila.Get(pila.Size() - 1)
	pila.Remove(posicion)
	return posicion
}

// devuelve el valor del tope de la pila
func (pila *LinkedList[T]) PilaTop() T {
	posicion, _ := pila.Get(pila.Size() - 1)
	return posicion
}

// COLA
// agrega un valor a la cola
func (cola *LinkedList[T]) ColaEnqueue(value T) {
	cola.Append(value)
}

// devuelve el valor del tope de la pila y lo elimina
func (cola *LinkedList[T]) ColaDequeue() T {
	posicion, _ := cola.Get(0)
	cola.Remove(posicion)
	return posicion
}

// devuelve el valor del tope de la pila
func (cola *LinkedList[T]) ColaFront() T {
	posicion, _ := cola.Get(0)
	return posicion
}
