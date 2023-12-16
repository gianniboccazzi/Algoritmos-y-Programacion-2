package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	. "tp2/algogram"
	. "tp2/errores"
	. "tp2/hash"
)

func leerUsuarios(ruta string) (Algogram, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, &ErrorLeerArchivo{}
	}
	defer archivo.Close()

	usuarios := CrearHash[string, Usuario]()
	s := bufio.NewScanner(archivo)
	afinidad := 0
	for s.Scan() {
		nuevoUsuario := CrearUsuario(s.Text(), afinidad)
		usuarios.Guardar(nuevoUsuario.VerNombre(), nuevoUsuario)
		afinidad++
	}
	return CrearAlgogram(usuarios), nil
}

func ejecutarComandos(algogram Algogram, entrada []string) {
	publicacion := strings.Join(entrada[1:], " ")
	switch entrada[0] {
	case "login":
		fmt.Printf(algogram.Login(entrada[1]) + "\n")
	case "logout":
		fmt.Printf(algogram.Logout() + "\n")
	case "publicar":
		fmt.Printf(algogram.Publicar(publicacion) + "\n")
	case "ver_siguiente_feed":
		fmt.Printf(algogram.VerSiguienteFeed() + "\n")
	case "likear_post":
		id, _ := strconv.Atoi(entrada[1])
		fmt.Printf(algogram.LikearPost(id) + "\n")
	case "mostrar_likes":
		id, _ := strconv.Atoi(entrada[1])
		fmt.Printf(algogram.MostrarLikes(id) + "\n")
	}
}

func ejecutarPrograma(algogram Algogram) {
	var entrada []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		entrada = strings.Split(s.Text(), " ")
		ejecutarComandos(algogram, entrada)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf(ErrorParametros{}.Error() + "\n")
		return
	}
	algogram, err := leerUsuarios(args[0])
	if err != nil {
		fmt.Printf(ErrorLeerArchivo{}.Error() + "\n")
		return
	}
	ejecutarPrograma(algogram)
}
