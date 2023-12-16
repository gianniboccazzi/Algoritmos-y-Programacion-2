package diccionario

import (
	. "tp2/abb/pila"
)

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type iterAbb[K comparable, V any] struct {
	datos Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
	cmp   func(K, K) int
}

func (diccionario *abb[K, V]) crearNodo(clave K, valor V) *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	nodo.clave = clave
	nodo.dato = valor
	return nodo
}

func (diccionario *abb[K, V]) ubicacionNodo(nodo *nodoAbb[K, V], clave K) *nodoAbb[K, V] {
	if nodo == nil {
		return nil
	}
	comparacion := diccionario.cmp(nodo.clave, clave)
	if comparacion == 0 {
		return nodo
	}
	if comparacion > 0 {
		return diccionario.ubicacionNodo(nodo.izquierdo, clave)
	}
	return diccionario.ubicacionNodo(nodo.derecho, clave)
}

func (diccionario *abb[K, V]) ubicacionPadre(nodo *nodoAbb[K, V], clave K) *nodoAbb[K, V] {
	if diccionario.cmp(nodo.clave, clave) == 0 {
		return nil
	}
	if (nodo.izquierdo != nil && diccionario.cmp(nodo.izquierdo.clave, clave) == 0) || (nodo.derecho != nil && diccionario.cmp(nodo.derecho.clave, clave) == 0) {
		return nodo
	}
	if diccionario.cmp(nodo.clave, clave) < 0 {
		if nodo.derecho == nil {
			return nodo
		}
		return diccionario.ubicacionPadre(nodo.derecho, clave)
	} else {
		if nodo.izquierdo == nil {
			return nodo
		}
	}
	return diccionario.ubicacionPadre(nodo.izquierdo, clave)
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.raiz = nil
	abb.cantidad = 0
	abb.cmp = funcion_cmp
	return abb
}

func (diccionario *abb[K, V]) guardarEnDiccionarioNoVacio(clave K, dato V) {
	if diccionario.Pertenece(clave) {
		nodo := diccionario.ubicacionNodo(diccionario.raiz, clave)
		nodo.dato = dato
	} else {
		padre := diccionario.ubicacionPadre(diccionario.raiz, clave)
		comparacion := diccionario.cmp(padre.clave, clave)
		if comparacion > 0 {
			padre.izquierdo = diccionario.crearNodo(clave, dato)
		} else if comparacion < 0 {
			padre.derecho = diccionario.crearNodo(clave, dato)
		} else {
			padre.dato = dato
		}
		diccionario.cantidad++
	}
}

func (diccionario *abb[K, V]) Guardar(clave K, dato V) {
	if diccionario.raiz == nil {
		diccionario.raiz = diccionario.crearNodo(clave, dato)
		diccionario.cantidad++
	} else {
		diccionario.guardarEnDiccionarioNoVacio(clave, dato)
	}
}

func (diccionario *abb[K, V]) Pertenece(clave K) bool {
	return diccionario.ubicacionNodo(diccionario.raiz, clave) != nil
}

