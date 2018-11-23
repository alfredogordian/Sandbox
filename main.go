package main

import (
	"awesomeProject/routes"
	"fmt"
)

var i int
var f float64
var b bool
var s string


func main() {


	routes.Calcular()
	fmt.Println("*********************************")
	fmt.Println("Hello, 世界 ❤️❤️❤️")
	// Esto es un comentario
	/*
	Este es un bloque

	 */

	 routes.Routes()
	 routes.Routes2()

	var foo string = "Este es un string"

	variable := "Esta es una variable que está seteada como se hace generalmente"
	fmt.Printf("Hola %s", foo)
	fmt.Printf("Hola %s /n", variable)


	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	/*var s string

	soySlice:= []string{"Uno","Dos","Tres"}*/

	/*for i:=0 :i <len(soySlice.): i++ {
		s += soySlice[i] + " "
	}*/

	fmt.Println(s)
}

func init() { fmt.Printf("Se ejecuta init ************************") }
