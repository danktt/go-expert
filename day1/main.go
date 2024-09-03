package main

const scope = "global"

var (
	a bool   = true
	b int    = 10
	c string = "Hello, World!"
)

func main() {
	// Inferencia de tipo, nao precisa declarar o tipo da variavel
	variavel := "variavel" // só pode ser usada dentro da função
	println(variavel)
	println("Hello, World!")

	println(a)

	a = true
	println(a)

	println(scope)
}
