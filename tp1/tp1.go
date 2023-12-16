package main

import (
	"bufio"
	"fmt"
	"os"
	. "rerepolez/cola"
	. "rerepolez/errores"
	. "rerepolez/votos"
	"strconv"
	"strings"
)

func countingPorCriterio(votantes []Votante, criterio func(int) int) []Votante {
	colas := make([]Cola[Votante], 10)
	for i := range colas {
		colas[i] = CrearColaEnlazada[Votante]()
	}

	for _, votante := range votantes {
		colas[criterio(votante.LeerDNI())].Encolar(votante)
	}

	arrOrdenado := make([]Votante, len(votantes))
	i := 0
	for _, cola := range colas {
		for !cola.EstaVacia() {
			arrOrdenado[i] = cola.Desencolar()
			i++
		}
	}
	return arrOrdenado
}

func ordenarVotantesPorDNI(votantes []Votante) []Votante {
	ordenadoPorUnidad := countingPorCriterio(votantes, func(num int) int {
		return num % 10
	})
	ordenadoPorDecena := countingPorCriterio(ordenadoPorUnidad, func(num int) int {
		return (num / 10) % 10
	})
	ordenadoPorCentena := countingPorCriterio(ordenadoPorDecena, func(num int) int {
		return (num / 100) % 10
	})
	ordenadoPorUnidadDeMil := countingPorCriterio(ordenadoPorCentena, func(num int) int {
		return (num / 1000) % 10
	})
	ordenadoPorDecenaDeMil := countingPorCriterio(ordenadoPorUnidadDeMil, func(num int) int {
		return (num / 10000) % 10
	})
	ordenadoPorCentenaDeMil := countingPorCriterio(ordenadoPorDecenaDeMil, func(num int) int {
		return (num / 100000) % 10
	})
	ordenadoPorUnidadDeMillon := countingPorCriterio(ordenadoPorCentenaDeMil, func(num int) int {
		return (num / 1000000) % 10
	})
	ordenadoPorDecenaDeMillon := countingPorCriterio(ordenadoPorUnidadDeMillon, func(num int) int {
		return (num / 10000000) % 10
	})
	return ordenadoPorDecenaDeMillon
}

func busquedaBinaria(votantes []Votante, inicio, fin, elemento int) int {
	if inicio > fin {
		return -1
	}
	medio := (inicio + fin) / 2
	if votantes[medio].LeerDNI() == elemento {
		return medio
	}
	if votantes[medio].LeerDNI() < elemento {
		return busquedaBinaria(votantes, medio+1, fin, elemento)
	} else {
		return busquedaBinaria(votantes, inicio, medio-1, elemento)
	}
}

func leerPartidos(ruta string) ([]Partido, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, &ErrorLeerArchivo{}
	}
	defer archivo.Close()

	var partidos []Partido
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		params := strings.Split(s.Text(), ",")
		nuevoPartido := CrearPartido(params[0], params[1:])
		partidos = append(partidos, nuevoPartido)
	}
	return partidos, nil
}

func leerPadrones(ruta string) ([]Votante, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, &ErrorLeerArchivo{}
	}
	defer archivo.Close()

	var votantes []Votante
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		dni, _ := strconv.Atoi(s.Text())
		nuevoVotante := CrearVotante(dni)
		votantes = append(votantes, nuevoVotante)
	}
	return ordenarVotantesPorDNI(votantes), nil
}

func ejecutarIngresar(votantes []Votante, entrada []string, fila Cola[Votante]) {
	if len(entrada) <= 1 {
		return
	}
	dniIngresado, _ := strconv.Atoi(entrada[1])
	if dniIngresado <= 0 {
		fmt.Printf(DNIError{}.Error() + "\n")
	} else {
		indice := busquedaBinaria(votantes, 0, len(votantes)-1, dniIngresado)
		if indice != -1 {
			fila.Encolar(votantes[indice])
			fmt.Printf("OK\n")
		} else {
			fmt.Printf(DNIFueraPadron{}.Error() + "\n")
		}
	}
}

func ejecutarVotar(partidos []Partido, entrada []string, fila Cola[Votante]) {
	if len(entrada) <= 1 {
		return
	}
	if fila.EstaVacia() {
		fmt.Printf(FilaVacia{}.Error() + "\n")
	} else {
		votante := fila.VerPrimero()
		var tipo TipoVoto
		switch entrada[1] {
		case "Presidente":
			tipo = PRESIDENTE
		case "Gobernador":
			tipo = GOBERNADOR
		case "Intendente":
			tipo = INTENDENTE
		default:
			fmt.Printf(ErrorTipoVoto{}.Error() + "\n")
			return
		}
		numeroLista, err := strconv.Atoi(entrada[2])
		if numeroLista > len(partidos) || err != nil {
			fmt.Printf(ErrorAlternativaInvalida{}.Error() + "\n")
		} else {
			err := votante.Votar(tipo, numeroLista)
			if err != nil {
				fila.Desencolar()
				fmt.Printf(err.Error() + "\n")
			} else {
				fmt.Printf("OK\n")
			}
		}
	}
}

