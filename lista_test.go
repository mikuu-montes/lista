package lista_test

import (
	Lista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const tamañoVolumen = 10000

// Pre condicion: La lista debe haber sido creada previamente con "CrearListaEnlazada()"
// Post condicion: Hace las pruebas que comprueban que una lista vacia se comporte como tal.
func pruebasListaVacia[T any](t *testing.T, lista Lista.Lista[T]) {
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

// Pre condicion: -
// Post Condicion: Rellena el array de forma ascendente (0, 1, 2, ...)
func rellenarArray(array []int) {
	for i := range array {
		array[i] = i
	}
}

func TestListaVacia(t *testing.T) {
	nuevaLista := Lista.CrearListaEnlazada[int]()
	pruebasListaVacia(t, nuevaLista)
}

func TestUnSoloElemento(t *testing.T) {
	listaBool := Lista.CrearListaEnlazada[bool]()

	pruebasListaVacia(t, listaBool)

	listaBool.InsertarPrimero(true)

	require.False(t, listaBool.EstaVacia())
	require.Equal(t, 1, listaBool.Largo())
	require.Equal(t, true, listaBool.VerPrimero())
	require.Equal(t, true, listaBool.VerUltimo())
	require.Equal(t, true, listaBool.BorrarPrimero())

	pruebasListaVacia(t, listaBool)
}

func TestEsLista(t *testing.T) {
	listaString := Lista.CrearListaEnlazada[string]()

	pruebasListaVacia(t, listaString)

	elementosString := []string{"Que", "Tal", "¿Como", "Se", "Encuentra?"}

	for _, elemento := range elementosString {
		listaString.InsertarUltimo(elemento)
	}

	require.False(t, listaString.EstaVacia())
	require.Equal(t, len(elementosString), listaString.Largo())
	require.Equal(t, "Que", listaString.VerPrimero())
	require.Equal(t, "Encuentra?", listaString.VerUltimo())

	for i := 0; i < len(elementosString); i++ {
		require.Equal(t, len(elementosString)-i, listaString.Largo())
		require.Equal(t, elementosString[i], listaString.VerPrimero())
		require.Equal(t, elementosString[i], listaString.BorrarPrimero())
	}

	pruebasListaVacia(t, listaString)
}

func TestVolumen(t *testing.T) {
	listaInt := Lista.CrearListaEnlazada[int]()

	pruebasListaVacia(t, listaInt)

	arrayInts := make([]int, tamañoVolumen)
	rellenarArray(arrayInts)

	for _, elemento := range arrayInts {
		listaInt.InsertarPrimero(elemento)
	}

	require.False(t, listaInt.EstaVacia())
	require.Equal(t, tamañoVolumen, listaInt.Largo())
	require.Equal(t, arrayInts[tamañoVolumen-1], listaInt.VerPrimero())
	require.Equal(t, arrayInts[0], listaInt.VerUltimo())

	for i := 0; i < len(arrayInts); i++ {
		require.Equal(t, tamañoVolumen-i, listaInt.Largo())
		require.Equal(t, arrayInts[tamañoVolumen-i-1], listaInt.VerPrimero())
		require.Equal(t, arrayInts[tamañoVolumen-i-1], listaInt.BorrarPrimero())
	}

	pruebasListaVacia(t, listaInt)
}

func TestIntercalado(t *testing.T) {
	listaFloat := Lista.CrearListaEnlazada[float64]()

	pruebasListaVacia(t, listaFloat)

	listaFloat.InsertarPrimero(1.0)
	require.Equal(t, 1.0, listaFloat.VerPrimero())
	require.Equal(t, 1.0, listaFloat.VerUltimo())
	require.Equal(t, 1, listaFloat.Largo())

	listaFloat.InsertarPrimero(2.0)
	require.Equal(t, 2.0, listaFloat.VerPrimero())
	require.Equal(t, 1.0, listaFloat.VerUltimo())
	require.Equal(t, 2, listaFloat.Largo())

	require.Equal(t, 2.0, listaFloat.BorrarPrimero())
	require.Equal(t, 1, listaFloat.Largo())
	require.Equal(t, 1.0, listaFloat.VerPrimero())
	require.Equal(t, 1.0, listaFloat.VerUltimo())

	listaFloat.InsertarUltimo(3.0)
	require.Equal(t, 1.0, listaFloat.VerPrimero())
	require.Equal(t, 3.0, listaFloat.VerUltimo())
	require.Equal(t, 2, listaFloat.Largo())

	require.Equal(t, 1.0, listaFloat.BorrarPrimero())
	require.Equal(t, 1, listaFloat.Largo())
	require.Equal(t, 3.0, listaFloat.VerPrimero())
	require.Equal(t, 3.0, listaFloat.VerUltimo())

	require.Equal(t, 3.0, listaFloat.BorrarPrimero())
	require.Equal(t, 0, listaFloat.Largo())

	pruebasListaVacia(t, listaFloat)
}
