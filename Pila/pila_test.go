package pila_test

import (
	TDAPila "pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

}

func TestPilaStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("Ciclon")
	pila.Apilar("teamo")
	pila.Apilar("Cuervos")
	require.Equal(t, "Cuervos", pila.VerTope())
	require.Equal(t, "Cuervos", pila.Desapilar())
	require.False(t, pila.EstaVacia())
	require.Equal(t, "teamo", pila.VerTope())
	require.Equal(t, "teamo", pila.Desapilar())
	require.Equal(t, "Ciclon", pila.Desapilar())
	require.True(t, pila.EstaVacia())

}

func TestPilaEnteros(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	require.Equal(t, 2, pila.VerTope())
	pila.Apilar(3)
	require.Equal(t, 3, pila.VerTope())
	require.Equal(t, 3, pila.Desapilar())
	require.False(t, pila.EstaVacia())
	require.Equal(t, 2, pila.VerTope())
	require.Equal(t, 2, pila.Desapilar())
	require.Equal(t, 1, pila.Desapilar())
	require.True(t, pila.EstaVacia())

}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i == 1000; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}
	for i := 1000; i == 0; i-- {
		require.Equal(t, i, pila.VerTope())
		require.Equal(t, i, pila.Desapilar())
	}
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.True(t, pila.EstaVacia())

}
