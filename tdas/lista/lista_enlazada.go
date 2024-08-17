package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	anterior *nodoLista[T]
	actual   *nodoLista[T]
	lista    *listaEnlazada[T]
}

func crearNodo[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato: dato, siguiente: nil}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{anterior: nil, actual: l.primero, lista: l}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(dato T) {
	nodo := crearNodo(dato)
	if l.EstaVacia() {
		l.ultimo = nodo
	} else {
		nodo.siguiente = l.primero
	}
	l.primero = nodo
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodo := crearNodo(dato)
	if l.EstaVacia() {
		l.primero = nodo
	} else {
		l.ultimo.siguiente = nodo
	}
	l.ultimo = nodo
	l.largo++

}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	elem := l.VerPrimero()
	l.primero = l.primero.siguiente

	if l.primero == nil {
		l.ultimo = nil
	}
	l.largo--

	return elem
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := l.primero
	for actual != nil {
		if !visitar(actual.dato) {
			break
		}
		actual = actual.siguiente
	}

}

func (iter *iterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterListaEnlazada[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iterListaEnlazada[T]) Insertar(dato T) {
	nodo := crearNodo(dato)
	nodo.siguiente = iter.actual

	//Iterador al principio de la lista
	if iter.anterior == nil {
		iter.lista.primero = nodo //actualizo la referencia del primer elemento en la lista
	} else {
		//Iterador en el medio de la lista
		iter.anterior.siguiente = nodo
	}

	//Iterador al final de la lista
	if !iter.HaySiguiente() {
		iter.lista.ultimo = nodo
	}
	iter.actual = nodo
	iter.lista.largo++

}

func (iter *iterListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	dato := iter.VerActual()

	//Iterador al principio de la lista
	if iter.anterior == nil {
		iter.lista.primero = iter.lista.primero.siguiente //actualizo la referencia del primer elemento en la lista

	} else {
		//Iterador al medio de la lista
		iter.anterior.siguiente = iter.actual.siguiente
	}

	//Iterador al final de la lista
	if iter.actual.siguiente == nil {
		iter.lista.ultimo = iter.anterior
	}

	iter.actual = iter.actual.siguiente
	iter.lista.largo--
	return dato
}
