package errores

type ErrorUsuarioYaLoggeado struct{}

func (e ErrorUsuarioYaLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUsuarioNoExiste struct{}

func (e ErrorUsuarioNoExiste) Error() string {
	return "Error: usuario no existente"
}

type ErrorUsuarioNoLoggeado struct{}

func (e ErrorUsuarioNoLoggeado) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "Error: Lectura de archivos"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "Error: Faltan parámetros"
}

type ErrorPostInexistenteoSinLikes struct{}

func (e ErrorPostInexistenteoSinLikes) Error() string {
	return "Error: Post inexistente o sin likes"
}

type ErrorUsuarioNoLoggeadoOPostInexistente struct{}

func (e ErrorUsuarioNoLoggeadoOPostInexistente) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}
