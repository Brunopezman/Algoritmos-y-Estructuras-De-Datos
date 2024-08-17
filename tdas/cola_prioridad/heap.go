package cola_prioridad

const (
	_TAM_INICIAL      = 5
	_AUMENTAR         = 2
	_REDUCIR          = 0.5
	_CANTIDAD_INICIAL = 0
	_CTE_ACHICAR      = 4
)

type colaPrioridad[T comparable] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func crearArr[T comparable](tam int) []T {
	return make([]T, tam)
}

func CrearHeap[T comparable](cmp func(T, T) int) ColaPrioridad[T] {
	return &colaPrioridad[T]{crearArr[T](_TAM_INICIAL), _CANTIDAD_INICIAL, cmp}
}

func CrearHeapArr[T comparable](arreglo []T, cmp func(T, T) int) ColaPrioridad[T] {
	cp := new(colaPrioridad[T])
	cp.cmp = cmp
	if len(arreglo) > _TAM_INICIAL {
		cp.datos = make([]T, len(arreglo))
	} else {
		cp.datos = make([]T, _TAM_INICIAL)

	}

	cp.cantidad = len(arreglo)
	copy(cp.datos, arreglo)
	heapify(cp.datos, (cp.cantidad-1)/2, cmp, cp.cantidad)
	return cp
}

func downHeap[T any](arr []T, pos int, cmp func(T, T) int, cantidad int) {
	if pos >= cantidad {
		return
	}

	//Posiciones de los hijos
	hijoIzq := 2*pos + 1
	hijoDer := 2*pos + 2

	max := pos
	if hijoIzq < cantidad && cmp(arr[hijoIzq], arr[max]) > 0 {
		max = hijoIzq
	}
	if hijoDer < cantidad && cmp(arr[hijoDer], arr[max]) > 0 {
		max = hijoDer
	}
	if max != pos {
		swap(&arr[pos], &arr[max])
		downHeap(arr, max, cmp, cantidad)
	}
}

func upHeap[T any](arr []T, pos int, cmp func(T, T) int) {
	if cmp(arr[pos], arr[0]) == 0 {
		return
	}
	padre := buscarPadre(pos)
	if cmp(arr[pos], arr[padre]) > 0 {
		swap(&arr[pos], &arr[padre])
		upHeap(arr, padre, cmp)
	}
}

func heapify[T any](arr []T, pos int, cmp func(T, T) int, cantidad int) {
	for pos >= 0 {
		downHeap(arr, pos, cmp, cantidad)
		pos--
	}
}

func HeapSort[T comparable](arr []T, cmp func(T, T) int) {
	heapify(arr, (len(arr)/2)-1, cmp, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		swap(&arr[0], &arr[i])
		downHeap(arr[:i], 0, cmp, i)
	}
}

func (cp *colaPrioridad[T]) EstaVacia() bool {
	return cp.cantidad == 0
}
func (cp *colaPrioridad[T]) Encolar(elem T) {
	if len(cp.datos) == cp.cantidad {
		cp.redimensionar(len(cp.datos) * _AUMENTAR)
	}
	cp.datos[cp.cantidad] = elem
	upHeap(cp.datos, cp.cantidad, cp.cmp)
	cp.cantidad++
}

func (cp *colaPrioridad[T]) VerMax() T {
	if cp.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cp.datos[0]
}
func (cp *colaPrioridad[T]) Desencolar() T {
	if cp.EstaVacia() {
		panic("La cola esta vacia")
	}
	desencolado := cp.VerMax()
	swap(&cp.datos[0], &cp.datos[cp.cantidad-1])
	cp.cantidad--
	downHeap(cp.datos, 0, cp.cmp, cp.cantidad)

	if cp.cantidad > _TAM_INICIAL && cp.cantidad*_CTE_ACHICAR <= len(cp.datos) {
		cp.redimensionar(int(float64(len(cp.datos)) * _REDUCIR))
	}

	return desencolado
}

func (cp *colaPrioridad[T]) Cantidad() int {
	return cp.cantidad
}

func (cp *colaPrioridad[T]) redimensionar(nuevoTam int) {
	cpViejo := cp.datos
	cp.datos = make([]T, nuevoTam)
	copy(cp.datos, cpViejo)
}

func swap[T any](i, j *T) {
	*i, *j = *j, *i
}

func buscarPadre(pos int) int {
	return valorAbs((pos - 1) / 2)
}
func valorAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
