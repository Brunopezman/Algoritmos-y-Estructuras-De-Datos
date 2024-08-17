package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	require.True(t, cola.EstaVacia())

	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })

	cola.Encolar(1)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 1, cola.VerPrimero())
	require.EqualValues(t, 1, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestInvarianteFIFO(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	cola.Encolar(4)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 1, cola.Desencolar())
	require.EqualValues(t, 2, cola.Desencolar())
	require.EqualValues(t, 3, cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 4, cola.Desencolar())

	cola.Encolar(5)
	cola.Encolar(6)

	require.EqualValues(t, 5, cola.Desencolar())
	require.EqualValues(t, 6, cola.Desencolar())

	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.True(t, cola.EstaVacia())
}

func TestEstaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()

	require.True(t, cola.EstaVacia())
	cola.Encolar("Juan Carlos")
	require.False(t, cola.EstaVacia())
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
}

func TestPanics(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[bool]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestStr(t *testing.T) {
	colaStr := TDACola.CrearColaEnlazada[string]()
	colaStr.Encolar("Juan Roman Riquelme")
	require.Equal(t, "Juan Roman Riquelme", colaStr.Desencolar(), "El primero debe ser 'Juan Roman Riquelme'")
	require.True(t, colaStr.EstaVacia())
}

func TestTipoCola(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	colaDeColas := TDACola.CrearColaEnlazada[TDACola.Cola[string]]()
	colaDeColas.Encolar(cola)
	require.Equal(t, cola, colaDeColas.Desencolar(), "El primero debe ser la cola")
	require.True(t, colaDeColas.EstaVacia())
}

func TestInt(t *testing.T) {
	colaInt := TDACola.CrearColaEnlazada[int]()
	colaInt.Encolar(5)
	require.Equal(t, 5, colaInt.Desencolar(), "El primero debe ser 5")
	require.True(t, colaInt.EstaVacia())
}

func TestBool(t *testing.T) {
	colaBool := TDACola.CrearColaEnlazada[bool]()
	colaBool.Encolar(true)
	require.Equal(t, true, colaBool.Desencolar(), "El primero debe ser 'true'")
	require.True(t, colaBool.EstaVacia())
}

func TestColaVaciada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	cola.Encolar(10)
	cola.Encolar(20)
	cola.Encolar(30)
	require.False(t, cola.EstaVacia())
	cola.Desencolar()
	cola.Desencolar()
	cola.Desencolar()

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	cola.Encolar(1)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 1, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestVolumen(t *testing.T) {
	colaVolumen := TDACola.CrearColaEnlazada[int]()

	for valor := 0; valor < 10000; valor++ {
		colaVolumen.Encolar(valor)
		require.False(t, colaVolumen.EstaVacia())
		require.EqualValues(t, 0, colaVolumen.VerPrimero())

	}
	for valor := 0; valor < 10000; valor++ {
		require.EqualValues(t, valor, colaVolumen.VerPrimero())
		e := colaVolumen.Desencolar()
		require.EqualValues(t, valor, e)

	}
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaVolumen.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaVolumen.VerPrimero() })
}
