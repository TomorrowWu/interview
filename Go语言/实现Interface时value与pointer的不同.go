package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {
	// https://play.golang.org/p/TvR758rfre

	//为什么Cat{}不算实现了Animal interface？
	// &Llama{}算实现了interface？

	//go 会把指针进行隐式转换得到 value，但反过来则不行

	animals := []Animal{Dog{}, Cat{}, Llama{}, &Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
