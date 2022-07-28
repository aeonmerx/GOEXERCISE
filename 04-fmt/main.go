package main

import "fmt"

func main() {
	hola := "hola"
	mundo := "mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Alex"
	edad := 26
	fmt.Printf("hola, %s y su edad es %d \n",nombre,edad)
	fmt.Printf("hola, %s y su edad es %d \t",nombre,edad)
	fmt.Printf("hola, %v y su edad es %v \n",nombre,edad) 
	mensage := fmt.Sprintf("Hola, %s y su edad es %d", nombre, edad)
	
	fmt.Println(mensage)
	fmt.Printf("nombre: %T \n",edad) //SABER TIPO DE DATO
	fmt.Print("Ingrese otro nombre")
	fmt.Scanln(&nombre) // GUARDAR VARIABLES CON EL SCANNER 
	print("Otro nombre:",nombre) //IMPRIMIR SIN FMT CON FUNCIÒN PRINT
}