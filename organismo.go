package main

type organismo struct {
	_nome   string
	_idade  int
	_status string
	_fase   string
}

func Organismonovo(nome string) *organismo {

	p := organismo{_nome: nome}
	p._idade = 0
	p._status = "vivo"
	p._fase = "nascido"

	return &p
}

func (p *organismo) vivendo() {

	if p._status == "vivo" {

		p._idade++

		if p._fase == "nascido" {

		}

	}

}

func (p *organismo) nome() string   { return p._nome }
func (p *organismo) fase() string   { return p._fase }
func (p *organismo) status() string { return p._status }
func (p *organismo) ciclos() int    { return p._idade }