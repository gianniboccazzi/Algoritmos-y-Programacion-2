package votos

import (
	"fmt"
)

type partidoImplementacion struct {
	nombre         string
	candidatos     [CANT_VOTACION]string
	votosRecibidos [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	votosRecibidos [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos []string) Partido {
	partido := new(partidoImplementacion)
	partido.nombre = nombre
	for i, _ := range candidatos {
		partido.candidatos[i] = candidatos[i]
	}
	return partido
}

func CrearVotosEnBlanco() Partido {
	return new(partidoEnBlanco)
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.votosRecibidos[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	if partido.votosRecibidos[tipo] == 1 {
		return fmt.Sprintf("%s - %s: %d voto", partido.nombre, partido.candidatos[tipo], partido.votosRecibidos[tipo])
	}
	return fmt.Sprintf("%s - %s: %d votos", partido.nombre, partido.candidatos[tipo], partido.votosRecibidos[tipo])
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votosRecibidos[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	if blanco.votosRecibidos[tipo] == 1 {
		return fmt.Sprintf("Votos en Blanco: %d voto", blanco.votosRecibidos[tipo])
	}
	return fmt.Sprintf("Votos en Blanco: %d votos", blanco.votosRecibidos[tipo])
}
