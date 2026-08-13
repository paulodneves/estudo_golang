package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
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

	wesley.Desativar()

}
