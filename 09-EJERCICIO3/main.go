package main

import "fmt"
/*Practica 03: Calcular el Precio de Venta
Enunciado: dado el valor de un producto a costo proveedor,
 hallar el IVA (19%) y el precio de venta .
Análisis: para la solución de este problema,
 se requiere que el usuario ingrese el valor de
  venta del producto y el sistema realice el
   cálculo respectivo
 para hallar el IGV y el precio de venta.
 */
 //Función que calcula el igv
func calcularIva(valorproducto, iva float64) float64 {
    return iva * valorproducto
}
//Funcion que calcula precio de venta
func precioVenta(valorproducto, iva float64) float64 {
    return valorproducto + iva
}
func main() {

    //Variables
	var valorproducto, iva, precioventa float64
   //Entrada  de datos
	fmt.Print("Ingrese Valor de Venta: ")
    fmt.Scanln(&valorproducto)
	//Calcular IVA
	iva = calcularIva(valorproducto, 0.19) 
		//Calcular PRECIO DE VENTA
			precioventa = precioVenta(valorproducto, iva)
	//Salida de datos
    fmt.Println("==========================")
    fmt.Println("Valor de Venta:", valorproducto)
    fmt.Println("IVA:", iva)
    fmt.Println("Precio de Venta:", precioventa)
    fmt.Println("==========================")
}