package main

func soma(a int, b *int) int {
	a = 50
	*b = 100 //altera o valor na memória
	return a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	resultado := soma(minhaVar1, &minhaVar2)
	println(minhaVar1, minhaVar2)
	println(resultado)
}
