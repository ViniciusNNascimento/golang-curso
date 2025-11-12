package main

import (
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
}

type Profissao struct  {
	Pessoa
	Tipo string
}

func main() {


	// fmt.Println(Pessoa{"Vinic", 23})
	// fmt.Println(Pessoa{Nome: "Vinic", Idade: 23}) // jeito recomendado de fazer
	// fmt.Println(Pessoa{Nome: "Vinic"})

	p1 := Pessoa{Nome: "Vinic", Idade: 23}

	p1.Nome = "AnakinSkywalker"
	// fmt.Println(p1.Nome)

	p1.Idade = 27
	// fmt.Println(p1.Idade)
	

	p2 := Pessoa{Nome: "Yoda", Idade: 750}
	// p3 := Pessoa{Nome: "Obi wan", Idade: 25}
	// // fmt.Println(p2)

	// pessoas := []Pessoa{}
	// pessoas2 := []Pessoa{ p1}
	// pessoas2 = append(pessoas2, p3 )
	// pessoas = append(pessoas, p1, p2)
	// fmt.Println(pessoas)

	// alunos := map[string][]Pessoa{}
	// alunos2 := map[string][]Pessoa{}
	// alunos["Programação"] = pessoas
	// alunos2["Engenharia"] = pessoas2
	// fmt.Println(alunos)
	// fmt.Println(alunos2)


	// var alunos = map[string][]Pessoa{
	// 	"Programação": {{Nome: "Obiwan", Idade: 35}, {Nome: "Darth vader", Idade: 35}},
	// 	"Engenharia": {{Nome: "Obiwan", Idade: 35}, {Nome: "Darth vader", Idade: 35}},
	// }
	// fmt.Println(alunos)


	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Pessoa.Idade)
	fmt.Println(prof.Tipo)


}