func ejecutarDeshacer(fila Cola[Votante]) {
	if fila.EstaVacia() {
		fmt.Printf(FilaVacia{}.Error() + "\n")
	} else {
		votante := fila.VerPrimero()
		err := votante.Deshacer()
		if err != nil {
			switch err.(type) {
			case *ErrorVotanteFraudulento:
				fila.Desencolar()
				fmt.Printf(err.Error() + "\n")
			default:
				fmt.Printf(err.Error() + "\n")
			}
		} else {
			fmt.Printf("OK\n")
		}
	}
}

func ejecutarFinVotar(fila Cola[Votante], partidos []Partido, blanco Partido, impugnados *int) {
	if fila.EstaVacia() {
		fmt.Printf(FilaVacia{}.Error() + "\n")
	} else {
		votante := fila.VerPrimero()
		votoFinal, err := votante.FinVoto()
		if err != nil {
			fila.Desencolar()
			fmt.Printf(err.Error() + "\n")
		} else {
			if votoFinal.Impugnado {
				*impugnados++
			} else {
				sumarVoto(votoFinal, blanco, partidos, PRESIDENTE)
				sumarVoto(votoFinal, blanco, partidos, INTENDENTE)
				sumarVoto(votoFinal, blanco, partidos, GOBERNADOR)
			}
			fila.Desencolar()
			fmt.Printf("OK\n")
		}
	}
}

func sumarVoto(votofinal Voto, blanco Partido, partidos []Partido, tipo TipoVoto) {
	if votofinal.VotoPorTipo[tipo] == 0 {
		blanco.VotadoPara(tipo)
	} else {
		partidos[votofinal.VotoPorTipo[tipo]-1].VotadoPara(tipo)
	}
}

func ejecutarComandos(votantes []Votante, partidos []Partido, entrada []string, fila Cola[Votante], blanco Partido, impugnados *int) {
	switch entrada[0] {
	case "ingresar":
		ejecutarIngresar(votantes, entrada, fila)
	case "votar":
		ejecutarVotar(partidos, entrada, fila)
	case "deshacer":
		ejecutarDeshacer(fila)
	case "fin-votar":
		ejecutarFinVotar(fila, partidos, blanco, impugnados)
	}
}

func ejecutarPrograma(votantes []Votante, partidos []Partido, blanco Partido, impugnados *int, fila Cola[Votante]) {
	var entrada []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		entrada = strings.Split(s.Text(), " ")
		ejecutarComandos(votantes, partidos, entrada, fila, blanco, impugnados)
	}
}

func imprimirResultados(partidos []Partido, blanco Partido, impugnados int, fila Cola[Votante]) {
	if !fila.EstaVacia() {
		fmt.Printf(ErrorCiudadanosSinVotar{}.Error() + "\n")
	}
	fmt.Printf("Presidente:" + "\n")
	imprimirPartido(partidos, PRESIDENTE, blanco)
	fmt.Printf("Gobernador:" + "\n")
	imprimirPartido(partidos, GOBERNADOR, blanco)
	fmt.Printf("Intendente:" + "\n")
	imprimirPartido(partidos, INTENDENTE, blanco)
	if impugnados == 1 {
		fmt.Printf("Votos Impugnados: %d voto\n", impugnados)
	} else {
		fmt.Printf("Votos Impugnados: %d votos\n", impugnados)
	}
}

func imprimirPartido(partidos []Partido, tipo TipoVoto, blanco Partido) {
	fmt.Printf(blanco.ObtenerResultado(tipo) + "\n")
	for _, partido := range partidos {
		fmt.Printf(partido.ObtenerResultado(tipo) + "\n")
	}
	fmt.Printf("\n")
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Printf(ErrorParametros{}.Error() + "\n")
		return
	}
	fila := CrearColaEnlazada[Votante]()
	partidos, errPartidos := leerPartidos(args[0])
	votantes, errVotantes := leerPadrones(args[1])
	votosEnBlanco, impugnados := CrearVotosEnBlanco(), 0
	if errPartidos != nil || errVotantes != nil {
		fmt.Printf(ErrorLeerArchivo{}.Error() + "\n")
		return
	}
	ejecutarPrograma(votantes, partidos, votosEnBlanco, &impugnados, fila)
	imprimirResultados(partidos, votosEnBlanco, impugnados, fila)
}
