package routes

import "fmt"

func init()  {
	fmt.Println("Funcion init de ROUTES 2")
}

func Routes2(){
	routesPrivate()
}

func routesPrivate(){
	fmt.Println("Impreso desde funcion privada")
}