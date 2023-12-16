package cola_prioridad_test

import (
	TDAHeap "cola_prioridad"
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN_HEAP = []int{12500, 25000, 50000, 100000, 200000, 400000}

func TestCrearHeapVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int {
		return a - b
	})
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestCrearHeapArrVacio(t *testing.T) {
	arreglo := make([]string, 0)
	heap := TDAHeap.CrearHeapArr(arreglo, func(a, b string) int {
		return strings.Compare(a, b)
	})
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapUnElemento(t *testing.T) {
	heap := TDAHeap.CrearHeap[string](func(a, b string) int {
		return strings.Compare(a, b)
	})
	heap.Encolar("Hola")
	require.EqualValues(t, 1, heap.Cantidad())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, "Hola", heap.VerMax())
	require.EqualValues(t, "Hola", heap.Desencolar())
}

func TestHeapEncolarUpheap(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int {
		return a - b
	})
	datos := []int{40, 27, 13, 18, 22, 3, 10}
	for _, dato := range datos {
		heap.Encolar(dato)
	}
	heap.Encolar(30)
	require.EqualValues(t, 40, heap.VerMax())
	heap.Encolar(50)
	require.EqualValues(t, 50, heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, 40, heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, 30, heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, 27, heap.VerMax())
	heap.Encolar(80)
	require.EqualValues(t, 80, heap.VerMax())
	heap.Encolar(1200)
	require.EqualValues(t, 1200, heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, 80, heap.VerMax())
}

func TestHeapDesencolarDownheap(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int {
		return a - b
	})
	datos := []int{40, 30, 13, 27, 22, 3, 10, 18}
	for _, dato := range datos {
		heap.Encolar(dato)
	}
	require.EqualValues(t, heap.VerMax(), heap.Desencolar())
	require.EqualValues(t, len(datos)-1, heap.Cantidad())
	require.EqualValues(t, heap.VerMax(), heap.Desencolar())
}

func TestHeapDesencolarHastaVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int {
		return a - b
	})
	datos := []int{40, 30, 13, 27, 22, 3, 10, 18, 44, 677, 99, 00, 45, 86, 36, 94, 25, 85, 37, 75, 47, 86}
	for _, dato := range datos {
		heap.Encolar(dato)
	}
	cantidad := len(datos)
	for !heap.EstaVacia() {
		require.EqualValues(t, heap.VerMax(), heap.Desencolar())
		cantidad--
		require.EqualValues(t, cantidad, heap.Cantidad())
	}
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapSort(t *testing.T) {
	arreglo := []int{50, 1, 2, 3, 30, 10, 80}
	TDAHeap.HeapSort(arreglo, func(a, b int) int {
		return a - b
	})
	require.EqualValues(t, []int{1, 2, 3, 10, 30, 50, 80}, arreglo)
}

func TestCrearHeapArr(t *testing.T) {
	arreglo := []int{7, 1, 2, 3, 30, 10, 80}
	heap := TDAHeap.CrearHeapArr[int](arreglo, func(a, b int) int {
		return a - b
	})
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 80, heap.VerMax())
	require.EqualValues(t, 7, heap.Cantidad())
	datos := []int{5, 3, 14, 23, 26, 1, 33, 45, 87, 98, 150}
	for _, dato := range datos {
		heap.Encolar(dato)
	}
	for !heap.EstaVacia() {
		require.EqualValues(t, heap.VerMax(), heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func ejecutarPruebaVolumenHeap(b *testing.B, n int) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int {
		return a - b
	})
	datos := make([]int, n)
	for i := 0; i < n; i++ {
		datos[i] = n - 1 - i
	}

	valores_mezclados := rand.Perm(n)
	/* Inserta 'n' strings en el heap */
	for _, val := range valores_mezclados {
		heap.Encolar(val)
	}

	require.EqualValues(b, n, heap.Cantidad(), "La cantidad de elementos es incorrecta")
	require.False(b, heap.EstaVacia())
	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		res := heap.Desencolar()
		ok = res == datos[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Desencolar funciona correctamente")
	require.EqualValues(b, 0, heap.Cantidad(), "La cantidad de elementos es incorrecta")
	require.True(b, heap.EstaVacia(), "El heap no termina vacio")
}

func BenchmarkDiccionarioOrdenado(b *testing.B) {
	b.Log("Prueba de stress del heap. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN_HEAP {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenHeap(b, n)
			}
		})
	}
}
