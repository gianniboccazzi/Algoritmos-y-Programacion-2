package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func crearNodo[T any](dato T) *nodoCola[T] {
	nodoNuevo := new(nodoCola[T])
	nodoNuevo.dato = dato
	return nodoNuevo
}

func CrearColaEnlazada[T any]() Cola[T] {
	return new(colaEnlazada[T])
}

func (cola colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(dato T) {
	nodoNuevo := crearNodo(dato)
	if cola.EstaVacia() {
		cola.primero = nodoNuevo
	} else {
		cola.ultimo.prox = nodoNuevo
	}
	cola.ultimo = nodoNuevo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := cola.VerPrimero()
	cola.primero = cola.primero.prox
	if cola.EstaVacia() {
		cola.ultimo = nil
	}
	return dato
}
