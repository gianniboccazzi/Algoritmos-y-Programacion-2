package lista

type IteradorLista[T any] interface {

	// VerActual devuelve el dato actual.
	VerActual() T

	// HaySiguiente devuelve verdadero si hay algún elemento siguiente, false de lo contrario.
	HaySiguiente() bool

	// Siguiente devuelve el dato del siguiente.
	Siguiente() T

	// Insertar agrega un nuevo elemento antes del actual.
	Insertar(T)

	// Borrar saca el elemento que está en ese momento
	Borrar() T
}

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos insertados, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento a la lista, al principio de la misma.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento a la lista, al final de la misma.
	InsertarUltimo(T)

	// BorrarPrimero saca el primer elemento de la lista. Si la lista tiene elementos, se quita el primero de la misma,
	// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primero de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del primero de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerUltimo() T

	// Largo obtiene el valor del largo de la lista.
	Largo() int

	// Iterar
	Iterar(visitar func(T) bool)

	// Iterador
	Iterador() IteradorLista[T]
}
