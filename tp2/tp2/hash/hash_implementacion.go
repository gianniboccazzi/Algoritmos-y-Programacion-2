package diccionario

import (
	"fmt"
)

type estado int

const (
	VACIO estado = iota
	OCUPADO
	BORRADO
)

const CAPACIDAD_INICIAL int = 20
const CRITERIO_AGRANDAR float32 = 0.7
const CRITERIO_ACHICAR float32 = 0.1

type celda[K comparable, V any] struct {
	clave  K
	dato   V
	estado estado
}

type hashCerrado[K comparable, V any] struct {
	datos        []celda[K, V]
	cantOcupados int
	cantBorrados int
}

type iteradorDiccionario[K comparable, V any] struct {
	posicion  int
	tablahash *hashCerrado[K, V]
}

/*
Link de la funcion de hashing:
https://golangprojectstructure.com/hash-functions-go-code/ */

func sdbmHash(data []byte) int {
	var hash int

	for _, b := range data {
		hash = int(b) + (hash << 6) + (hash << 16) - hash
	}

	return hash
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	diccionario := new(hashCerrado[K, V])
	diccionario.datos = make([]celda[K, V], CAPACIDAD_INICIAL)
	return diccionario
}

// obtenerPosicion devuelve la posición que le corresponde a la clave pasada luego de aplicarle la función de diccionario
func (diccionario hashCerrado[K, V]) obtenerPosicion(clave K, largo int) int {
	pos := sdbmHash(convertirABytes(clave)) % largo
	if pos < 0 {
		pos *= -1
	}
	return pos
}

// calcularCarga calcula el factor de carga teniendo en cuenta los ocupados y los borrados
func (diccionario *hashCerrado[K, V]) calcularCarga() float32 {
	carga := float32(diccionario.cantBorrados+diccionario.cantOcupados) / float32(len(diccionario.datos))
	return carga
}

// buscarClave recorre la tabla de diccionario y devuelve la posicion en donde se encuentra la clave pasada
// o en su defecto, la primera posición vacia que encuentra
func (diccionario *hashCerrado[K, V]) buscarClave(posicion int, clave K) int {
	for diccionario.datos[posicion].estado != VACIO {
		if diccionario.datos[posicion].clave == clave {
			return posicion
		}
		if posicion != len(diccionario.datos)-1 {
			posicion++
		} else {
			posicion = 0
		}
	}
	return posicion
}

// redimensionar crea una nueva tabla con la capacidad pasada y va guardando los elementos de la antigua
// tabla en la nueva
func (diccionario *hashCerrado[K, V]) redimensionar(nuevaCapacidad int) {
	tablaAux := diccionario.datos
	diccionario.datos = make([]celda[K, V], nuevaCapacidad)
	diccionario.cantOcupados = 0
	for _, celda := range tablaAux {
		if celda.estado == OCUPADO {
			diccionario.Guardar(celda.clave, celda.dato)
		}
	}
}

func (diccionario *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if diccionario.calcularCarga() >= CRITERIO_AGRANDAR {
		diccionario.redimensionar(len(diccionario.datos) * 2)
	}
	pos := diccionario.obtenerPosicion(clave, len(diccionario.datos))
	buscarClave := diccionario.buscarClave(pos, clave)
	if !diccionario.Pertenece(clave) {
		diccionario.datos[buscarClave].clave = clave
		diccionario.datos[buscarClave].estado = OCUPADO
		diccionario.cantOcupados++
	}
	diccionario.datos[buscarClave].dato = dato
}

func (diccionario hashCerrado[K, V]) Pertenece(clave K) bool {
	pos := diccionario.obtenerPosicion(clave, len(diccionario.datos))
	return diccionario.datos[diccionario.buscarClave(pos, clave)].estado == OCUPADO
}

func (diccionario hashCerrado[K, V]) Obtener(clave K) V {
	if !diccionario.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	pos := diccionario.obtenerPosicion(clave, len(diccionario.datos))
	return diccionario.datos[diccionario.buscarClave(pos, clave)].dato
}

func (diccionario *hashCerrado[K, V]) Borrar(clave K) V {
	if !diccionario.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	if diccionario.calcularCarga() <= CRITERIO_ACHICAR && len(diccionario.datos) > CAPACIDAD_INICIAL {
		diccionario.redimensionar(len(diccionario.datos) / 2)
	}
	pos := diccionario.obtenerPosicion(clave, len(diccionario.datos))
	buscarClave := diccionario.buscarClave(pos, clave)
	diccionario.datos[buscarClave].estado = BORRADO
	diccionario.cantOcupados--
	diccionario.cantBorrados++
	return diccionario.datos[buscarClave].dato
}

func (diccionario hashCerrado[K, V]) Cantidad() int {
	return diccionario.cantOcupados
}

func (diccionario *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	i, res := -1, true
	for i < len(diccionario.datos)-1 && res {
		i = proximoDato(diccionario.datos, i)
		if i > len(diccionario.datos)-1 {
			break
		}
		res = visitar(diccionario.datos[i].clave, diccionario.datos[i].dato)
	}
}

func (diccionario *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iteradorDiccionario[K, V])
	iter.tablahash = diccionario
	primeraPosicion := proximoDato(diccionario.datos, -1)
	iter.posicion = primeraPosicion
	return iter
}

func proximoDato[K comparable, V any](arreglo []celda[K, V], posicionAnterior int) int {
	for i := posicionAnterior + 1; i < len(arreglo); i++ {
		if arreglo[i].estado == OCUPADO {
			return i
		}
	}
	return len(arreglo)
}

func (iter iteradorDiccionario[K, V]) HaySiguiente() bool {
	return iter.posicion != len(iter.tablahash.datos)
}

func (iter iteradorDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.tablahash.datos[iter.posicion].clave, iter.tablahash.datos[iter.posicion].dato
}

func (iter *iteradorDiccionario[K, V]) Siguiente() K {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	clave := iter.tablahash.datos[iter.posicion].clave
	nuevaPosicion := proximoDato(iter.tablahash.datos, iter.posicion)
	iter.posicion = nuevaPosicion
	return clave
}
