package algograms

import (
	"fmt"
	"math"
	. "tp2/heap"
)

type usuarioImplementacion struct {
	nombre   string
	feed     ColaPrioridad[Post]
	afinidad int
}

func CrearUsuario(nombre string, afinidad int) Usuario {
	usuario := new(usuarioImplementacion)
	usuario.nombre = nombre
	usuario.afinidad = afinidad
	usuario.feed = CrearHeap[Post](func(a, b Post) int {
		if math.Abs(float64(a.VerAfinidad())-float64(usuario.afinidad)) < math.Abs(float64(b.VerAfinidad())-float64(usuario.afinidad)) {
			return 1
		}
		if math.Abs(float64(a.VerAfinidad())-float64(usuario.afinidad)) > math.Abs(float64(b.VerAfinidad())-float64(usuario.afinidad)) {
			return -1
		}
		if a.VerID() < b.VerID() {
			return 1
		}
		return -1
	})

	return usuario
}

func (usuario *usuarioImplementacion) AgregarAlFeed(post Post) {
	usuario.feed.Encolar(post)
}

func (usuario usuarioImplementacion) VerNombre() string {
	return usuario.nombre
}

func (usuario usuarioImplementacion) VerAfinidad() int {
	return usuario.afinidad
}

func (usuario *usuarioImplementacion) VerSiguientePublicacion() string {
	post := usuario.feed.Desencolar()
	fmt.Println("Post ID ", post.VerID())
	fmt.Println(post.VerAutor() + " dijo: " + post.VerPublicacion())
	return fmt.Sprintln("Likes: ", post.CantidadLikes())
}

func (usuario usuarioImplementacion) HayPostsSinVer() bool {
	return !usuario.feed.EstaVacia()
}
