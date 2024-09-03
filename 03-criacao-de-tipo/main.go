package main

type ID int

type Person struct {
	Name   string
	Age    int
	Height float64
	Weight float64
	Gender string
}

type Address struct {
	Street  string
	Number  int
	City    string
	State   string
	Country string
}

type User struct {
	Person
	Address
}

func main() {

	sayHello("Danilo")
}

func sayHello(name string) {
	println("Hello, " + name)
}
