package lista_test

import (
	"github.com/stretchr/testify/require"
	TDALista "lista"
	"testing"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.EqualValues(t, true, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestComportamiento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	datos := [6]string{
		"Algoritmos",
		"y",
		"Programación",
		"II",
		"Catedra",
		"Buchwald",
	}
	for _, dato := range datos {
		lista.InsertarPrimero(dato)
		lista.InsertarUltimo(dato)
	}
	for i := 5; i >= 0; i-- {
		require.EqualValues(t, datos[i], lista.BorrarPrimero())
	}
	for j := 0; j <= 5; j++ {
		require.EqualValues(t, datos[j], lista.BorrarPrimero())
	}
}

func TestVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i <= 10000; i++ {
		lista.InsertarPrimero(i)
		lista.InsertarUltimo(i)
	}
	for j := 10000; j >= 0; j-- {
		require.EqualValues(t, j, lista.BorrarPrimero())
	}
	for k := 0; k <= 10000; k++ {
		require.EqualValues(t, k, lista.BorrarPrimero())
	}

}
func TestIterExternoInsertarAlCrear(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	iter := lista.Iterador()
	iter.Insertar("Algoritmos")
	require.EqualValues(t, "Algoritmos", lista.VerPrimero())
	require.EqualValues(t, "Algoritmos", lista.VerUltimo())
	require.EqualValues(t, lista.VerPrimero(), iter.VerActual())
	require.EqualValues(t, lista.VerUltimo(), iter.VerActual())
	require.EqualValues(t, false, lista.EstaVacia())
	require.EqualValues(t, 1, lista.Largo())
}

func TestIterExternoInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	iter.Insertar(10)
	require.EqualValues(t, 10, lista.VerUltimo())
	require.EqualValues(t, lista.VerUltimo(), iter.VerActual())
	require.EqualValues(t, false, lista.EstaVacia())
	require.EqualValues(t, 11, lista.Largo())
}

func TestIterExternoInsertarAlMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	datos := [8]string{
		"TDA",
		"Lista",
		"Enlazada",
		"y",
		"Programación",
		"II",
		"Catedra",
		"Buchwald",
	}
	for _, dato := range datos {
		lista.InsertarUltimo(dato)
	}
	iter := lista.Iterador()
	for i := 0; iter.HaySiguiente(); i++ {
		if i == 3 {
			iter.Insertar("Algoritmos")
		}
		iter.Siguiente()
	}
	require.EqualValues(t, lista.VerPrimero(), lista.BorrarPrimero())
	require.EqualValues(t, lista.VerPrimero(), lista.BorrarPrimero())
	require.EqualValues(t, lista.VerPrimero(), lista.BorrarPrimero())
	require.EqualValues(t, "Algoritmos", lista.VerPrimero())
}

func TestIterExternoBorrarAlCrear(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	iter := lista.Iterador()
	require.EqualValues(t, 1, iter.Borrar())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 1, lista.Largo())
}

func TestIterExternoBorrarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(40)
	lista.InsertarPrimero(20)
	lista.InsertarPrimero(10)
	iter := lista.Iterador()
	for i := 0; iter.HaySiguiente(); i++ {
		if i == lista.Largo() {
			require.EqualValues(t, lista.VerUltimo(), iter.Borrar())
		}
		iter.Siguiente()
	}
}

func TestIterExternoBorrarAlMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	datos := [6]string{
		"Algoritmos",
		"y",
		"Programación",
		"II",
		"Catedra",
		"Buchwald",
	}
	for _, dato := range datos {
		lista.InsertarUltimo(dato)
	}
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		if iter.VerActual() == datos[3] {
			iter.Borrar()
		}
		iter.Siguiente()
	}
	require.EqualValues(t, lista.VerPrimero(), lista.BorrarPrimero())
	require.EqualValues(t, lista.VerPrimero(), lista.BorrarPrimero())
	require.EqualValues(t, lista.VerPrimero(), lista.BorrarPrimero())
	require.EqualValues(t, true, lista.VerPrimero() != "II")
	require.EqualValues(t, "Catedra", lista.VerPrimero())
}

func TestIterInternoSumaElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(40)
	lista.InsertarPrimero(30)
	lista.InsertarPrimero(20)
	lista.InsertarPrimero(10)
	contador := 0
	contador_ptr := &contador
	lista.Iterar(func(v int) bool {
		*contador_ptr += v
		return true
	})
	require.Equal(t, 100, contador)
}

func TestIterInternoSumaPrimeros2Elementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(40)
	lista.InsertarPrimero(30)
	lista.InsertarPrimero(20)
	lista.InsertarPrimero(10)
	contador := 0
	cant_elementos := 0
	contador_ptr := &contador
	lista.Iterar(func(v int) bool {
		*contador_ptr += v
		cant_elementos++
		return cant_elementos < 2
	})
	require.Equal(t, 30, contador)
}
