package main

import (
	"fmt"
)

var (
	meuArray [3]int
)

func main(){
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30


	fmt.Println("O tamanho do meu array é", len(meuArray))
	fmt.Printf("O tipo de meuArray é %T e o conteúdo é %v\n", meuArray, meuArray)

	for i, v := range meuArray {
		fmt.Printf("O valor do indice %d é %d\n", i, v)
	}

}
