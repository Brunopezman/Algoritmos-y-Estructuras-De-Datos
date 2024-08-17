package cola

type nodoCola[T any] struct {
	dato      T
	siguiente *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func crearNodo[T any](dato T) *nodoCola[T] {
	return &nodoCola[T]{dato: dato, siguiente: nil}
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{primero: nil, ultimo: nil}
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil && c.ultimo == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(dato T) {
	nodo := crearNodo[T](dato)
	if c.EstaVacia() {
		c.primero = nodo
	} else {
		c.ultimo.siguiente = nodo
	}
	c.ultimo = nodo
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := c.primero.dato
	c.primero = c.primero.siguiente
	if c.primero == nil {
		c.ultimo = nil
	}
	return dato
}
