package algograms

type Algogram interface {
	Login(usuario string) string

	Logout() string

	Publicar(publicacion string) string

	VerSiguienteFeed() string

	LikearPost(id int) string

	MostrarLikes(id int) string

	VerUsuarioLoggeado() Usuario
}
