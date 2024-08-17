package pila

/* Definición del struct pila proporcionado por la cátedra. */

const (
	_CAPACIDAD_INICIAL int = 4
	_CTE_REDIMENSION   int = 2
	_CTE_ACHICAR       int = 4
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, _CAPACIDAD_INICIAL), cantidad: 0}
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(elemento T) {
	if p.cantidad == len(p.datos) {
		p.redimensionar(len(p.datos) * _CTE_REDIMENSION)
	}

	p.datos[p.cantidad] = elemento
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	e := p.datos[p.cantidad-1]
	p.cantidad--

	if p.cantidad <= len(p.datos)/_CTE_ACHICAR && len(p.datos) > _CTE_ACHICAR {
		p.redimensionar(len(p.datos) / _CTE_REDIMENSION)
	}

	return e
}

func (p *pilaDinamica[T]) redimensionar(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, p.datos)
	p.datos = nuevosDatos

}