func (diccionario *abb[K, V]) Obtener(clave K) V {
	nodo := diccionario.ubicacionNodo(diccionario.raiz, clave)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

func (diccionario *abb[K, V]) Borrar(clave K) V {
	if !diccionario.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	nodoABorrar := diccionario.ubicacionNodo(diccionario.raiz, clave)
	padre, datoBorrado := diccionario.ubicacionPadre(diccionario.raiz, nodoABorrar.clave), nodoABorrar.dato
	if nodoABorrar.izquierdo == nil && nodoABorrar.derecho == nil {
		diccionario.borrarSinHijos(nodoABorrar, padre)
		diccionario.cantidad--
	} else if (nodoABorrar.izquierdo == nil && nodoABorrar.derecho != nil) || (nodoABorrar.izquierdo != nil && nodoABorrar.derecho == nil) {
		diccionario.borrarConUnHijo(nodoABorrar, padre)
		diccionario.cantidad--
	} else {
		reemplazante := diccionario.encontrarReemplazante(nodoABorrar)
		nodoABorrar.dato = diccionario.Borrar(reemplazante.clave)
		nodoABorrar.clave = reemplazante.clave
	}
	return datoBorrado
}

func (diccionario *abb[K, V]) borrarSinHijos(nodo *nodoAbb[K, V], padre *nodoAbb[K, V]) {
	if padre == nil {
		diccionario.raiz = nil
		return
	}
	if padre.derecho == nil || padre.derecho.clave != nodo.clave {
		padre.izquierdo = nil
	} else {
		padre.derecho = nil
	}
}

func (diccionario *abb[K, V]) borrarConUnHijo(nodo *nodoAbb[K, V], padre *nodoAbb[K, V]) {
	var hijo *nodoAbb[K, V]
	if nodo.derecho == nil {
		hijo = nodo.izquierdo
	} else {
		hijo = nodo.derecho
	}
	if padre == nil {
		diccionario.raiz = hijo
		return
	}
	if padre.derecho == nil || padre.derecho.clave != nodo.clave {
		padre.izquierdo = hijo
	} else {
		padre.derecho = hijo
	}
}

//func (diccionario *abb[K, V]) borrarConDosHijos(nodo *nodoAbb[K, V], padre *nodoAbb[K, V]) {
//	if padre == nil {
//		return
//	}
//	reemplazante := diccionario.encontrarReemplazante(nodo)
//	padreReemplazante := diccionario.ubicacionPadre(diccionario.raiz, reemplazante.clave)
//	if padreReemplazante.derecho != nil && padreReemplazante.derecho.clave == reemplazante.clave {
//		padreReemplazante.derecho = nil
//	} else {
//		padreReemplazante.izquierdo = nil
//	}
//	nodo.clave = reemplazante.clave
//}

func (diccionario *abb[K, V]) encontrarReemplazante(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	reemplazante := nodo.izquierdo
	for reemplazante.derecho != nil {
		reemplazante = reemplazante.derecho
	}
	return reemplazante
}

func (diccionario *abb[K, V]) Cantidad() int {
	return diccionario.cantidad
}

func (nodo *nodoAbb[K, V]) iterarSinRango(visitar func(clave K, dato V) bool) {
	if nodo == nil {
		return
	}
	nodo.izquierdo.iterarSinRango(visitar)
	res := visitar(nodo.clave, nodo.dato)
	if !res {
		return
	}
	nodo.derecho.iterarSinRango(visitar)
}

func (nodo *nodoAbb[K, V]) iterar(desde *K, hasta *K, visitar func(clave K, dato V) bool, funcion_cmp func(K, K) int) {
	if nodo == nil {
		return
	}
	if desde == nil && hasta == nil {
		nodo.iterarSinRango(visitar)
	} else {
		if funcion_cmp(nodo.clave, *desde) > 0 {
			nodo.izquierdo.iterar(desde, hasta, visitar, funcion_cmp)
		}
		if funcion_cmp(nodo.clave, *desde) >= 0 && funcion_cmp(nodo.clave, *hasta) <= 0 {
			res := visitar(nodo.clave, nodo.dato)
			if !res {
				return
			}
		}
		if funcion_cmp(nodo.clave, *hasta) < 0 {
			nodo.derecho.iterar(desde, hasta, visitar, funcion_cmp)
		}
	}
}

func apilarSinRango[K comparable, V any](nodo *nodoAbb[K, V], pila Pila[*nodoAbb[K, V]]) {
	if nodo == nil {
		return
	}
	pila.Apilar(nodo)
	apilarSinRango(nodo.izquierdo, pila)
}

func (iter *iterAbb[K, V]) apilar(nodo *nodoAbb[K, V], pila Pila[*nodoAbb[K, V]], funcion_cmp func(K, K) int) {
	if nodo == nil {
		return
	}
	if iter.desde == nil && iter.hasta == nil {
		apilarSinRango(nodo, pila)
	} else {
		if (funcion_cmp(nodo.clave, *iter.desde) >= 0 || iter.desde == nil) && (funcion_cmp(nodo.clave, *iter.hasta) <= 0 || iter.hasta == nil) {
			pila.Apilar(nodo)
			iter.apilar(nodo.izquierdo, pila, funcion_cmp)
		}
		if funcion_cmp(nodo.clave, *iter.desde) > 0 || iter.desde == nil {
			iter.apilar(nodo.izquierdo, pila, funcion_cmp)
		} else if funcion_cmp(nodo.clave, *iter.hasta) < 0 || iter.hasta == nil {
			iter.apilar(nodo.derecho, pila, funcion_cmp)
		}
	}
}

func (diccionario abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	diccionario.raiz.iterar(nil, nil, visitar, nil)
}

func (diccionario abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterAbb[K, V])
	iter.datos = CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.desde = nil
	iter.hasta = nil
	iter.apilar(diccionario.raiz, iter.datos, nil)
	return iter
}

func (diccionario *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	diccionario.raiz.iterar(desde, hasta, visitar, diccionario.cmp)
}

func (diccionario *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iterAbb[K, V])
	iter.cmp = diccionario.cmp
	iter.datos = CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.desde = desde
	iter.hasta = hasta
	iter.apilar(diccionario.raiz, iter.datos, diccionario.cmp)
	return iter
}

func (iter iterAbb[K, V]) HaySiguiente() bool {
	return !iter.datos.EstaVacia()
}

func (iter iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.datos.VerTope().clave, iter.datos.VerTope().dato
}

func (iter *iterAbb[K, V]) Siguiente() K {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	desapilado := iter.datos.Desapilar()
	if desapilado.derecho != nil {
		iter.apilar(desapilado.derecho, iter.datos, iter.cmp)
	}
	return desapilado.clave
}
