package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Guilherme"
	var intNum = 12
	var floatNum float32 = 1.2

	fmt.Println("Olá", name)
	fmt.Println("Numero inteiro definido", intNum)
	fmt.Println("Numero float definido:", floatNum)

	fmt.Println("O tido da variavel name é:", reflect.TypeOf(name))
	fmt.Println("O tido da variavel intNum é:", reflect.TypeOf(intNum))
	fmt.Println("O tido da variavel intFloat é:", reflect.TypeOf(floatNum))
}
