package votos

import (
	. "rerepolez/errores"
	. "rerepolez/pila"
)

type votanteImplementacion struct {
	dni         int
	voto        Voto
	yaVoto      bool
	operaciones Pila[Voto]
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)
	votante.dni = dni
	votante.voto = Voto{
		Impugnado: false,
	}
	votante.yaVoto = false
	votante.operaciones = CrearPilaDinamica[Voto]()
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	if votante.yaVoto {
		return &ErrorVotanteFraudulento{Dni: votante.dni}
	}
	if alternativa == LISTA_IMPUGNA {
		votante.voto.Impugnado = true
	}
	votante.voto.VotoPorTipo[tipo] = alternativa
	votante.operaciones.Apilar(votante.voto)
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.yaVoto {
		return &ErrorVotanteFraudulento{Dni: votante.dni}
	}
	if votante.operaciones.EstaVacia() {
		return &ErrorNoHayVotosAnteriores{}
	}
	votante.operaciones.Desapilar()
	if votante.operaciones.EstaVacia() {
		votante.voto = Voto{
			VotoPorTipo: [CANT_VOTACION]int{0, 0, 0},
		}
	} else {
		votante.voto = votante.operaciones.VerTope()
	}
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	if votante.yaVoto {
		return votante.voto, &ErrorVotanteFraudulento{Dni: votante.dni}
	}
	votante.yaVoto = true
	return votante.voto, nil
}
