package algograms

// Post modela una publicación , con sus alternativas para cada uno de los tipos de votaciones
type Usuario interface {
	AgregarAlFeed(post Post)

	VerNombre() string

	VerAfinidad() int

	VerSiguientePublicacion() string

	HayPostsSinVer() bool
}
