package pila

/* Definición del struct pila proporcionado por la cátedra. */
const LONG_INICIAL = 10
const TAMAÑO_DUPLICADO = 2
const MITAD_TAMAÑO = 2
const TAMAÑO_CUADRUPLICADO = 4
const TAMAÑO_MINIMO = 0

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {

	pila := new(pilaDinamica[T])
	pila.datos = make([]T, LONG_INICIAL)
	pila.cantidad = TAMAÑO_MINIMO
	return pila
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if pila.cantidad == cap(pila.datos) {
		pila.datos = redimensionarespacio(pila.datos, cap(pila.datos)*TAMAÑO_DUPLICADO)
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++

}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	if pila.cantidad*TAMAÑO_CUADRUPLICADO <= cap(pila.datos) && cap(pila.datos) > TAMAÑO_MINIMO {
		pila.datos = redimensionarespacio(pila.datos, cap(pila.datos)/MITAD_TAMAÑO)
	}
	tope := pila.VerTope()
	pila.cantidad--
	return tope
}

func redimensionarespacio[T any](datos []T, capacidad int) []T {
	redimension := make([]T, capacidad)
	copy(redimension, datos)
	return redimension
}
