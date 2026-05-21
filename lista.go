package lista

type Lista[T any] interface {

	// Post condicion: Devuelve true si la lista no tiene elementos o false en caso contrario.
	EstaVacia() bool

	// Post condicion: Se inserta un nuevo elemento al comienzo de la lista.
	InsertarPrimero(T)

	// Post condicion: Se inserta un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	// Post condicion: Si la lista tiene elementos, se elimina el primero de la misma y retorna ese valor.
	//				   Si esta vacía entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// Post condicion: Si la lista tiene elementos, retorna el valor del primero.
	//				   Si esta vacía entra en pánico con un mensaje "La lista esta vacia"
	VerPrimero() T

	// Post condicion: Si la lista tiene elementos, retorna el valor del último.
	//				   Si esta vacía entra en pánico con un mensaje "La lista esta vacia"
	VerUltimo() T

	// Post condicion: Devuelve la cantidad de elementos de la lista.
	Largo() int

	// Post condicion: Recorre la lista (de primer a último elemento), aplicando en cada uno la funcion "visitar"
	//        	       Si la función "visitar" devuelve false, o si ya se recorrieron todos los elementos, la iteración se corta.
	// 				   La lista no se modifica.
	Iterar(visitar func(T) bool)

	// Post condicion: Retorna un iterador externo para la lista.
	// 				   Posicionandose en el primer elemento de la misma, o en estado "final" si esta vacía.
	//                 Su creación no modifica a la lista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	//	Post condicion: Devuelve el elemento de la posicion actual del iterador(sin modificarlo).
	// 					Si no hay elementos, entra en pánico con un mensaje "El iterador termino de iterar".
	VerActual() T

	//	Post condicion: Retorna true si el iterador tiene mas iteraciones para hacer, false en caso contrario.
	HayAlgoMas() bool

	//	Post condicion: Si todavia hay iteraciones para hacer, se avanza una iteración.
	// 					En caso contrario entra en pánico con un mensaje "El iterador termino de iterar".
	Avanzar()

	//	Post condicion: Inserta un elemento a la lista, en la posicion de la actual iteración. Dejando el iterador en la posicion del nuevo elemento.
	Insertar(T)

	//	Post condicion: Si el iterador tiene elementos, elimina el de la iteración actual, lo retorna, y avanza al siguiente
	// 					En caso contrario entra en pánico con un mensaje "El iterador termino de iterar".
	Borrar() T
}
