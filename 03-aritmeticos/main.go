package main

import "fmt"

func main() {
	a := 20
	b := 10
	//SUMA
	result := a + b
	fmt.Println("Suma:",result)
	//RESTA
	result = a - b
	fmt.Println("Resta:",result)
	//MULTIPLICACIÒN
	result = a * b
	fmt.Println("Suma:",result)
	//DIVISIÒN
	result = a / b
	fmt.Println("Divisiòn:",result)
	//DIVISIÒN EXACTA
	var div  float64 = 3.0 / 2.0
	fmt.Println("Divisiòn exacta:",div)
	//MODULO
	result = a % b
	fmt.Println("Mòdulo:",result)
}