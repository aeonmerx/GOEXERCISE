package main
/*Practica 02: Calcular cociente y el Residuo de dos
 números enteros
Enunciado: halar el cociente y el residuo (resto)
 de dos números enteros.
Análisis: para la solución de este problema,
 se requiere que el usuario ingrese dos 
 números enteros y el sistema realice el cálculo 
respectivo para hallar el cociente y residuo.

*/
import "fmt"
func cociente(a,b int)int{
return a/b
}
func residuo(a,b int)int{
	return a%b

}
func main() {
	var a,b,r,s int
	fmt.Println("Ingrese numero 1")
	fmt.Scanln(&a)
	fmt.Println("Ingrese numero 1")
	fmt.Scanln(&b)
	 
	cociente(a,b)
	residuo(a,b)
	r =  cociente(a,b)
	s =  residuo(a,b)
	fmt.Println("El cociente es:", r)
	fmt.Println("El residuo es:", s)
}