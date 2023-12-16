package algograms

import (
	"fmt"
	"strings"
	. "tp2/abb"
	. "tp2/errores"
)

type postImplementacion struct {
	publicacion   string
	autor         string
	id            int
	likes         DiccionarioOrdenado[string, Usuario]
	afinidadAutor int
}

func CrearPost(mensaje string, autor string, id int, afinidadAutor int) Post {
	post := new(postImplementacion)
	post.publicacion = mensaje
	post.autor = autor
	post.id = id
	post.likes = CrearABB[string, Usuario](func(k, k2 string) int {
		return strings.Compare(k, k2)
	})
	post.afinidadAutor = afinidadAutor
	return post
}

func (post postImplementacion) MostrarLikes() string {
	if post.likes.Cantidad() == 0 {
		return ErrorPostInexistenteoSinLikes{}.Error()
	}
	cantidadLikes := fmt.Sprintf("El post tiene %d likes:\n", post.likes.Cantidad())
	var autores string
	iter := post.likes.Iterador()
	for iter.HaySiguiente() {
		usuario, _ := iter.VerActual()
		autores += fmt.Sprintf("\t%s\n", usuario)
		iter.Siguiente()
	}
	return cantidadLikes + autores
}

func (post *postImplementacion) AgregarLike(usuario Usuario) {
	if post.YaDioLike(usuario.VerNombre()) {
		return
	}
	post.likes.Guardar(usuario.VerNombre(), usuario)
}

func (post postImplementacion) VerPublicacion() string {
	return post.publicacion
}

func (post postImplementacion) VerAutor() string {
	return post.autor
}

func (post postImplementacion) VerID() int {
	return post.id
}

func (post postImplementacion) VerAfinidad() int {
	return post.afinidadAutor
}

func (post postImplementacion) CantidadLikes() int {
	return post.likes.Cantidad()
}

func (post postImplementacion) YaDioLike(nombre string) bool {
	return post.likes.Pertenece(nombre)

}

func (post postImplementacion) NoTieneLikes() bool {
	return post.likes.Cantidad() == 0

}
