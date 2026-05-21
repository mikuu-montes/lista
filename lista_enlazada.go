package lista

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

// Pre condicion: -
// Post condicion: Retorna un puntero a un nodo con el dato enviado y apunta como "siguiente nodo" a "siguiente"
func nodoCrear[T any](dato T, siguiente *nodo[T]) *nodo[T] {
	return &nodo[T]{dato: dato, siguiente: siguiente}
}

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

// Pre condicion: -
// Post condicion: Retorna una lista enlazada.
func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

// Post condicion: Devuelve true si la lista no tiene elementos o false en caso contrario.
func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

// Post condicion: Se inserta un nuevo elemento al comienzo de la lista.
func (l *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoNodo := nodoCrear(elemento, l.primero)

	if l.EstaVacia() {
		l.ultimo = nuevoNodo
	}

	l.primero = nuevoNodo
	l.largo++
}

// Post condicion: Se inserta un nuevo elemento al final de la lista.
func (l *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevoNodo := nodoCrear(elemento, nil)

	if l.EstaVacia() {
		l.primero = nuevoNodo
	} else {
		l.ultimo.siguiente = nuevoNodo
	}

	l.ultimo = nuevoNodo
	l.largo++
}

// Post condicion: Si la lista tiene elementos, se elimina el primero de la misma y retorna ese valor.
// Si esta vacía entra en pánico con un mensaje "La lista esta vacia".
func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}

	datoBorrado := l.primero.dato
	l.primero = l.primero.siguiente

	if l.largo == 1 {
		l.ultimo = nil
	}

	l.largo--
	return datoBorrado
}

// Post condicion: Si la lista tiene elementos, retorna el valor del primero.
// Si esta vacía entra en pánico con un mensaje "La lista esta vacia"
func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}

	return l.primero.dato
}

// Post condicion: Si la lista tiene elementos, retorna el valor del último.
// Si esta vacía entra en pánico con un mensaje "La lista esta vacia"
func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}

	return l.ultimo.dato
}

// Post condicion: Devuelve la cantidad de elementos de la lista.
func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

// Post condicion: Recorre la lista (de primer a último elemento), aplicando en cada uno la funcion "visitar"
// Si la función "visitar" devuelve false, o si ya se recorrieron todos los elementos, la iteración se corta.
// La lista no se modifica.
func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	nodoActual := l.primero
	for nodoActual != nil && visitar(nodoActual.dato) {
		nodoActual = nodoActual.siguiente
	}
}

// Post condicion: Retorna un iterador externo para la lista.
// Posicionandose en el primer elemento de la misma, o en estado "final" si esta vacía.
func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return crearIteradorExterno(l)
}
