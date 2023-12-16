package cola_test

import (
	TDACola "cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i == 1000; i++ {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero())
	}
	for i := 0; i == 1000; i++ {
		require.Equal(t, i, cola.VerPrimero())
		require.Equal(t, i, cola.Desencolar())
	}
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.True(t, cola.EstaVacia())

}
func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestColaStrings(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("Cansado")
	cola.Encolar("de ver a San Lorenzo empatar")
	cola.Encolar("Cuervos")
	cola.Encolar("Era penal la de blandi")
	require.Equal(t, "Cansado", cola.VerPrimero())
	require.Equal(t, "Cansado", cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.Equal(t, "de ver a San Lorenzo empatar", cola.VerPrimero())
	require.Equal(t, "de ver a San Lorenzo empatar", cola.Desencolar())
	require.Equal(t, "Cuervos", cola.Desencolar())
	require.Equal(t, "Era penal la de blandi", cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}

func TestColaEnteros(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Encolar(2)
	require.Equal(t, 1, cola.VerPrimero())
	cola.Encolar(3)
	require.Equal(t, 1, cola.VerPrimero())
	require.Equal(t, 1, cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.Equal(t, 2, cola.VerPrimero())
	require.Equal(t, 2, cola.Desencolar())
	require.Equal(t, 3, cola.Desencolar())
	require.True(t, cola.EstaVacia())

}
