package cola_prioridad_test

import (
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func cmp(a, b int) int {
	return a - b
}

func func_cmp_str(s1, s2 string) int {
	if s1 == s2 {
		return 0
	} else if s1 > s2 {
		return 1
	}
	return -1
}
func TestHeapVacio(t *testing.T) {
	cp := TDAHeap.CrearHeap(cmp)
	require.True(t, cp.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.Desencolar() })
	require.EqualValues(t, 0, cp.Cantidad())
}

func TestUnElemento(t *testing.T) {
	cp := TDAHeap.CrearHeap(cmp)
	require.True(t, cp.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.Desencolar() })
	require.EqualValues(t, 0, cp.Cantidad())

	cp.Encolar(7)
	require.False(t, cp.EstaVacia())
	require.EqualValues(t, 7, cp.VerMax())
	require.EqualValues(t, 1, cp.Cantidad())
}

func TestVariosElementos(t *testing.T) {
	cp := TDAHeap.CrearHeap(cmp)
	require.True(t, cp.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.Desencolar() })
	require.EqualValues(t, 0, cp.Cantidad())

	cp.Encolar(7)
	require.False(t, cp.EstaVacia())
	require.EqualValues(t, 7, cp.VerMax())
	require.EqualValues(t, 1, cp.Cantidad())

	cp.Encolar(10)
	require.False(t, cp.EstaVacia())
	require.EqualValues(t, 10, cp.VerMax())
	require.EqualValues(t, 2, cp.Cantidad())

	cp.Encolar(5)
	require.False(t, cp.EstaVacia())
	require.EqualValues(t, 10, cp.VerMax())
	require.EqualValues(t, 3, cp.Cantidad())
}

func TestDesencolar(t *testing.T) {
	cp := TDAHeap.CrearHeap(cmp)
	require.True(t, cp.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.Desencolar() })
	require.EqualValues(t, 0, cp.Cantidad())

	cp.Encolar(10)
	cp.Encolar(8)
	cp.Encolar(20)
	cp.Encolar(11)
	cp.Encolar(15)
	require.False(t, cp.EstaVacia())
	require.EqualValues(t, 20, cp.VerMax())
	require.EqualValues(t, 5, cp.Cantidad())

	cp.Desencolar()
	require.False(t, cp.EstaVacia())
	require.EqualValues(t, 15, cp.VerMax())
	require.EqualValues(t, 4, cp.Cantidad())

	cp.Desencolar()
	require.False(t, cp.EstaVacia())
	require.EqualValues(t, 11, cp.VerMax())
	require.EqualValues(t, 3, cp.Cantidad())

	cp.Desencolar()
	require.False(t, cp.EstaVacia())
	require.EqualValues(t, 10, cp.VerMax())
	require.EqualValues(t, 2, cp.Cantidad())

	cp.Desencolar()
	require.False(t, cp.EstaVacia())
	require.EqualValues(t, 8, cp.VerMax())
	require.EqualValues(t, 1, cp.Cantidad())

	cp.Desencolar()
	require.True(t, cp.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.Desencolar() })
	require.EqualValues(t, 0, cp.Cantidad())
}

func TestCrearHeapArreglo(t *testing.T) {
	arr := []int{12, 6, 9, 8}
	cp := TDAHeap.CrearHeapArr(arr, cmp)

	require.False(t, cp.EstaVacia())
	require.EqualValues(t, 12, cp.VerMax())
	require.EqualValues(t, 4, cp.Cantidad())

	cp.Desencolar()
	require.False(t, cp.EstaVacia())
	require.EqualValues(t, 9, cp.VerMax())
	require.EqualValues(t, 3, cp.Cantidad())

	cp.Desencolar()
	require.False(t, cp.EstaVacia())
	require.EqualValues(t, 8, cp.VerMax())
	require.EqualValues(t, 2, cp.Cantidad())

	cp.Encolar(13)
	require.False(t, cp.EstaVacia())
	require.EqualValues(t, 13, cp.VerMax())
	require.EqualValues(t, 3, cp.Cantidad())

	cp.Desencolar()
	require.False(t, cp.EstaVacia())
	require.EqualValues(t, 8, cp.VerMax())
	require.EqualValues(t, 2, cp.Cantidad())

	cp.Desencolar()
	require.False(t, cp.EstaVacia())
	require.EqualValues(t, 6, cp.VerMax())
	require.EqualValues(t, 1, cp.Cantidad())

	cp.Desencolar()
	require.True(t, cp.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.Desencolar() })
	require.EqualValues(t, 0, cp.Cantidad())
}

func TestCrearHeapArregloVacio(t *testing.T) {
	arr := []int{}
	cp := TDAHeap.CrearHeapArr(arr, cmp)
	require.True(t, cp.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.Desencolar() })
	require.EqualValues(t, 0, cp.Cantidad())
}

func TestHeapSort(t *testing.T) {
	arr := []string{"d", "b", "g", "j", "d", "c", "h", "e"}
	arrEsperado := []string{"b", "c", "d", "d", "e", "g", "h", "j"}
	TDAHeap.HeapSort(arr, func_cmp_str)
	require.EqualValues(t, arrEsperado, arr)
}

func TestVolumen(t *testing.T) {

	heap := TDAHeap.CrearHeap(cmp)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())

	for i := 0; i < 10000; i++ {
		require.EqualValues(t, i, heap.Cantidad())
		heap.Encolar(i)
		require.False(t, heap.EstaVacia())
		require.EqualValues(t, i, heap.VerMax())
		require.EqualValues(t, i+1, heap.Cantidad())
	}

	for i := 10000; i > 0; i-- {
		require.EqualValues(t, i, heap.Cantidad())
		require.EqualValues(t, i-1, heap.VerMax())
		require.False(t, heap.EstaVacia())
		heap.Desencolar()
	}

	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestPanic(t *testing.T) {
	cp := TDAHeap.CrearHeap(cmp)
	require.True(t, cp.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.Desencolar() })

	cp.Encolar(5)
	require.False(t, cp.EstaVacia())

	cp.Desencolar()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cp.Desencolar() })
}
