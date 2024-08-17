package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())

	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })

}

func TestInsertarSinIterar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())
	require.False(t, lista.EstaVacia())

	lista.InsertarPrimero(2)
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())
	require.False(t, lista.EstaVacia())

	lista.InsertarUltimo(3)
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
	require.False(t, lista.EstaVacia())
}

func TestBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarUltimo(3)

	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.EqualValues(t, 3, lista.BorrarPrimero())

	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

func TestVaciarLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarUltimo(3)

	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())

	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
	require.False(t, lista.EstaVacia())

	require.EqualValues(t, 3, lista.VerUltimo())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())

	require.EqualValues(t, 3, lista.VerUltimo())
	require.EqualValues(t, 3, lista.VerPrimero())
	require.EqualValues(t, 3, lista.BorrarPrimero())

	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })

	require.True(t, lista.EstaVacia())

}

func TestPanicsLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[bool]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

func TestPanicsIteradorExt(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[bool]()

	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

}

func TestIterExtInsertarPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	iter := lista.Iterador()

	iter.Insertar(0)

	require.EqualValues(t, 0, iter.VerActual())
	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, 2, lista.VerUltimo())
}

func TestIterExtInsertarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	lista.InsertarPrimero("Hola")
	lista.InsertarUltimo("Mundo")

	iter := lista.Iterador()
	iter.Siguiente()
	iter.Insertar(",")

	require.EqualValues(t, "Hola", lista.BorrarPrimero())
	require.EqualValues(t, ",", lista.BorrarPrimero())
	require.EqualValues(t, "Mundo", lista.BorrarPrimero())
}

func TestIterExtInsertarFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[bool]()

	lista.InsertarPrimero(false)
	lista.InsertarPrimero(false)

	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()

	iter.Insertar(true)

	require.EqualValues(t, true, iter.VerActual())
	require.EqualValues(t, true, lista.VerUltimo())
	require.EqualValues(t, false, lista.VerPrimero())

}

func TestInsertarIterExt(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	iter := lista.Iterador()

	require.False(t, iter.HaySiguiente())
	// Insertar en el primer lugar
	iter.Insertar("Hola")
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, "Hola", iter.VerActual())
	require.EqualValues(t, "Hola", lista.VerPrimero())
	require.EqualValues(t, "Hola", lista.VerUltimo())

	iter.Siguiente()
	// Insertar al final
	iter.Insertar("mundo")
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, "mundo", iter.VerActual())
	require.EqualValues(t, "Hola", lista.VerPrimero())
	require.EqualValues(t, "mundo", lista.VerUltimo())

	// Insertar en el medio
	iter.Insertar(",")
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, ",", iter.VerActual())
	require.EqualValues(t, "Hola", lista.VerPrimero())
	require.EqualValues(t, "mundo", lista.VerUltimo())

	require.EqualValues(t, "Hola", lista.BorrarPrimero())
	require.EqualValues(t, ",", lista.BorrarPrimero())
	require.EqualValues(t, "mundo", lista.BorrarPrimero())

}

func TestBorrarIterExt(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	require.EqualValues(t, 1, iter.VerActual())
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, 1, iter.Borrar())
	require.EqualValues(t, 2, iter.VerActual())
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, 2, iter.Borrar())
	require.EqualValues(t, 3, iter.VerActual())
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, 3, iter.Borrar())

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}

func TestIterExtBorrarPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()

	require.EqualValues(t, 1, iter.Borrar())

	require.EqualValues(t, 2, iter.VerActual())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
}

func TestIterExtBorrarFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	iter := lista.Iterador()
	iter.Siguiente()

	require.EqualValues(t, 2, iter.Borrar())

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())
}

func TestPruebaVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for valor := 0; valor >= 10000; valor++ {
		lista.InsertarUltimo(valor)
		require.EqualValues(t, valor, lista.VerUltimo())
		require.EqualValues(t, 0, lista.VerPrimero())

	}
	for valor := 0; valor >= 10000; valor++ {
		require.EqualValues(t, valor, lista.VerPrimero())
		require.EqualValues(t, 10000, lista.VerUltimo())
		require.EqualValues(t, valor, lista.BorrarPrimero())

	}
	require.True(t, lista.EstaVacia())
}

func TestSumaElem(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)
	lista.InsertarUltimo(7)
	lista.InsertarUltimo(3)

	suma := 0
	lista.Iterar(func(v int) bool {
		suma += v
		return true
	})

	require.EqualValues(t, 11, suma)
}

func TestCondicionCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("a")
	lista.InsertarUltimo("b")
	lista.InsertarUltimo("c")
	lista.InsertarUltimo("d")
	lista.InsertarUltimo("e")
	lista.InsertarUltimo("f")

	contador := 0
	lista.Iterar(func(v string) bool {
		contador += 1
		return contador < 5
	})
}
