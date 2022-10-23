package main

import "fmt"

// interfacenya
type HasName interface {
	GetName() string
}

type DetailAnimal interface {
	GetName() string
	GetJenis() string
	GetBuas() bool
}

// mendefinisikan struct
type Person struct {
	Name string
}

type Animal struct {
	Name  string
	Jenis string
	Buas  bool
}

func GetDetailsAnimal(animal DetailAnimal) {
	fmt.Println("Nama", animal.GetName())
	fmt.Println("Jenis", animal.GetJenis())
	fmt.Println("Buas", animal.GetBuas())
}

// punya kontrak dengan interface HasName
func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// struct method
func (person Person) GetName() string {
	return person.Name
}

func (a Animal) GetName() string {
	return a.Name
}

func (a Animal) GetJenis() string {
	return a.Jenis
}

func (a Animal) GetBuas() bool {
	return a.Buas
}

func main() {
	var arief Person
	arief.Name = "Arief"

	var cat Animal
	cat.Name = "Kucing"
	cat.Jenis = "Mamalia"
	cat.Buas = false

	GetDetailsAnimal(cat)

	SayHello(arief)
}
