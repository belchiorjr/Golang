package main

import (
	"errors"
	"fmt"
)


func main() {
	//int8 int16 int32 int64
	// soment int e usara o padrao do SO
	var numero8 int32 = 1000000000
	var numero16 int16 = 100
	fmt.Println(numero8)
	fmt.Println(numero16)


	// alias rune = int32
	var numeroRune rune = 13456
	fmt.Println(numeroRune)

	// alias byte = int8
	var numeroByte byte = 100
	fmt.Println(numeroByte)


	/// FLOAT float32. float64

	var numeroFloat32 float32 = 65465464546546654646465465465464654643.465465456456564646445654564654564644645664
	var numeroFloat64 float64 = 6546546464564646546465464665546465465987979887989987987987987987987897987987987987978987987465465465464654645.54646445446545454644456546546666666666666666666666666666464565465466545646465
	fmt.Println(numeroFloat32)
	fmt.Println(numeroFloat64)



	/// String não tem o char
	var str string = "Texto"
	fmt.Println(str)

	// Numero ASCCI
	ascii := 'B'
	fmt.Println(ascii)



	/// Valor Standards
	var textoStandard string
	fmt.Println(textoStandard) // Imprime espaco

	var numeroStandard int
	fmt.Println(numeroStandard) // Imprime o zero


	/// BOOLEAN
	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleanoStandard bool
	fmt.Println(booleanoStandard)


	/// ERROR
	var erroStandard error
	fmt.Println(erroStandard)

	var erroCreated error = errors.New("Erro Interno")
	fmt.Println(erroCreated)




}