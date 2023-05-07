package main

import (
	"fmt"
	"guia3/slicelist"
)

func main() {
	//l := linkedlist.NewLinkedList[int]()
	// fmt.Println("Agregamos 1, 2 y 3 al final de la lista")
	// l.Append(1)
	// l.Append(2)
	// l.Append(3)
	// fmt.Println(l)
	// fmt.Println("Agregamos 0 al inicio de la lista")
	// l.Prepend(0)
	// fmt.Println(l)
	// fmt.Println("Buscamos el numero 3")
	// fmt.Println("Se encuentra en la posicion: ", l.Search(3))

	// l.Remove(2)
	// l.Remove(0)
	// l.Remove(3)
	// fmt.Println("Buscamos el numero 3 luego de removerlo de la lista")
	// fmt.Println(l)
	// fmt.Println("Se encuentra en la posicion: ", l.Search(3))
	// l.Remove(1)
	// l.Remove(1) //No deberia hacer nada
	// l.Remove(1) //No deberia hacer nada
	// fmt.Println(l)

	// fmt.Println("Agregamos 0, 1 , 3, 4 y 5 al final de la lista")
	// l.Append(0)
	// l.Append(1)
	// l.Append(3)
	// l.Append(4)
	// l.Append(5)

	// fmt.Println(l)
	// fmt.Println("Agregamos 2 en la posicion 2")
	// l.InsertAt(2, 2)
	// fmt.Println("Agregamos 20 en la posicion 1")
	// l.InsertAt(20, 1)
	// fmt.Println(l)
	// fmt.Println(l.Size())
	// fmt.Println(l.SizeFlor())
	// l.Append(1)
	// l.Append(4)
	// l.Append(9)
	// l.Append(13)
	// l.Append(22)
	// fmt.Println(l)

	// l.Prepend(0)
	// fmt.Println(l)

	// l.InsertAt(10, 2)
	// fmt.Println(l)

	// l.Remove(22)
	// l.Remove(22)
	// l.Remove(22)
	// fmt.Println(l)

	// fmt.Println(l.Size())
	// fmt.Println(l.SizeFlor())

	// /*lista1 := slicelist.NewSliceList[int]()
	// fmt.Println("Agregamos valores la lista")
	// lista1.Append(1)
	// lista1.Append(2)
	// lista1.Prepend(3)
	// fmt.Println(lista1)
	// lista1.InsertAt(22, 1)
	// fmt.Println(lista1)
	// lista1.Remove(22)
	// fmt.Println(lista1)
	// lista1.Append(9)
	// lista1.Append(7)
	// lista1.Append(13)
	// fmt.Println(lista1)
	// lista1.Remove(9)
	// fmt.Println(lista1)
	// fmt.Println(lista1.Search(2))

	// l1 := linkedlist.NewLinkedList[int]()
	// l2 := linkedlist.NewLinkedList[int]()
	// l1.Append(0)
	// l1.Append(1)
	// l1.Append(2)
	// l1.Append(3)
	// l1.Append(4)
	// l2.Append(11)
	// fmt.Println(l1)
	// fmt.Println(l1.Size())

	// l2.Append(5)
	// l2.Append(6)
	// l2.Append(7)
	// l2.Append(8)
	// l2.Append(9)

	// fmt.Println(l2)
	// fmt.Println(l2.Size())
	// l1.ConcatenarLista(l1, l2)
	// fmt.Println(l1)

	// l3 := linkedlist.NewLinkedList[int]()
	// l3 = l3.ConcatenarLista(l1, l2)
	// fmt.Println(l3)

	// l3 = l3.IntercalarElementos(l1, l2)
	// fmt.Println(l3)

	// l4 := linkedlist.NewLinkedList[int]()
	// l4 = l4.MixList(l1, l2)
	// fmt.Println(l4)*/

	// l1 := linkedlist.NewLinkedList[int]()
	// l1.PilaPush(6)
	// l1.PilaPush(7)
	// l1.PilaPush(8)
	// l1.PilaPush(9)
	// l1.PilaPush(22)
	// fmt.Println(l1)

	// fmt.Println(l1.PilaPop())
	// fmt.Println(l1)

	// l1.PilaPush(11)
	// l1.PilaPush(12)
	// fmt.Println(l1)

	// l2 := linkedlist.NewLinkedList[int]()
	// l2.ColaEnqueue(1)
	// l2.ColaEnqueue(2)
	// l2.ColaEnqueue(3)
	// l2.ColaEnqueue(4)
	// l2.ColaEnqueue(5)
	// fmt.Println(l2)

	// fmt.Println(l2.ColaDequeue())
	// fmt.Println(l2)

	// l2.ColaEnqueue(5)
	// l2.ColaEnqueue(6)
	// fmt.Println(l2)

	// list := linkedlist.NewLinkedList[int]()
	// list.Append(1)
	// list.Append(2)
	// list.Append(3)
	// list.Append(4)
	// list.Append(5)
	// fmt.Println(list.Get(1))

	s := slicelist.NewSliceList[int]()
	s.Append(1)
	s.Append(2)
	s.Append(4)
	s.Append(5)
	s.InsertAt(3, 3)
	fmt.Println(s.String())
	s.Remove(5)
	fmt.Println(s.String())
	s.Remove(1)
	fmt.Println(s.String())

}
