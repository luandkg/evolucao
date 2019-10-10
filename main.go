package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

// Projeto para Linguagens de Progrmacao 2019.02

// AUTOR : LUAN ALVES - 17/0003191
// AUTOR : MARUAN -

func main() {

	tb := Tabuleiro_novo("MATRIZ")

	tb.limpar()

	lsplantas := list.New()

	lsplantas.PushBack(Planta_novo("Capim Gordura", 5))
	lsplantas.PushBack(Planta_novo("Capim Verde", 10))

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var x int = r1.Intn(50)
	var y int = r1.Intn(50)

	tb.mudar(x, y, 1)

	//produtores := &produtor{}
	//p1 := Produtor_Novo("Capim Gordura", 5)
	//produtores.Append(&p1)expression

	tb.mostrar()

	var ciclo int = 0

	for {

		fmt.Println("---------------- Ciclo :  ", ciclo, " --------------------------------")
		time.Sleep(time.Second)
		fmt.Println("")

		fmt.Println("PRODUTORES")

		for plantac := lsplantas.Front(); plantac != nil; plantac = plantac.Next() {

			p := plantac.Value
			fmt.Println(p)

			//fmt.Println("      - ", p.nome, " [", p.fase, ",", p.ciclos, "]")
			//if capimgorgura.status() == "vivo" {
			//			.nome()//	capimgorgura.vivendo()
			//	}planta)

			//if capimverde.status() == "vivo" {
			//	capimverde.vivendo()
			//}
		}

		ciclo++

		fmt.Println("")

		if ciclo >= 15 {
			break
		}

	}

	fmt.Println("Fim da Simulação !!!")

}
