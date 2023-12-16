package cola_prioridad

type heap[T comparable] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

const _CAPACIDAD_INICIAL int = 500000
const _COEFICIENTE_REDIMENSION int = 2
const _COEFICIENTE_DISMINUCION int = 4

func swap[T comparable](a, b *T) {
	*a, *b = *b, *a
}

func esHoja(posicion int, cantidad int) bool {
	return posicion >= cantidad
}

func upheap[T comparable](arreglo []T, hijo int, funcion_cmp func(T, T) int) {
	if arreglo[hijo] == arreglo[0] {
		return
	}
	padre := (hijo - 1) / 2
	if funcion_cmp(arreglo[hijo], arreglo[padre]) > 0 {
		swap(&arreglo[hijo], &arreglo[padre])
		upheap(arreglo, padre, funcion_cmp)
	}
}

func downheap[T comparable](arreglo []T, padre int, funcion_cmp func(T, T) int, cantidad int) {
	if esHoja(padre, cantidad) {
		return
	}
	hijoIzq := 2*padre + 1
	hijoDer := 2*padre + 2
	max := padre
	if hijoIzq < cantidad && funcion_cmp(arreglo[hijoIzq], arreglo[max]) > 0 {
		max = hijoIzq
	}
	if hijoDer < cantidad && funcion_cmp(arreglo[hijoDer], arreglo[max]) > 0 {
		max = hijoDer
	}
	if max != padre {
		swap(&arreglo[padre], &arreglo[max])
		downheap(arreglo, max, funcion_cmp, cantidad)
	}
}

func heapify[T comparable](arreglo []T, padre int, funcion_cmp func(T, T) int, cantidad int) {
	for padre >= 0 {
		downheap(arreglo, padre, funcion_cmp, cantidad)
		padre--
	}
}

func (heap *heap[T]) redimensionar(capacidad int) {
	datos := make([]T, capacidad)
	copy(datos, heap.datos)
	heap.datos = datos
}

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.datos = make([]T, _CAPACIDAD_INICIAL)
	heap.cmp = funcion_cmp
	return heap
}

func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.cmp = funcion_cmp
	if len(arreglo) > _CAPACIDAD_INICIAL {
		heap.datos = make([]T, len(arreglo))
	} else {
		heap.datos = make([]T, _CAPACIDAD_INICIAL)
	}
	heap.cantidad = len(arreglo)
	copy(heap.datos, arreglo)
	heapify(heap.datos, (heap.cantidad-1)/2, funcion_cmp, heap.cantidad)
	return heap
}

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, len(elementos)/2-1, funcion_cmp, len(elementos))
	for i := len(elementos) - 1; i > 0; i-- {
		swap(&elementos[0], &elementos[i])
		downheap(elementos[:i], 0, funcion_cmp, i)
	}
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(elemento T) {
	if heap.cantidad == len(heap.datos) {
		heap.redimensionar(len(heap.datos) * _COEFICIENTE_REDIMENSION)
	}
	heap.datos[heap.cantidad] = elemento
	upheap(heap.datos, heap.cantidad, heap.cmp)
	heap.cantidad++
}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	if heap.cantidad*_COEFICIENTE_DISMINUCION <= len(heap.datos) && len(heap.datos) >= _CAPACIDAD_INICIAL*2 {
		heap.redimensionar(len(heap.datos) / _COEFICIENTE_REDIMENSION)
	}
	maximo := heap.VerMax()
	swap(&heap.datos[0], &heap.datos[heap.cantidad-1])
	heap.cantidad--
	downheap(heap.datos, 0, heap.cmp, heap.cantidad)
	return maximo
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}
