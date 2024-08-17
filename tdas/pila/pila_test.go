package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pilaInt := TDAPila.CrearPilaDinamica[int]()

	require.True(t, pilaInt.EstaVacia())

	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaInt.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaInt.VerTope() })

	pilaInt.Apilar(1)
	require.False(t, pilaInt.EstaVacia())
	pilaInt.Desapilar()
	require.True(t, pilaInt.EstaVacia())
}

func TestInvariante(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(4)
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 4, pila.VerTope())
	pila.Desapilar()
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 3, pila.VerTope())
	pila.Desapilar()
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 2, pila.VerTope())
	pila.Desapilar()
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 1, pila.VerTope())
	pila.Desapilar()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestUint(t *testing.T) {
	pilaUint := TDAPila.CrearPilaDinamica[uint]()

	pilaUint.Apilar(1)
	require.False(t, pilaUint.EstaVacia())
	pilaUint.Apilar(2)
	pilaUint.Apilar(3)
	pilaUint.Apilar(4)
	require.EqualValues(t, 4, pilaUint.VerTope())
	pilaUint.Desapilar()
	require.EqualValues(t, 3, pilaUint.VerTope())
	pilaUint.Desapilar()
	require.EqualValues(t, 2, pilaUint.VerTope())
	pilaUint.Desapilar()
	require.EqualValues(t, 1, pilaUint.VerTope())
	pilaUint.Desapilar()
	require.True(t, pilaUint.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaUint.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaUint.VerTope() })

}

func TestStr(t *testing.T) {
	pilaStr := TDAPila.CrearPilaDinamica[string]()

	pilaStr.Apilar("a")
	pilaStr.Apilar("b")
	pilaStr.Apilar("c")
	pilaStr.Apilar("d")
	require.EqualValues(t, "d", pilaStr.VerTope())
	pilaStr.Desapilar()
	require.False(t, pilaStr.EstaVacia())
	require.EqualValues(t, "c", pilaStr.VerTope())
	pilaStr.Desapilar()
	require.False(t, pilaStr.EstaVacia())
	require.EqualValues(t, "b", pilaStr.VerTope())
	pilaStr.Desapilar()
	require.False(t, pilaStr.EstaVacia())
	require.EqualValues(t, "a", pilaStr.VerTope())
	pilaStr.Desapilar()
	require.True(t, pilaStr.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaStr.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaStr.VerTope() })
}

func TestCondicionesBorde(t *testing.T) {
	pilaInt := TDAPila.CrearPilaDinamica[int]()

	pilaInt.Apilar(10)
	pilaInt.Apilar(20)
	pilaInt.Apilar(30)
	pilaInt.Apilar(40)
	pilaInt.Apilar(50)
	pilaInt.Apilar(60)
	pilaInt.Desapilar()
	pilaInt.Desapilar()
	pilaInt.Desapilar()
	pilaInt.Desapilar()
	pilaInt.Desapilar()
	pilaInt.Desapilar()

	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaInt.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaInt.VerTope() })

	pilaInt.Apilar(1)
	require.False(t, pilaInt.EstaVacia())
	pilaInt.Desapilar()
	require.True(t, pilaInt.EstaVacia())
}

func TestVolumen(t *testing.T) {
	pilaVolumen := TDAPila.CrearPilaDinamica[int]()

	for valor := 0; valor < 10000; valor++ {
		pilaVolumen.Apilar(valor)
		require.False(t, pilaVolumen.EstaVacia())
		require.EqualValues(t, valor, pilaVolumen.VerTope())

	}

	for valor := 9999; valor >= 0; valor-- {
		require.EqualValues(t, valor, pilaVolumen.VerTope())
		e := pilaVolumen.Desapilar()
		require.EqualValues(t, valor, e)

	}
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaVolumen.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaVolumen.VerTope() })
}
