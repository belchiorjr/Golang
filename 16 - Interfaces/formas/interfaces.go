package main

import (
	"math"
	"fmt" 
)

type forma interface{
	area() float64
}

func escreverArea (f forma) {
	fmt.Printf("A Área da forma é %0.2f\n", f.area())
}

type retangulo struct {
	altura float64
	largura float64
}

func (r retangulo) area () float64{
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area () float64{
	return math.Pow(c.raio , 2) * math.Pi
}

func main() {
	fmt.Println("Interfaces em GO")

	r := retangulo{10,15}
	escreverArea(r)
	
	c := circulo{10}
	escreverArea(c)

}