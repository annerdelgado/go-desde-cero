package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, Gophers")
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())
}

//inicializar un modulo en go es -> got mod init
//obtener un paquete en go es -> got get y la ruta del repositorio del paquete
// go run hola.go
// go build hola.go
