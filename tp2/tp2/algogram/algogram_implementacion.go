package algograms

import (
	"fmt"
	. "tp2/errores"
	. "tp2/hash"
)

type algogramImplementacion struct {
	usuarios        Diccionario[string, Usuario]
	posts           Diccionario[int, Post]
	usuarioLoggeado Usuario
}

func CrearAlgogram(usuarios Diccionario[string, Usuario]) Algogram {
	algogram := new(algogramImplementacion)
	algogram.usuarios = usuarios
	algogram.posts = CrearHash[int, Post]()
	algogram.usuarioLoggeado = nil
	return algogram
}

func (algogram *algogramImplementacion) Login(usuario string) string {
	if algogram.usuarioLoggeado != nil {
		return ErrorUsuarioYaLoggeado{}.Error()
	}
	if !algogram.usuarios.Pertenece(usuario) {
		return ErrorUsuarioNoExiste{}.Error()
	}
	algogram.usuarioLoggeado = algogram.usuarios.Obtener(usuario)
	return fmt.Sprintf("Hola %s", usuario)
}

func (algogram *algogramImplementacion) Logout() string {
	if algogram.usuarioLoggeado != nil {
		algogram.usuarioLoggeado = nil
		return fmt.Sprintf("Adios")
	}
	return ErrorUsuarioNoLoggeado{}.Error()
}

func (algogram *algogramImplementacion) Publicar(publicacion string) string {
	if algogram.usuarioLoggeado == nil {
		return ErrorUsuarioNoLoggeado{}.Error()
	}
	post := CrearPost(publicacion, algogram.usuarioLoggeado.VerNombre(), algogram.posts.Cantidad(), algogram.usuarioLoggeado.VerAfinidad())
	algogram.posts.Guardar(post.VerID(), post)

	// Se agrega a los feeds del resto de usuarios
	iter := algogram.usuarios.Iterador()
	for iter.HaySiguiente() {
		nombre, usuario := iter.VerActual()
		if nombre != algogram.usuarioLoggeado.VerNombre() {
			usuario.AgregarAlFeed(post)
		}
		algogram.usuarios.Guardar(nombre, usuario)
		iter.Siguiente()
	}
	return fmt.Sprintf("Post publicado")
}

func (algogram *algogramImplementacion) VerSiguienteFeed() string {
	if algogram.usuarioLoggeado == nil || !algogram.usuarioLoggeado.HayPostsSinVer() {
		return "Usuario no loggeado o no hay mas posts para ver."
	}
	return algogram.usuarioLoggeado.VerSiguientePublicacion()
}

func (algogram *algogramImplementacion) LikearPost(id int) string {
	if algogram.usuarioLoggeado == nil || !algogram.posts.Pertenece(id) {
		return ErrorUsuarioNoLoggeadoOPostInexistente{}.Error()
	}
	algogram.posts.Obtener(id).AgregarLike(algogram.usuarioLoggeado)
	return fmt.Sprintf("Post likeado")
}

func (algogram algogramImplementacion) MostrarLikes(id int) string {
	if !algogram.posts.Pertenece(id) || algogram.posts.Obtener(id).NoTieneLikes() {
		return ErrorPostInexistenteoSinLikes{}.Error()
	}
	return algogram.posts.Obtener(id).MostrarLikes()
}

func (algogram algogramImplementacion) VerUsuarioLoggeado() Usuario {
	return algogram.usuarioLoggeado
}
