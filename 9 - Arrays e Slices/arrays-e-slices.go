package main

import (
	"reflect"
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")
	
	var array1 [5]int
	fmt.Println(array1)
	
	var array2 [5]string
	array2[0] =  "Posição 1"
	array2[1] =  "Posição 2"
	array2[2] =  "Posição 3"
	array2[3] =  "Posição 4"
	array2[4] =  "Posição 5"
	fmt.Println(array2)

	array3 := [3] string {"Posicção 1", "Posicção 2", "Posicção 3",}
	fmt.Println(array3)
	
	array4 := [...] int {1,2,3,4,5,6,7,8,9}
	fmt.Println(array4)
	
	slice := [] int {2,5,6,5,8}

	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))
	
	// Append
	slice = append(slice, 100)
	fmt.Println(slice)
	
	slice2 := array3[1:3]
	fmt.Println(slice2)
	
	// Arrays Internos
	fmt.Println(".................................")
	
	slice3 := make([]float32, 10,11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Tamanho
	fmt.Println(cap(slice3)) // Capacidade

	slice3 = append(slice3, 5) 
	fmt.Println(len(slice3)) // Tamanho
	fmt.Println(cap(slice3)) // Capacidade
	
	slice3 = append(slice3, 12)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Tamanho
	fmt.Println(cap(slice3)) // Capacidade
	
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // Tamanho
	fmt.Println(cap(slice4)) // Capacidade
	
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4)) // Tamanho
	fmt.Println(cap(slice4)) // Capacidade

}