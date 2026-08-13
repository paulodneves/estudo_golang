package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Empresa struct {
	Nome string
}

// implementa automaticamente
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func (e Empresa) Desativar() {
	fmt.Println("A empresa foi desativada")
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", wesley.Nome, wesley.Idade, wesley.Ativo)

	wesley.Ativo = false

	wesley.Cidade = "Maceió"
	wesley.Endereco.Numero = 370

	fmt.Println(wesley.Ativo)
	fmt.Println(wesley.Endereco)
	
	Desativacao(wesley)
	
	minhaEmpresa := Empresa{}
	Desativacao(minhaEmpresa)

}
