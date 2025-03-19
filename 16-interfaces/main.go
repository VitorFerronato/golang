package main

import "fmt"

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A area da forma é %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)
}
