package main

import "fmt"

func main() {
	//ARREGLOS
	var numeros [5]int
	fmt.Println(numeros)
	numeros[0] = 10
	numeros[1] = 20
	numeros[2]= 30
	fmt.Println(numeros)
	fmt.Println(numeros[4])
	//ARRAYS CON VALORES
	nombres := [3]string{"Alex","Juana","Luisa"}
	fmt.Println(nombres)
	colores := [...]string {
		"Rojo",
		"Verde",
		"Negro",
		"Azul",
	}
	fmt.Println(colores,len(colores))

	//INDICES DEFINIDOS
	monedas := [...]string{
		0: "Dólar", 
		2: "Soles",
		3: "Pesos",
		5: "Euro",
	}
	fmt.Println(monedas,len(monedas))


	//SUB ARRAY
	subMoneda := monedas[3:]
	fmt.Println(subMoneda)
}