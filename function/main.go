package main

import "fmt"

type Carro struct {
	Nome string

}

func(c Carro) andar(){
	fmt.Println(c.Nome, "andou")
}

func main() {

	carro := Carro{
		Nome: "Fusca",
	}
	
	carro.andar()
	
	resultado := func(x ...int) func() int{
		resultado := 0
		
		for _,v := range x {
			resultado += v
		}

		return func() int{
			return resultado * resultado
		}

	}
	fmt.Println(resultado(54,54,54,54)())
}

func soma(a int, b int) (result int) {
	result = a + b
	return
}

func somaTudo(x ...int) int{
	resultado := 0
	for _, v := range x {
		resultado += v
	}
	return resultado
}
