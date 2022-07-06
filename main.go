package main

import (
	Dog "github.com/fish895623/hello/modules"
	IAnimal "github.com/fish895623/hello/modules/interfaces"
)

func main() {
	var dog IAnimal.Animal = Dog.NewDog("asdf")
	println(dog.Sounds())
}
