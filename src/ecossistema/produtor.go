package ecossistema

import (
	"fmt"
	"strconv"

	"utils"
)

type produtor struct {
	organismo

	_nomecompleto       string
	_adultociclo        int
	_reproduzirciclo    int
	_reproduzircontador int

	_vida int

	_produtores   (map[string]*produtor)
	_ecossistemaC *Ecossistema
}

// Plantanovo : Criar instancia de planta
func Plantanovo(nome string, adulto int, reproducao int, vida int, cor uint32, ecossistemaC *Ecossistema) *produtor {

	p := produtor{_adultociclo: adulto}

	p.organismo = *Organismonovo(nome)
	p._nome = nome
	p._nomecompleto = ""
	p._idade = 0
	p._status = "vivo"
	p._fase = "nascido"
	p._adultociclo = adulto

	p._reproduzirciclo = reproducao
	p._reproduzircontador = 0

	p._vida = vida
	p._posx = 0
	p._posy = 0

	p._cor = cor
	p._ecossistemaC = ecossistemaC

	return &p
}

func (p *produtor) vivendo() {

	p.organismo.vivendo()

	if p._status == "vivo" {

		if p._fase == "nascido" {
			if p._idade >= p._adultociclo {
				p._fase = "adulto"

				fmt.Println("       --- Produtor : ", p.nome(), " Evoluiu : Adulto !!!")

			}
		}

		// Se o organismo for adulto inicia o ciclo de reproducao
		if p._fase == "adulto" {

			p._reproduzircontador += 1

			if p._reproduzircontador >= p._reproduzirciclo {
				p._reproduzircontador = 0
				fmt.Println("       --- Produtor : ", p.nome(), " Reproduzindo !!!")

				var pg = Plantanovo(p._nome, p._adultociclo, p._reproduzirciclo, p._vida, p._cor, p._ecossistemaC)
				var x int = utils.Aleatorionumero(50)
				var y int = utils.Aleatorionumero(50)

				pg.mudarposicao(x, y)

				p._ecossistemaC.AdicionarProdutor(pg)
			}

		}

		if p._idade >= p._vida {
			p._status = "morto"
			fmt.Println("       --- Produtor : ", p.nome(), " Morreu !!!")
		}

	}

}

func (p *produtor) toString() string {

	var str = p._nomecompleto + " [" + p.fase() + " " + strconv.Itoa(p.ciclos()) + "]" + " POS[" + strconv.Itoa(p.x()) + " " + strconv.Itoa(p.y()) + "] - Status : " + p._status

	return str
}