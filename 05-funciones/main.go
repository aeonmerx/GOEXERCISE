package main
// FUNCIONES EN GO , LLAMAR FUNCIONES EN MAIN
import "fmt"

func saludar(nombre string) {
	// fmt.Println("Hola mundo")
fmt.Println("Hola", nombre)
}
func sumar(a,b int) int {
return a+b 
}

func main() {
saludar("Andres")
r := sumar(10,20)
print("La suma es ", r)
}