package main

func main() {
	a:= 10
	println(&a)
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	
	println(ponteiro)
	println(a)
	println(b)
	println(*b)
	*b=30
	println(a)
	println(b)
	a=50
	println(b)
	println(*b)
} 