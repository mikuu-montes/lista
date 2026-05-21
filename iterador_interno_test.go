package lista_test

import (
	Lista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVolumenIterador(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()

	arrayInts := make([]int, tamañoVolumen)
	rellenarArray(arrayInts)

	for _, elemento := range arrayInts {
		lista.InsertarUltimo(elemento)
	}

	contador := 0

	lista.Iterar(func(info int) bool {
		contador++
		return true
	})

	require.Equal(t, len(arrayInts), contador)
}

func TestIterarSinCorte(t *testing.T) {
	lista := Lista.CrearListaEnlazada[int]()

	arrayInts := []int{1, 2, 3, 4, 5}

	for _, elemento := range arrayInts {
		lista.InsertarUltimo(elemento)
	}

	suma := 0

	lista.Iterar(func(dato int) bool {
		suma += dato
		return true
	})

	require.Equal(t, 15, suma)
}

func TestIterarConCorte(t *testing.T) {
	lista := Lista.CrearListaEnlazada[string]()

	arrayString := []string{"A", "E", "I", "O", "U"}

	for _, elemento := range arrayString {
		lista.InsertarUltimo(elemento)
	}

	contador := 0

	lista.Iterar(func(dato string) bool {
		contador++
		return dato != "I"
	})

	require.Equal(t, 3, contador)
}
