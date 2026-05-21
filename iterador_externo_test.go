package lista_test

import (
	Lista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

// Pre condicion: La lista debe haber sido creada previamente con "CrearListaEnlazada()"
// Post condicion: Chequea si un iterador cumple ciertas cosas cuando la lista esta recien creada/vacia.
func iteradorRecienCreado[T any](t *testing.T, lista Lista.Lista[T]) {
	iterador := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.False(t, iterador.HayAlgoMas())
}
func TestRecienCreado(t *testing.T) {
	nuevaLista := Lista.CrearListaEnlazada[bool]()
	iteradorRecienCreado(t, nuevaLista)
}

func TestInsertarElementos(t *testing.T) {
	listaFloat := Lista.CrearListaEnlazada[float64]()

	iteradorRecienCreado(t, listaFloat)

	arrayFloats := []float64{2.0, 4.0}
	for _, elemento := range arrayFloats {
		listaFloat.InsertarUltimo(elemento)
	}

	iterador := listaFloat.Iterador()

	require.Equal(t, 2.0, iterador.VerActual())
	require.True(t, iterador.HayAlgoMas())

	iterador.Insertar(1.0)
	require.Equal(t, 1.0, listaFloat.VerPrimero())

	iterador.Avanzar()
	iterador.Avanzar()

	iterador.Insertar(3.0)
	require.Equal(t, 3.0, iterador.VerActual())

	iterador.Avanzar()
	iterador.Avanzar()

	iterador.Insertar(5.0)
	require.Equal(t, 5.0, listaFloat.VerUltimo())
	require.Equal(t, 5, listaFloat.Largo())

	iterador.Avanzar()

	require.False(t, iterador.HayAlgoMas())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Avanzar() })
}

func TestVolumenIteradorExterno(t *testing.T) {
	listaInt := Lista.CrearListaEnlazada[int]()

	iteradorRecienCreado(t, listaInt)

	arrayInts := make([]int, tamañoVolumen)
	rellenarArray(arrayInts)

	iterador := listaInt.Iterador()

	for _, elemento := range arrayInts {
		iterador.Insertar(elemento)
	}

	for i := len(arrayInts) - 1; i >= 0; i-- {
		require.Equal(t, arrayInts[i], iterador.VerActual())
		require.Equal(t, arrayInts[i], listaInt.VerPrimero())
		iterador.Borrar()
	}

	require.False(t, iterador.HayAlgoMas())
	require.True(t, listaInt.EstaVacia())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Avanzar() })
}

func TestBorrarElementos(t *testing.T) {
	listaString := Lista.CrearListaEnlazada[string]()

	iteradorRecienCreado(t, listaString)

	arrayString := []string{"A", "E", "I", "O", "U"}

	for _, elemento := range arrayString {
		listaString.InsertarUltimo(elemento)
	}

	iterador := listaString.Iterador()

	iterador.Borrar()
	require.Equal(t, "E", listaString.VerPrimero())
	require.Equal(t, "E", iterador.VerActual())

	iterador.Avanzar()

	iterador.Borrar()
	require.Equal(t, "O", iterador.VerActual())
	encontrado := false
	listaString.Iterar(func(dato string) bool {
		if dato == "I" {
			encontrado = true
		}
		return true
	})
	require.False(t, encontrado)

	iterador.Avanzar()

	iterador.Borrar()
	require.Equal(t, "O", listaString.VerUltimo())

	require.False(t, iterador.HayAlgoMas())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Avanzar() })
}
