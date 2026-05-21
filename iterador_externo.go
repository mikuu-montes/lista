package lista

type iteradorLista[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodo[T]
	anterior *nodo[T]
}

// Pre condicion: -
// Post condicion: Retorna un iterador asociado a la lista.
// Queda posicionado en el primer elemento de la misma o en estado "final" si esta vacia.
// El iterador no modifica a la lista.
func crearIteradorExterno[T any](lista *listaEnlazada[T]) *iteradorLista[T] {
	return &iteradorLista[T]{lista: lista, actual: lista.primero}
}

// Post condicion: Retorna true si el iterador tiene mas iteraciones para hacer, false en caso contrario.
func (i *iteradorLista[T]) HayAlgoMas() bool {
	return i.actual != nil
}

// Post condicion: Devuelve el elemento de la posicion actual del iterador(sin modificarlo).
// Si no hay elementos, entra en pánico con un mensaje "El iterador termino de iterar".
func (i *iteradorLista[T]) VerActual() T {
	if !i.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}

	return i.actual.dato
}

// Post condicion: Si todavia hay iteraciones para hacer, se avanza una iteración.
// En caso contrario entra en pánico con un mensaje "El iterador termino de iterar".
func (i *iteradorLista[T]) Avanzar() {
	if !i.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	i.anterior = i.actual
	i.actual = i.actual.siguiente
}

// Post condicion: Inserta un elemento a la lista, en la posicion de la actual iteración. Dejando el iterador en la posicion del nuevo elemento.
func (i *iteradorLista[T]) Insertar(dato T) {
	nuevoNodo := nodoCrear(dato, i.actual)

	if i.anterior == nil {
		i.lista.primero = nuevoNodo
	} else {
		i.anterior.siguiente = nuevoNodo
	}

	if i.actual == nil {
		i.lista.ultimo = nuevoNodo
	}

	i.actual = nuevoNodo
	i.lista.largo++
}

// Post condicion: Si el iterador tiene elementos, elimina el de la iteración actual, lo retorna, y avanza al siguiente elemento.
// En caso contrario entra en pánico con un mensaje "El iterador termino de iterar".
func (i *iteradorLista[T]) Borrar() T {
	if !i.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}

	datoEliminado := i.actual.dato

	if i.anterior == nil {
		i.lista.primero = i.actual.siguiente
	} else {
		i.anterior.siguiente = i.actual.siguiente
	}

	if i.actual.siguiente == nil {
		i.lista.ultimo = i.anterior
	}

	i.actual = i.actual.siguiente
	i.lista.largo--

	return datoEliminado
}
