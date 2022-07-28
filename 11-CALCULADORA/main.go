package main
import "fmt"
func sumar(x, y int) int {
	return x+y
}
func restar(x, y int) int {
	return x+y
}
func multiplicar(x, y int) int {
	return x+y
}
func dividir(x, y int) int {
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