package pila

const _CAPACIDAD_INICIAL int = 5
const _COEFICIENTE_REDIMENSION int = 2
const _COEFICIENTE_DISMINUCION int = 4

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (pila *pilaDinamica[T]) redimensionar(capacidad int) {
	datosCopia := make([]T, capacidad)
	copy(datosCopia, pila.datos)
	pila.datos = datosCopia
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, _CAPACIDAD_INICIAL)
	pila.cantidad = 0
	return pila
}

func (pila pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elem T) {
	if pila.cantidad == len(pila.datos) {
		pila.redimensionar(len(pila.datos) * _COEFICIENTE_REDIMENSION)
	}
	pila.datos[pila.cantidad] = elem
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	if pila.cantidad*_COEFICIENTE_DISMINUCION <= len(pila.datos) {
		capacidad := len(pila.datos) / _COEFICIENTE_REDIMENSION
		pila.redimensionar(capacidad)
	}
	elemDesapilado := pila.VerTope()
	pila.cantidad--
	return elemDesapilado
}
