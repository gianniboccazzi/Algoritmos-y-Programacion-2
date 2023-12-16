package algograms

// Post modela una publicación, con sus alternativas para cada uno de los tipos de votaciones
type Post interface {
	MostrarLikes() string

	AgregarLike(usuario Usuario)

	VerPublicacion() string

	VerAutor() string

	VerID() int

	VerAfinidad() int

	CantidadLikes() int

	NoTieneLikes() bool
}
