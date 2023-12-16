package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iteradorExterno[T any] struct {
	anterior *nodoLista[T]
	actual   *nodoLista[T]
	lista    *listaEnlazada[T]
}

func crearNodo[T any](dato T) *nodoLista[T] {
	nodoNuevo := new(nodoLista[T])
	nodoNuevo.dato = dato
	return nodoNuevo
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	lista.primero = nil
	lista.ultimo = nil
	lista.largo = 0
	return lista
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nodoNuevo := crearNodo(dato)
	if lista.EstaVacia() {
		lista.ultimo = nodoNuevo
	} else {
		nodoNuevo.siguiente = lista.primero
	}
	lista.primero = nodoNuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodoNuevo := crearNodo(dato)
	if lista.EstaVacia() {
		lista.primero = nodoNuevo
	} else {
		lista.ultimo.siguiente = nodoNuevo
	}
	lista.ultimo = nodoNuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	datoCopia := lista.VerPrimero()
	lista.primero = lista.primero.siguiente
	lista.largo--
	if lista.EstaVacia() {
		lista.ultimo = nil
	}
	return datoCopia
}

func (lista listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual, res := lista.primero, true
	for actual != nil && res {
		res = visitar(actual.dato)
		actual = actual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iter := new(iteradorExterno[T])
	iter.actual = lista.primero
	iter.anterior = nil
	iter.lista = lista
	return iter
}

func (iter iteradorExterno[T]) VerActual() T {
	if iter.HaySiguiente() {
		return iter.actual.dato
	}
	panic("El iterador termino de iterar")
}

func (iter iteradorExterno[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iteradorExterno[T]) Siguiente() T {
	if iter.HaySiguiente() {
		datoCopia := iter.actual.dato
		iter.anterior = iter.actual
		iter.actual = iter.actual.siguiente
		return datoCopia
	}
	panic("El iterador termino de iterar")

}

func (iter *iteradorExterno[T]) Insertar(dato T) {
	nodoNuevo := crearNodo(dato)
	if iter.lista.EstaVacia() {
		iter.lista.primero = nodoNuevo
		iter.lista.ultimo = nodoNuevo
	} else if iter.anterior == nil { // Esto es si el iterador está en la primera posicion
		nodoNuevo.siguiente = iter.lista.primero
		iter.lista.primero = nodoNuevo
	} else if !iter.HaySiguiente() { // Esto es si el iterador está en la última posicion
		iter.lista.ultimo.siguiente = nodoNuevo
		iter.lista.ultimo = nodoNuevo
	} else {
		nodoNuevo.siguiente = iter.actual
		iter.anterior.siguiente = nodoNuevo
	}
	iter.actual = nodoNuevo
	iter.lista.largo++
}

func (iter *iteradorExterno[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	datoCopia := iter.actual.dato
	if iter.anterior == nil {
		iter.lista.primero = iter.actual.siguiente
	} else if iter.actual.siguiente == nil {
		iter.lista.ultimo = iter.anterior
	} else {
		iter.anterior.siguiente = iter.actual.siguiente
	}
	iter.actual = iter.actual.siguiente
	iter.lista.largo--
	return datoCopia
}
