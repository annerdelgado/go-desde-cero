package main

import (
	"fmt"
	"strconv"

	"rsc.io/quote"
)

// var firstName, lastName, age = "Anner", "Delgado", 23

// Declaracion de constantes
// const pi float32 = 3.14
const Pi = 3.14
const (
	X = 100
	Y = 0b1010 //binario
	Z = 0o12   //octal
	W = 0xFF   //hexadecimal
)

// const (
//
//	Domingo = 1
//	Lunes = 2
//	Martes = 3
//
// )
const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Hello, Gophers")
	// fmt.Println(quote.Hello())
	fmt.Println(quote.Go())

	///////Declaracion de variables///////

	// var firstName, lastName string
	// var age int

	// var (
	// 	firstName, lastName string
	// 	age                 int
	// )

	// firstName = "Anner"
	// lastName = "Delgado"
	// age = 23

	// var (
	// 	firstName string = "Anner"
	// 	lastName  string = "Delgado"
	// 	age       int    = 23
	// )

	// var (
	// 	firstName = "Anner"
	// 	lastName  = "Delgado"
	// 	age       = 23
	// )

	// var firstName, lastName, age = "Anner", "Delgado", 23

	// firstName, lastName, age := "Anner", "Delgado", 23 // operador := sirve para crear nuevas variables las cuales solo se pueden usar donde son declaradas

	// fmt.Println(firstName, lastName, age)
	// fmt.Println(Pi)
	// fmt.Println(X, Y, Z, W)
	// fmt.Println(Viernes)

	//////////////Tipos de datos//////////////

	//el tipo de int indica el numero de bits que almacena
	//int8 desde -128 hasta 127
	// var integer int8 = 127 //-128

	//cuando colocamos simplemente int la cantidad de numeros que almacenará depende del SO
	//si la arquitectura del SO es de 32bits le asignará int32, si es de 64 bits le asignará int64
	// var integer int = 10

	//cuando utilizamos el uint alamcenará unicmaente numeros positivos
	// var integer uint = 10

	//para almacenar numeros con decimales, además para almacenar numeros muy grandes, luego del int64 viene el float32
	// var float float32

	// fmt.Println(math.MinInt64, math.MaxInt64) //valores minimo y maximo que puede almacenar int64
	// fmt.Println(math.MaxFloat32)              //valor maximo que puede almacenar float32
	// fmt.Println(math.MaxUint8)                //valor minimo 0 y maximo que puede almacenar uint8

	//declarar booleanos
	// var valueBool bool = true //false

	//una forma de declarar strings
	// fullname := "Anner Delgado \t(alias \"adelgado\")\n" //la \t simboliza un espacio, el caracter \ es para que permita colocar " como texto del string
	// fmt.Println(fullname)

	//el tipo de dato byte es practicamente igual al uint8 que almacena numeros de 0 a 255 unicamente positivos
	//este se utiliza comunmente para almacenar caracteres ascii
	// var a byte = 'a' //imprime el valor en ascii o decimal de a
	// fmt.Println(a)

	// s := "hola"
	// fmt.Println(s[3]) //imprime el valor en bytes del caracter en x posicion del string o cadena

	//el tipo de dato rune es practicamente igual al int32, el cual es utiliza para representara caracteres de tipo unicode
	// var r rune = '😬'
	// fmt.Println(r)

	//////////////Valores predeterminados//////////////
	//todas las variables al ser declaradas en go, tienen un valor predeterminado conocido como valor 0, el cual depende del tipo de dato que sea la variable
	//valor predeterminado para int, float o uint es 0
	//valor predeterminado para bool es false
	//valor predeterminado para string es una cadena vacía ""

	// var (
	// 	defaultInt    int
	// 	defaultUint   uint
	// 	defaultFloat  float32
	// 	defaultBool   bool
	// 	defaultString string
	// )

	// fmt.Println(defaultInt, defaultUint, defaultFloat, defaultBool, defaultString)

	//////////////Conversiones de tipos de datos//////////////
	var integer16 int16 = 50
	var integer32 int32 = 100

	fmt.Println(integer16 + int16(integer32))
	fmt.Println(int32(integer16) + integer32)

	s := "100"
	//la librería strconv sirve para convertir cadenas a otros tipos de datos y la funcion Atoi() convierte a int el string
	i, _ := strconv.Atoi(s) //el _ se usa cuando una funcion devuelve dos o mas valores pero uno de ellos no deseamos recuperarlo
	fmt.Println(i + i)

	n := 42
	s = strconv.Itoa(n) //convertir int a string
	fmt.Println(s)
}
