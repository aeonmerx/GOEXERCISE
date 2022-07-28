package main

import "fmt"

/* Practica 01: Suma de dos números enteros
Enunciado: dado dos números enteros, hallar la suma.

Análisis: para la solución de este problema,
 se requiere que el usuario ingrese dos números
  enteros y el sistema realice el cálculo respectivo
   para hallar la suma.
*/
//1 CREAR FUNCIÒN SUMAR  DOS NÙMEROS
func sumar(x, y int) int {
	 return x+y
}
// FUNCIÒN PRINCIPAL
func main() {
	// 2. DEFINIR LAS VARIABLES
	var x,y, suma int

	//ENTRADA DE DATOS
	fmt.Print("Ingrese nùmero 1")
	fmt.Scanln(&x)
	fmt.Print("Ingrese nùmero 2")
	fmt.Scanln(&y)
//LLAMAR FUNCIÒN SUMAR
suma = sumar(x,y)
//SALIDA DE DATOS
fmt.Println("La suma es:",suma)
}